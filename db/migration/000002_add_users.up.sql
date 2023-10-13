CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "hashed_password" varchar not null,
    "full_name" varchar not null,
    "email" varchar unique not null,
    "password_change_at" timestamptz not null default '0001-01-01 00:00:00Z',
    "created_at" timestamptz not null default (now())
);

alter table "accounts" add foreign key ("owner") references "users" ("username");

-- create unique index on "accounts" ("owner", "currency")
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
