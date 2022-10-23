-- 2020年最后一次登录  The Latest login in 2020

SELECT user_id ,max(time_stamp) last_stamp  FROM Logins 
where time_stamp  BETWEEN '2020-01-01' AND '2021-01-01'
GROUP by user_id 