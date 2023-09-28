CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "currency" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL default (now())
);

CREATE table "entries" (
                           "id" bigserial primary key,
                           "account_id" bigint,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL default (now())
);

create table "transfer" (
                            "id" bigserial primary key,
                            "from_account_id" bigint,
                            "to_account_id" bigint,
                            "amount" bigint not null,
                            "created_at" timestamptz not null default (now())
);

alter table "entries" add foreign key ("account_id") references "accounts" ("id");

alter table "transfer" add foreign key ("from_account_id") references "accounts" ("id");

alter table "transfer" add foreign key ("to_account_id") references "accounts" ("id");

create index on "accounts" ("owner");

create index on "entries" ("account_id");

create index on "transfer" ("from_account_id");

create index on "transfer" ("to_account_id");

create index on "transfer" ("from_account_id", "to_account_id");

comment on column entries.amount IS 'can be negative or positive';

comment on column transfer.amount IS 'must be positive';

