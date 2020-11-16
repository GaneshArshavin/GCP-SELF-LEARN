    CREATE TABLE accounts_user (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username character varying(255) UNIQUE,
    passowrd character varying,
    email character varying(255) UNIQUE,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_active boolean
);


CREATE TABLE accounts_activity (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id character varying(255),
    company_name character varying(255),
    operation_type character varying(255),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_success boolean
);
