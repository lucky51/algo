-- 查询结果的质量和占比

SELECT
    query_name,
    round(
        sum(rating / position) / count(query_name),
        2
    ) as quality,
    round(
        sum(if(rating < 3, 1, 0)) / count(query_name) * 100,
        2
    ) as poor_query_percentage
from Queries
group by query_name