CREATE TABLE "tokens" (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR NOT NULL DEFAULT '',
    "address" VARCHAR NOT NULL DEFAULT '',
    "symbol" VARCHAR NOT NULL DEFAULT '',
    "logo_url" VARCHAR NOT NULL DEFAULT '',
    "decimals" bigint NOT NULL DEFAULT 0,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
