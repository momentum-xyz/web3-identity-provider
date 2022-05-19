#!/bin/bash
set -e

echo "Creating databases!"

mysql -u root <<-EOSQL
    CREATE DATABASE web3_idp_dev;
    CREATE DATABASE hydra_dev;
EOSQL
