l
-- 超过经理收入的员工

SELECT e1.name as Employee  from Employee  e1
JOIN Employee e2
on e1.managerId = e2.id
WHERE e1.salary > e2.salary and e1.managerId is not NULL
