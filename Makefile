# https://github.com/prisma-labs/get-graphql-schema
.PHONY: get-schema
get-schema:
	get-graphql-schema-h "X-AUTH-EMAIL=" -h "X-AUTH-KEY=" https://api.cloudflare.com/client/v4/graphql > schema.graphql

