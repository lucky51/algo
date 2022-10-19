-- 175 组合两个表

SELECT p.firstName ,p.lastName,a.city,a.state FROM Person p left join Address a on p.personId = a.personId 