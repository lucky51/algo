-- 部门工资前三高的所有员工

SELECT
    Department,
    Employee,
    Salary
from (
        SELECT
            Department.name as Department,
            Employee.salary as Salary,
            Employee.name as Employee,
            dense_rank() over(
                partition by Employee.departmentId
                ORDER BY
                    Employee.salary desc
            ) rn
        FROM Employee
            JOIN Department ON Employee.departmentId = Department.id
    ) as t
where rn < 4