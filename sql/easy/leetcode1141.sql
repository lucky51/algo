-- 查询近30天活跃用户数

SELECT a.activity_date day,count(DISTINCT user_id) active_users from Activity a 
where a.activity_date BETWEEN   dateadd(dd,-29,'2019-07-27')  and '2019-07-27' GROUP BY a.activity_date