-- 买下所有产品的客户

select Customer.customer_id
from Customer
GROUP BY Customer.customer_id
HAVING
    count(DISTINCT Customer.product_key) = (
        select
            count(DISTINCT product_key)
        from product
    )