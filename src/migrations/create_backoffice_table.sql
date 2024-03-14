CREATE TABLE IF NOT EXISTS backoffice_requests (
                                                   id bigserial NOT NULL,
                                                   nome text NOT NULL,
                                                   endereco text NULL,
                                                   telefone text NULL,
                                                   processed boolean NOT NULL DEFAULT false,
                                                   created_at timestamptz NULL,
                                                   updated_at timestamptz NULL,
                                                   deleted_at timestamptz NULL,
                                                   CONSTRAINT backoffice_pkey PRIMARY KEY (id)
    );