#!/bin/sh

normalized_db_str=$(echo "$DATABASE_URL" | awk -F \? '{ print $1 }')
log_level='info'

echo $DATABASE_URL

output=$(/usr/local/bin/goose \
    -dir "/sql/$DATABASE_SCHEMA/migrations" \
    -table schema_migrations \
    postgres "$normalized_db_str"\
    up \
    2>&1)

if [ $? -gt 0 ]; then log_level='error'; fi

# output as json
jq --null-input --compact-output --arg msg "$output" --arg lvl "$log_level" '{ level: $lvl, msg: $msg }'
