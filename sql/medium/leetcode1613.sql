-- 找到遗失的ID
-- 递归生成1-n的数字
WITH recursive a as (
        SELECT 1 as n
        UNION ALL
        SELECT n + 1
        from a
        where n < 100
    )
SELECT n as ids
FROM a
where n not in (
        SELECT customer_id
        from customers
    ) AND n <= (
        SELECT max(customer_id)
        from customers
    )