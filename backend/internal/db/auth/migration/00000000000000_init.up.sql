-- Tabel  
CREATE TABLE public.employee (
    employee_id uuid NOT NULL CONSTRAINT tms_role_pk PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    app_domain VARCHAR(15) NOT NULL,
    description TEXT,
    is_deleted BOOLEAN NOT NULL,
    is_active BOOLEAN NOT NULL,
    created_user_id uuid NOT NULL,
    created_time_utc TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_user_id uuid NOT NULL,
    updated_time_utc TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    row_version uuid NOT NULL
);

CREATE INDEX idx_tms_role_app_domain ON public.tms_role (app_domain);
CREATE INDEX idx_tms_role_name ON public.tms_role (role_name);
ON public.oa_customer (id_of_keycloak_client);