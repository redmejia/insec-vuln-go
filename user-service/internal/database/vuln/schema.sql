CREATE TABLE IF NOT EXISTS vuln_register_users (
    bus_id SMALLSERIAL PRIMARY KEY,
    bus_name TEXT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS vuln_login_users (
    bus_id SMALLINT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS vuln_users_new_deal (
	deal_id SMALLSERIAL PRIMARY KEY,
	bus_id SMALLINT NOT NULL,
    bus_name TEXT NOT NULL,
	pro_name TEXT NOT NULL, -- product description
	pro_description TEXT NOT NULL,
	created_at TIMESTAMP, 
	price NUMERIC(5, 2)
);









