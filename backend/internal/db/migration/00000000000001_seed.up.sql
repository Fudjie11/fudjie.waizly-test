BEGIN;

INSERT INTO "employee" ("employee_id", "name", "job_title", "salary", "department", "joined_date") VALUES
('904f0360-0bab-43a4-acac-9bf6b0c6b1e1', 'John Smith', 'Manager', 60000, 'Sales', '2022-01-15'),
('904f0360-0bab-43a4-acac-9bf6b0c6b1e2', 'Jane Doe', 'Analyst', 45000, 'Marketing', '2022-02-01'),
('904f0360-0bab-43a4-acac-9bf6b0c6b1e3', 'Mike Brown', 'Developer', 55000, 'IT', '2022-03-10'),
('904f0360-0bab-43a4-acac-9bf6b0c6b1e4', 'Anna Lee', 'Manager', 65000, 'Sales', '2022-12-05'),
('904f0360-0bab-43a4-acac-9bf6b0c6b1e5', 'Mark Wong', 'Developer', 50000, 'IT', '2022-05-20'),
('904f0360-0bab-43a4-acac-9bf6b0c6b1e6', 'Emily Chen', 'Analyst', 48000, 'Marketing', '2022-06-02');

INSERT INTO "sales" ("sales_id", "employee_id", "sales") VALUES
('904f0360-0bab-43a4-acac-9bf6b0c6b001', '904f0360-0bab-43a4-acac-9bf6b0c6b1e1', 15000),
('904f0360-0bab-43a4-acac-9bf6b0c6b002', '904f0360-0bab-43a4-acac-9bf6b0c6b1e2', 12000),
('904f0360-0bab-43a4-acac-9bf6b0c6b003', '904f0360-0bab-43a4-acac-9bf6b0c6b1e3', 18000),
('904f0360-0bab-43a4-acac-9bf6b0c6b004', '904f0360-0bab-43a4-acac-9bf6b0c6b1e1', 20000),
('904f0360-0bab-43a4-acac-9bf6b0c6b005', '904f0360-0bab-43a4-acac-9bf6b0c6b1e4', 22000),
('904f0360-0bab-43a4-acac-9bf6b0c6b006', '904f0360-0bab-43a4-acac-9bf6b0c6b1e5', 19000),
('904f0360-0bab-43a4-acac-9bf6b0c6b007', '904f0360-0bab-43a4-acac-9bf6b0c6b1e6', 13000),
('904f0360-0bab-43a4-acac-9bf6b0c6b008', '904f0360-0bab-43a4-acac-9bf6b0c6b1e2', 14000);

COMMIT;