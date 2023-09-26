DROP TABLE IF EXISTS accounts;
CREATE SEQUENCE IF NOT EXISTS account_id_seq;
CREATE TABLE accounts (
    "id" int4 NOT NULL DEFAULT nextval('account_id_seq'::regclass),
    "nama" varchar(255),
    "nik" varchar(255),
    "no_hp" varchar(255),
    "no_rekening" varchar(255),
    "saldo" float8,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS mutations;
CREATE SEQUENCE IF NOT EXISTS mutations_id_seq;
CREATE TABLE mutations (
    "id" int4 NOT NULL DEFAULT nextval('mutations_id_seq'::regclass),
    "no_rekening" varchar(255),
    "kode_transaksi" varchar(1),
    "nominal" float8,
    "created_at" timestamp,
    PRIMARY KEY ("id")
);