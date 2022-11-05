-- 每件商品的最新订单

SELECT
    Products.product_name,
    Orders.product_id,
    Orders.order_id,
    Orders.order_date
from Orders
    JOIN Products ON Products.product_id = Orders.product_id
where (
        Orders.product_id,
        Orders.order_date
    ) in (
        select
            product_id,
            max(order_date)
        FROM Orders
        GROUP BY product_id
    )
    ORDER BY Products.product_name,Orders.product_id,Orders.order_id