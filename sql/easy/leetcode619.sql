-- 只出现一次的最大数字

select max(num) from (
    select num from MyNumbers GROUP BY num HAVING count(num) =1 
)  as t  --派生出来的临时表必须具有别名。