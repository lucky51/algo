-- 购买了产品A和产品B却没有购买产品C的顾客

SELECT
    o.customer_id,
    c.customer_name
FROM (
        select distinct *
        from Orders
    ) o
    JOIN Customers c on o.customer_id = c.customer_id
GROUP BY o.customer_id
having
    sum(if(o.product_name = 'A', 1, 0)) = 1
    and sum(if(o.product_name = 'B', 1, 0)) = 1
    and sum(if(o.product_name = 'C', 1, 0)) = 0