CREATE TABLE accounts_user (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username character varying(255) UNIQUE,
    passowrd character varying,
    email character varying(255) UNIQUE,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);
