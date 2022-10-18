package main

import (
	"net/http"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/database"
	web3idp "github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/idp"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/log"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/web3oidc"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/web3oidc/hydra"
)

const configFileName = "config.yaml"

func main() {
	cfg := web3idp.GetConfig(configFileName, true)
	cfg.PrettyPrint()
	db, cleanupDB, err := database.NewDB(&cfg.Database)
	if err != nil {
		panic(err)
	}
	defer cleanupDB()
	hydraClient, err := hydra.NewClient(&cfg.Hydra)
	if err != nil {
		panic(err)
	}
	web3OIDCHandler := web3oidc.NewHandler(hydraClient, db)

	log.Logf(0, "Listening on %q", cfg.Settings.URL)
	log.Log(0, http.ListenAndServe(cfg.Settings.URL, web3OIDCHandler.Handle()))

}
