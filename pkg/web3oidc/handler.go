package web3oidc

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/ChainSafe/go-schnorrkel"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3challenge"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/xhttp"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/xorm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
)

const wallet_polkadot string = "polkadot"

func NewHandler(hydra *client.OryHydra, db *sql.DB) *Handler {
	ent := xorm.NewEnt(db)

	return &Handler{
		hydra: hydra,
		ent:   ent,
	}
}

type Handler struct {
	hydra *client.OryHydra
	ent   *ent.Client
}

// Response for getting login information.
type LoginInfoResponse struct {
	Subject    string   `json:"subject,omitempty"`
	RequestURL string   `json:"requestURL"`
	Display    string   `json:"display,omitempty"`
	LoginHint  string   `json:"loginHint,omitempty"`
	UILocales  []string `json:"uiLocales,omitempty"`
}

type GetLoginResponse struct {
	Redirect string `json:"redirect"`
}

type GetChallengeResponse struct {
	AddressChallenge string `json:"address_challenge"`
}

type PostLoginRequest struct {
	SignedAddressChallenge string `json:"signed_address_challenge"`
	LoginChallenge         string `json:"login_challenge"`
	RememberMe             bool   `json:"remember_me"`
	WalletType             string `json:"wallet_type"`
}

type PostLoginResponse struct {
	Redirect string `json:"redirect"`
}

type PostConsentRequest struct {
	ConsentChallenge string `json:"consent_challenge"`
}

type PostConsentResponse struct {
	Redirect string `json:"redirect"`
}

func (h *Handler) Handle() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`ok`))
	})

	r.Get("/v0/web3/check-skip-login", func(w http.ResponseWriter, r *http.Request) {
		challenge := r.URL.Query().Get("login_challenge")

		res, err := h.hydra.Admin.GetLoginRequest(admin.NewGetLoginRequestParams().
			WithLoginChallenge(challenge))
		if xhttp.Error(w, err, 500) {
			return
		}
		if *res.Payload.Skip {
			// TODO set eth address in the session?

			res, err := h.hydra.Admin.AcceptLoginRequest(admin.NewAcceptLoginRequestParams().
				WithLoginChallenge(challenge).
				WithBody(&models.AcceptLoginRequest{Remember: true, RememberFor: 3600,
					Subject: res.Payload.Subject}))
			if xhttp.Error(w, err, 500) {
				return
			}
			xhttp.Respond(w, r, &GetLoginResponse{
				Redirect: *res.Payload.RedirectTo,
			})
			return
		}
	})

	r.Get("/v0/web3/challenge", func(w http.ResponseWriter, r *http.Request) {
		challenge := r.URL.Query().Get("login_challenge")
		address := r.URL.Query().Get("address")

		_, err := h.hydra.Admin.GetLoginRequest(
			admin.NewGetLoginRequestParams().
				WithLoginChallenge(challenge))
		if xhttp.Error(w, err, 500) {
			return
		}

		w3u, err := getOrCreateWeb3User(h.ent, address)
		if xhttp.Error(w, err, 500) {
			return
		}

		web3Challenge := generateWeb3Challenge(address)

		// Create a new web3 challenge or update the existing one
		err = h.ent.Web3Challenge.Create().
			SetLoginChallenge(challenge).
			SetWeb3User(w3u).
			SetWeb3Challenge(web3Challenge).
			OnConflict().UpdateNewValues().
			Exec(context.Background())
		if xhttp.Error(w, err, 500) {
			return
		}

		xhttp.Respond(w, r, &GetChallengeResponse{
			AddressChallenge: web3Challenge,
		})
	})

	r.Get("/v0/web3/login", func(w http.ResponseWriter, r *http.Request) {
		challenge := r.URL.Query().Get("login_challenge")
		if challenge == "" {
			if xhttp.Error(w, errors.New("Missing login_challenge query parameter."), 400) {
				return
			}
		}
		res, err := h.hydra.Admin.GetLoginRequest(admin.NewGetLoginRequestParams().WithLoginChallenge(challenge))
		if xhttp.Error(w, err, 500) {
			return
		}
		payload := res.GetPayload()
		xhttp.Respond(w, r, &LoginInfoResponse{
			Subject:    *payload.Subject,
			RequestURL: *payload.RequestURL,
			Display:    payload.OidcContext.Display,
			LoginHint:  payload.OidcContext.LoginHint,
			UILocales:  payload.OidcContext.UILocales,
		})
	})

	r.Post("/v0/web3/login", func(w http.ResponseWriter, r *http.Request) {
		req := new(PostLoginRequest)

		err := xhttp.Decode(r, req)
		if xhttp.Error(w, err, 500) {
			return
		}

		walletType := web3user.WalletTypeEth
		if len(req.WalletType) > 0 && wallet_polkadot == req.WalletType {
			walletType = web3user.WalletTypePolkadot
		}

		w3c, err := h.ent.Web3Challenge.Query().
			Where(web3challenge.LoginChallenge(req.LoginChallenge)).
			WithWeb3User().
			Only(context.Background())
		if xhttp.Error(w, err, 500) {
			return
		}

		w3u := w3c.Edges.Web3User

		if walletType == web3user.WalletTypePolkadot {
			/*
				Polkadot address is a decoded hex string passed by the client refer the implementation
				https://github.com/OdysseyMomentumExperience/web3-client-auth/blob/master/src/hooks/web3-hooks.ts
			*/
			err = verifyPolkadotSignature(w3u.Address, req.SignedAddressChallenge, w3c.Web3Challenge)

		} else {
			err = verify(w3u.Address, req.SignedAddressChallenge, w3c.Web3Challenge)
		}

		if xhttp.Error(w, err, 500) {
			return
		}

		sub := w3u.UUID.String()

		res, err := h.hydra.Admin.AcceptLoginRequest(admin.NewAcceptLoginRequestParams().
			WithLoginChallenge(req.LoginChallenge).
			WithBody(&models.AcceptLoginRequest{
				RememberFor: 3600, Remember: req.RememberMe,
				Subject: &sub,
			}))

		if xhttp.Error(w, err, 500) {
			return
		}

		w3u, err = w3u.Update().SetWalletType(walletType).Save(context.Background())
		if xhttp.Error(w, err, 500) {
			return
		}

		xhttp.Respond(w, r, &PostLoginResponse{
			Redirect: *res.Payload.RedirectTo,
		})
	})

	r.Get("/v0/web3/consent", func(w http.ResponseWriter, r *http.Request) {
		challenge := r.URL.Query().Get("consent_challenge")
		if challenge == "" {
			if xhttp.Error(w, errors.New("Missing consent_challenge query parameter."), 400) {
				return
			}
		}
		res, err := h.hydra.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(challenge))
		if xhttp.Error(w, err, 500) {
			return
		}
		payload := res.GetPayload()
		xhttp.Respond(w, r, &LoginInfoResponse{
			Subject:    payload.Subject,
			RequestURL: payload.RequestURL,
			Display:    payload.OidcContext.Display,
			LoginHint:  payload.OidcContext.LoginHint,
			UILocales:  payload.OidcContext.UILocales,
		})
	})

	r.Post("/v0/web3/consent", func(w http.ResponseWriter, r *http.Request) {
		req := new(PostConsentRequest)

		err := xhttp.Decode(r, req)
		if xhttp.Error(w, err, 500) {
			return
		}

		cr, err := h.hydra.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().
			WithConsentChallenge(req.ConsentChallenge))
		if xhttp.Error(w, err, 500) {
			return
		}

		w3u, err := h.ent.Web3User.Query().Where(web3user.UUID(uuid.MustParse(cr.Payload.Subject))).Only(context.Background())
		if xhttp.Error(w, err, 500) {
			return
		}

		res, err := h.hydra.Admin.AcceptConsentRequest(admin.NewAcceptConsentRequestParams().
			WithConsentChallenge(req.ConsentChallenge).
			WithBody(&models.AcceptConsentRequest{
				Session: &models.ConsentRequestSession{
					IDToken: map[string]interface{}{
						"eth":          w3u.Address, // deprecated, TODO remove
						"web3_type":    w3u.WalletType,
						"web3_address": w3u.Address,
					},
				},
				GrantScope:               cr.Payload.RequestedScope,
				GrantAccessTokenAudience: cr.Payload.RequestedAccessTokenAudience,
			}))

		if xhttp.Error(w, err, 500) {
			return
		}

		xhttp.Respond(w, r, &PostLoginResponse{
			Redirect: *res.Payload.RedirectTo,
		})
	})

	return r
}

