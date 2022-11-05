-- 查找重复的电子邮箱

SELECT Email FROM Person GROUP BY Email
HAVING Count(Email) >1