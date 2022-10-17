-- 按日期分组销售产品

SELECT a.sell_date,
        count(distinct a.product ) AS num_sold,
         stuff( 
    (SELECT DISTINCT ',' + product
    FROM Activities a1
    WHERE a1.sell_date = a.sell_date
    ORDER BY  ',' + a1.product FOR xml path('')) ,1,1,'') AS products
FROM Activities a
GROUP BY  a.sell_date 