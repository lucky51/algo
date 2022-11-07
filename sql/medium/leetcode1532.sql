-- 最近的三笔订单

select
    customer_name,
    customer_id,
    order_id,
    order_date
FROM (
        select
            *,
            row_number() over(
                partition by customer_id
                ORDER BY order_date desc
            ) as rn
        from Orders
            JOIN Customers USING(customer_id)
    ) as t
where t.rn < 4
order by
    customer_name,
    customer_id,
    order_date desc