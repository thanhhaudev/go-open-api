CREATE TABLE tenants (
    id integer PRIMARY KEY,
    scope VARCHAR(125) NOT NULL,
    name VARCHAR(125) NOT NULL,
    api_key VARCHAR(125) NOT NULL,
    api_secret VARCHAR(125) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

INSERT INTO tenants (id, scope, name, api_key, api_secret)
VALUES
    (1, 'default', 'Default Tenant', 'KRY2oikKQ4DEgG5VOC57', 'CJxNmBP07PfH1GYZqu1O'),
    (2, 'tenant1', 'Tenant 1', '6yDd4PnFH9MMIdGgOOkf', 'NWHZUUiqTqbIBGMfcLyS'),
    (3, 'tenant2', 'Tenant 2', '4b7Ph2hsJP4ohC0tlw5J', '2UF9c2jvKsUfamAeISli');
