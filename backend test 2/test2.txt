

-- 1. Tampilkan seluruh data dari tabel "employee" (5 Points)
SELECT employee_id, name, job_title, salary, department, joined_date FROM employee;

-- 2. Berapa banyak karyawan yang memiliki posisi pekerjaan (job title) "Manager"? (5 Points)
SELECT COUNT(employee_id) FROM employee WHERE job_title = 'Manager';

-- 3. Tampilkan daftar nama dan gaji (salary) dari karyawan yang bekerja di departemen "Sales" atau "Marketing" (10 Points)
SELECT name, salary FROM employee WHERE department IN('Sales', 'Marketing');

-- 4. Hitung rata-rata gaji (salary) dari karyawan yang bergabung (joined) dalam 5 tahun terakhir (berdasarkan kolom "joined_date") (10 Points)
SELECT AVG(salary) AS average_salary FROM employee WHERE joined_date >= CURRENT_DATE - INTERVAL '5 years';

-- 5. Tampilkan 5 karyawan dengan total penjualan (sales) tertinggi dari tabel "employee" dan "sales" (10 Points)
SELECT e.employee_id, e.name, SUM(s.sales) AS total_sales FROM employee e
LEFT JOIN sales s ON e.employee_id = s.employee_id
GROUP BY e.employee_id, e.name
ORDER BY total_sales DESC
LIMIT 5;

-- 6. Tampilkan nama, gaji (salary), dan rata-rata gaji (salary) dari semua karyawan yang bekerja di departemen yang memiliki rata-rata gaji lebih tinggi dari gaji rata-rata di semua departemen (15 Points)
SELECT e.name, e.salary, avg_dept_salary.avg_salary AS department_avg_salary
FROM employee e
JOIN (
    SELECT department, AVG(salary) AS avg_salary
    FROM employee
    GROUP BY department
) AS avg_dept_salary ON e.department = avg_dept_salary.department
WHERE e.salary  > avg_dept_salary.avg_salary;

-- 7. Tampilkan nama dan total penjualan (sales) dari setiap karyawan, bersama dengan peringkat (ranking) masing-masing karyawan berdasarkan total penjualan. Peringkat 1 adalah karyawan dengan total penjualan tertinggi (25 Points)
SELECT 
    name,
    total_sales,
    RANK() OVER (ORDER BY total_sales DESC) AS rank
FROM (
    SELECT 
        e.employee_id,
        e.name,
        COALESCE(SUM(s.sales), 0) AS total_sales
    FROM 
        employee e
    LEFT JOIN 
        sales s ON e.employee_id = s.employee_id
    GROUP BY 
        e.employee_id, e.name
) AS sales_totals;

-- 8. Buat sebuah stored procedure yang menerima nama departemen sebagai input, dan mengembalikan daftar karyawan dalam departemen tersebut bersama dengan total gaji (salary) yang mereka terima (20 Points)
CREATE OR REPLACE PROCEDURE show_employee_salary_by_department(IN department_name VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
    employee_record RECORD;
BEGIN
    FOR employee_record IN 
        SELECT name, salary 
        FROM employee 
        WHERE department = department_name
    LOOP
        RAISE NOTICE 'employee:%, salary:%', employee_record.name, employee_record.salary;
    END LOOP;    
END;
$$;

-- Run
CALL show_employee_salary_by_department('Sales');
-- Hasilnya 
-- Employee: John Smith, Salary: 60000.0000
-- Employee: Anna Lee, Salary: 65000.0000
