-- 1965. 丢失信息的雇员

SELECT ISNULL(e.employee_id,
        s.employee_id) AS employee_id
FROM Employees e FULL
JOIN Salaries s
    ON e.employee_id = s.employee_id
WHERE e.employee_id is NULL
        OR s.employee_id is NULL
ORDER BY  employee_id