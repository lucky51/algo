-- 每位顾客最经常订购的商品

SELECT Orders.customer_id  FROM (
    select customer_id ,product_id ,count(product_id) as cnt FROM Orders GROUP BY customer_id ,product_id
) as t
JOIN Customers ON t.customer_id=Customers.customer_id
JOIN Products on t.product_id = Products.product_id
