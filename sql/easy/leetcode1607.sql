-- 没有卖出的商家

select seller_name
from seller
where seller_id not in (
        SELECT DISTINCT seller_id
        from orders
        WHERE
            sale_date BETWEEN '2020-01-01' AND '2021-01-01'
    ) ORDER BY seller_name 