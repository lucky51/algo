-- 查找每个员工花费的总时间

SELECT e.event_day day, e.emp_id, sum(e.out_time-e.in_time) as  total_time  
FROM Employees  e GROUP by e.emp_id  ,e.event_day  