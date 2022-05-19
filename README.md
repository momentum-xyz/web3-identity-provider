# WEB3-OIDC IDP

The web3-identity-provider bridges 'Connect your wallet' types of authentication with OIDC auth. It allows APIs for passwordless authentication based on public key cryptography with any type of web3 wallet. It can be used with Metamask, other WalletConnect wallets and Polkadot. 

## Deployment Prerequisites
---


### [Ory Hydra](https://www.ory.sh/hydra/docs/)

. [Hydra](https://www.ory.sh/hydra/docs/production) behind an API gateway or a load balancer

. ORY Hydra serves APIs via two ports:

    - Public port (default 4444)
    - Administrative port (default 4445)

. The public port should be exposed to public internet traffic via a custom Odyssey domain for example

 - https://web3.idp.odyssey

. Administrative port (4445) api's can be accessed via service name in internal odyssey netowrk 

. [A MySQL 5.7+ or Postgres database](https://www.ory.sh/hydra/docs/dependencies-environment) for hydra migration and client creation

. Post Database creation DSN env variable should point to the new hydra database for example 

  - export DSN=mysql://root@tcp(db:3306)/hydra_dev?interpolateParams=true&parseTime=true
  
. Once database is setup hydra migrate sql is required to create necessary SQL schemas , refer to [this link](https://www.ory.sh/hydra/docs/cli/hydra-migrate-sql) for more information about hydra migration   

 - hydra migrate -c /etc/config/hydra/hydra.yaml sql -e --yes

. Once migration is successfully done , which can be verified with tables exists in the database like ```hydra_client ``` then a  [new oauth hydra client](https://www.ory.sh/hydra/docs/cli/hydra-clients-create/#hydra-clients-create) need to be create for the authorization code flow , for example 

   ```
   hydra clients create \
        --endpoint http://hydra:4445 \
        --id auth-code-client \
        --secret secret \
        --grant-types authorization_code,refresh_token \
        --response-types code,id_token \
        --scope openid,offline \
        --callbacks http://127.0.0.1:3000/callback
   ```

   > **Note** 
 >  - While creating a hydra client make sure `--secret` flag value should be a secured random string if `--secret` flag is not provided hydra will generate the secret by itself which will print on the cli at the time of client creation , this secret needs to keep safe to revoke the token later or performing some admin operations. 
 > -  `--id` flag is for the client name which can be anything for this example its `auth-code-client`

 > - `--callbacks ` value should be the landing page URL where user will redirect post successful authentication (React client URL)

. If hydra deployement , migration and client creation is successful then to verify everything is working a URL from the browser or curl can be initiate like 

- https://web3.idp.odyssey/.well-known/openid-configuration  

  *sample response from this request should be similar*

```
{
   "issuer":"http://localhost:4444/",
   "authorization_endpoint":"http://localhost:4444/oauth2/auth",
   "token_endpoint":"http://localhost:4444/oauth2/token",
   "jwks_uri":"http://localhost:4444/.well-known/jwks.json",
   "subject_types_supported":[
      "pairwise",
      "public"
   ],
   "response_types_supported":[
      "code",
      "code id_token",
      "id_token",
      "token id_token",
      "token",
      "token id_token code"
   ],
   "claims_supported":[
      "sub"
   ],
   "grant_types_supported":[
      "authorization_code",
      "implicit",
      "client_credentials",
      "refresh_token"
   ],
   "response_modes_supported":[
      "query",
      "fragment"
   ],
   "userinfo_endpoint":"http://localhost:4444/userinfo",
   "scopes_supported":[
      "offline_access",
      "offline",
      "openid"
   ],
   "token_endpoint_auth_methods_supported":[
      "client_secret_post",
      "client_secret_basic",
      "private_key_jwt",
      "none"
   ],
   "userinfo_signing_alg_values_supported":[
      "none",
      "RS256"
   ],
   "id_token_signing_alg_values_supported":[
      "RS256"
   ],
   "request_parameter_supported":true,
   "request_uri_parameter_supported":true,
   "require_request_uri_registration":true,
   "claims_parameter_supported":false,
   "revocation_endpoint":"http://localhost:4444/oauth2/revoke",
   "backchannel_logout_supported":true,
   "backchannel_logout_session_supported":true,
   "frontchannel_logout_supported":true,
   "frontchannel_logout_session_supported":true,
   "end_session_endpoint":"http://localhost:4444/oauth2/sessions/logout",
   "request_object_signing_alg_values_supported":[
      "RS256",
      "none"
   ],
   "code_challenge_methods_supported":[
      "plain",
      "S256"
   ]
}
```


## [web3-identity-provider](https://github.com/OdysseyMomentumExperience/web3-identity-provider/) 

### Deployment prerequisites

1. Go >= v1.17.2
2. Ent - https://entgo.io/docs/tutorial-setup
3. Makefile
4. Docker
5. docker-compose


### Environment variables 

. Database config 
 ```
 mysql:
  database: web3_idp_dev
  url: localhost
  port: 3306
  username: root
  migrate: true
 ```
 . Log config 

```
settings:
  loglevel: 1
  url: 0.0.0.0:4000

```
. Hydra config 
```
hydra:
  adminURL: http://localhost:4445

```

> web3-identity-provider service runtime configuration fetched from the [config.yaml file](https://github.com/OdysseyMomentumExperience/web3-identity-provider/blob/master/config.yaml)

### Development env

#### Run
`go run cmd/main.go`


> 
>  A docker compose file having all the required steps included DB creation , hydra migration and hydra client creation can be found [here](https://github.com/OdysseyMomentumExperience/web3-identity-provider/blob/master/docker-compose.yaml) and hydra configuration is included in [hydra.yaml](https://github.com/OdysseyMomentumExperience/web3-identity-provider/blob/master/hydra.yaml) file 

## Contributors ✨

The web3-identity-provider is a project initiated by Odyssey. Thanks to these contributors 😎

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>

  <tr>
  <td align="center"><a href="https://github.com/jellevdp"><img src="https://avatars.githubusercontent.com/jellevdp?v=3?s=100" width="100px;" alt=""/><br /><sub><b>Jelle van der Ploeg </b></sub></a><br />
    </td>
<td align="center"><a href="https://github.com/tech-sam"><img src="https://avatars.githubusercontent.com/tech-sam?v=3?s=100" width="100px;" alt=""/><br /><sub><b>Sumit</b></sub></a><br />
</td>
   <td align="center"><a href="https://github.com/e-nikolov"><img src="https://avatars.githubusercontent.com/e-nikolov" width="100px;" alt=""/><br /><sub><b>Emil Nikolov  </b></sub></a><br />
    </td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/nwasiqUC"><img src="https://avatars.githubusercontent.com/nwasiqUC" width="100px;" alt=""/><br /><sub><b>Wasiq  </b></sub></a><br />
    </td>
    <td align="center"><a href="https://github.com/antst"><img src="https://avatars.githubusercontent.com/antst" width="100px;" alt=""/><br /><sub><b>Anton Starikov</b></sub></a><br />
    </td>
    <td align="center"><a href="https://github.com/jor-rit"><img src="https://avatars.githubusercontent.com/jor-rit" width="100px;" alt=""/><br /><sub><b>Jorrit</b></sub></a><br />
    </td>
  </tr>
</table>