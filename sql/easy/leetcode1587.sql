-- 银行账户概要II

--mysql
select u.name,ifnull(count(t.amount),0) balance from Users u  LEFT join  Transactions t
on u.account = t.account GROUP BY  u.account HAVING balance >10000

--tsql

select u.name,isnull(count(t.amount),0) balance from Users u  LEFT join  Transactions t
on u.account = t.account GROUP BY  u.account HAVING isnull(count(t.amount),0) >10000