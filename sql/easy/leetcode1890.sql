<<<<<<< HEAD
-- 2020年最后一次登录  The Latest login in 2020

=======

-- 2020年最后一次登录
>>>>>>> ee2b639436bc3580cf53f0faac26ac10b4b7393c
SELECT user_id ,max(time_stamp) last_stamp  FROM Logins 
where time_stamp  BETWEEN '2020-01-01' AND '2021-01-01'
GROUP by user_id 