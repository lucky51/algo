-- 部门工资最高的员工

-- 第一次看到 IN 居然可以两个字段

SELECT
    Department.name as Department,
    Employee.name as Employee,
    Salary
FROM Employee
    JOIN Department ON Employee.DepartmentId = Department.id
WHERE (Employee.DepartmentId, Salary) IN (
        SELECT
            DepartmentId,
            Max(Salary)
        FROM Employee
        GROUP BY DepartmentId
    )