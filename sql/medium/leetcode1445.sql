-- 苹果和橘子
-- MYSQL
SELECT
    sale_date,
    SUM(IF(
        fruit = 'apples',
        sold_num,
        -sold_num
    )) as diff
from sales
GROUP BY sale_date
ORDER BY sale_date