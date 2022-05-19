#!/bin/sh
set -e

echo "Migrating hydra!"

hydra migrate -c /etc/config/hydra/hydra.yaml sql -e --yes
