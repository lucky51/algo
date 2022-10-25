-- 平均售价

-- leetcode mysql 能通过， mssql 通过不了小数点没了
SELECT
    tb.product_id,
    ROUND(sum(sales) / sum(tb.units),2)  average_price FROM(
        SELECT
            us.product_id,
            us.units,
           ROUND( p.price *  us.units,2)  sales
        FROM   Prices p
           JOIN  UnitsSold us ON us.product_id = p.product_id
        WHERE
            us.purchase_date  BETWEEN p.start_date
            and p.end_date
    ) as tb GROUP BY tb.product_id