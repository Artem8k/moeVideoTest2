CREATE TABLE IF NOT EXISTS "advertiser" (
    id                  INTEGER generated always as identity PRIMARY KEY,
    title               TEXT NOT NULL,
    create_date         TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "adCompany" (
    id                  INTEGER generated always as identity PRIMARY KEY,
    title               TEXT NOT NULL,
    create_date         TIMESTAMP NOT NULL,
    advertiser_id       INTEGER NOT NULL REFERENCES "advertiser"(id)
);

CREATE TABLE IF NOT EXISTS "advertisement" (
    id                  INTEGER generated always as identity PRIMARY KEY,
    title               TEXT NOT NULL,
    CPM                 INTEGER NOT NULL,
    create_date         TIMESTAMP NOT NULL,
    adCompany_id       INTEGER NOT NULL REFERENCES "adCompany"(id)
);
