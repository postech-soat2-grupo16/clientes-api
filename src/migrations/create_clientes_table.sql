CREATE TABLE IF NOT EXISTS clientes (
	id bigserial NOT NULL,
	email text NULL,
	cpf text NULL,
	"name" text NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT clientes_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_clientes_deleted_at ON clientes USING btree (deleted_at);
