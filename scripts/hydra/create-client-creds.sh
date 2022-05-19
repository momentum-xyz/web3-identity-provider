#!/bin/sh
set -e

export CLIENT_EXISTS=$(hydra clients list --endpoint http://hydra:4445 | grep auth-code-client)

if [ -z "$CLIENT_EXISTS" ]; then
    echo "Creating test client credentials!"

    hydra clients create \
        --endpoint http://hydra:4445 \
        --id auth-code-client \
        --secret secret \
        --grant-types authorization_code,refresh_token \
        --response-types code,id_token \
        --scope openid,offline \
        --callbacks http://127.0.0.1:5555/callback
fi