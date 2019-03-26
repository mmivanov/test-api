#!/bin/sh

export MIGRATION_DB_DSN="user=docker password=docker dbname=postgres host=localhost port=5432 sslmode=disable"

if [ "$1" = "--dryrun" ]; then
  goose -dir migrations postgres "$MIGRATION_DB_DSN" status
else
  goose -dir migrations postgres "$MIGRATION_DB_DSN" up
fi
