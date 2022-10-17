-- 首字母大写 Fix Names in a Table

SELECT u.user_id, upper(left(u.name,1)) + lower(substring(u.name,2,datalength(u.name)-1))  as name FROM Users u order by u.user_id