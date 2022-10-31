-- 每位经理的下属员工数量

select
    e1.reports_to as employee_id ,
    e2.name,
    count(e1.reports_to) as reports_count,
    round(avg(e1.age), 0) as average_age
from Employees e1
    JOIN Employees e2 on e1.reports_to = e2.employee_id
where
    e1.reports_to is NOT NULL
GROUP BY e1.reports_to, e2.name
ORDER BY employee_id 