-- 即时食物配送I

SELECT
    round(
        sum(
            IF(
                order_date = customer_pref_delivery_date,
                1,
                0
            )
        ) / (
            select count(1)
            FROM
                Delivery
        ) * 100,
        2
    ) as immediate_percentage
FROM Delivery