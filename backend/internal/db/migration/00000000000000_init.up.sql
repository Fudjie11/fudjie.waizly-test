BEGIN;

-- Tabel Employee
CREATE TABLE public.employee (
    employee_id uuid NOT NULL CONSTRAINT employee_pk PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    job_title VARCHAR(50) NOT NULL,
    salary decimal(15, 4) not null,
    department VARCHAR(50) NOT NULL,
    joined_date TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Index For Tabel Employee
CREATE INDEX idx_employee_name ON public.employee (name);
CREATE INDEX idx_employee_job_title ON public.employee (job_title);
CREATE INDEX idx_employee_department ON public.employee (department);

-- Tabel Employee
CREATE TABLE public.sales (
    sales_id uuid NOT NULL CONSTRAINT sales_pk PRIMARY KEY,
    employee_id uuid NOT NULL,
    sales decimal(15, 4) not null
);

CREATE INDEX idx_sales_employee_id ON public.sales (employee_id);

COMMIT;