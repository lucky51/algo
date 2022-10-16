-- 计算特殊奖金

-- leetcode   官方提交不知道为什么通过不了，记录的顺序不一致。
SELECT e1.employee_id ,ISNULL(tb.bonus,0)  as bonus FROM Employees  e1 left join 
( SELECT e.employee_id ,salary as bonus  FROM Employees e 
where e.name not like 'M%' and e.employee_id %2=1) as tb 
on e1.employee_id  = tb.employee_id 


select employee_id,
case when name  not like 'M%' and employee_id % 2=1 then salary 
else  0 
end as bonus
from employees