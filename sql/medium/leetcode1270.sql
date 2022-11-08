-- 向公司CEO汇报工作的所有人

SELECT e1.employee_id
from employees e1
JOIN employees e2 on e1.manager_id=e2.employee_id
JOIN employees e3 on e2.manager_id=e3.employee_id
where e1.employee_id!=1 AND e3.manager_id =1



with recursive temp as (
    -- 查找直接汇报人为1的员工
    select * from employees where manager_id =1 and employee_id !=1
    UNION ALL
    -- 查找间接汇报人为1的员工
    select employees.* from employees JOIN temp on temp.employee_id=employees.manager_id
)
select employee_id from temp