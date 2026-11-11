#!/usr/bin/env bash
set -euo pipefail

if [ -z "${POSTGRES_DSN:-}" ]; then
  echo "POSTGRES_DSN not set"
  exit 1
fi

# Apply SQL migrations if you manage them externally. This script is a placeholder.
psql "$POSTGRES_DSN" -f migrations/001_init.sql
