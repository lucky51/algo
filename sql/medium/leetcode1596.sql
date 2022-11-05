-- 每位顾客最经常订购的商品


SELECT o.customer_id,o.product_id,p.product_name from products p,

(
    SELECT customer_id,product_id,rank() over(partition by customer_id order by count(1) desc) rk
FROM orders GROUP BY customer_id,product_id
) as o
where o.rk =1 and p.product_id = o.product_id