func getOrCreateWeb3User(ent *ent.Client, address string) (*ent.Web3User, error) {
	id, err := ent.Web3User.Create().SetAddress(address).OnConflict().Ignore().ID(context.Background())
	if err != nil {
		return nil, err
	}

	w3u, err := ent.Web3User.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return w3u, nil
}
func generateWeb3Challenge(address string) string {
	return fmt.Sprintf("Please sign this message with the private key for address %s to prove that you own it\n%s", address, uuid.New().String())
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func verify(address, signature, challenge string) error {
	if !strings.HasPrefix(signature, "0x") {
		signature = "0x" + signature
	}
	sigBytes := hexutil.MustDecode(signature)
	if sigBytes[64] != 27 && sigBytes[64] != 28 {
		return errors.New("unsupported signature format")
	}

	sigBytes[64] -= 27

	msgBytes := signHash([]byte(challenge))

	pubKey, err := crypto.SigToPub(msgBytes, sigBytes)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	if recoveredAddr != address {
		return errors.New("the challenge was not signed by the correct address")
	}

	fmt.Printf("Recovered address: %s\n", recoveredAddr)

	return nil
}

func verifyPolkadotSignature(address, signature, challenge string) error {
	pub, err := schnorrkel.NewPublicKeyFromHex(address)
	if err != nil {
		panic(err)
	}
	sig, err := schnorrkel.NewSignatureFromHex(signature)

	if err != nil {
		panic(err)
	}

	ctx := []byte("substrate")
	transcript := schnorrkel.NewSigningContext(ctx, []byte("<Bytes>"+challenge+"</Bytes>"))
	ok, err := pub.Verify(sig, transcript)
	if !ok || err != nil {
		return errors.New("the challenge was not signed by the correct address")
	}
	fmt.Printf("Polkadot address: %s\n", address)
	return nil
}
