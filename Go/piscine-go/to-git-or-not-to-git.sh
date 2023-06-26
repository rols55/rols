#! /bin/bash
curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"rols55\"}}){id}}"}' | jq ".[]" | jq ".[]"