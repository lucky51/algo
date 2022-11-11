-- 报告系统状态的连续日期
-- 官方题解一 mysql 变量法没具体研究过，看不懂
SELECT period_state, MIN(date) as start_date, MAX(date) as end_date
FROM (
    SELECT
        success_date AS date,
        "succeeded" AS period_state,
        IF(DATEDIFF(@pre_date, @pre_date := success_date) = -1, @id, @id := @id+1) AS id 
    FROM Succeeded, (SELECT @id := 0, @pre_date := NULL) AS temp
    UNION
    SELECT
        fail_date AS date,
        "failed" AS period_state,
        IF(DATEDIFF(@pre_date, @pre_date := fail_date) = -1, @id, @id := @id+1) AS id 
    FROM Failed, (SELECT @id := 0, @pre_date := NULL) AS temp
) T  WHERE date BETWEEN "2019-01-01" AND "2019-12-31"
GROUP BY T.id
ORDER BY start_date ASC

#作者：力扣官方题解
#链接：https://leetcode.cn/problems/report-contiguous-dates/solutions/305366/bao-gao-xi-tong-zhuang-tai-de-lian-xu-ri-qi-by-lee/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

-- 
select period_state,start_date,end_date 
from(
select type as period_state,diff, min(date) as start_date, max(date) as end_date
from
(
    select type, date, subdate(date,row_number()over(partition by type order by date)) as diff,
    row_number()over(partition by type order by date) as rn
    from
    (
        select 'failed' as type, fail_date as date from Failed
        union all
        select 'succeeded' as type, success_date as date from Succeeded
    ) a
)a
where date between '2019-01-01' and '2019-12-31'
group by type,diff
) t
order by start_date

#作者：leo
#链接：https://leetcode.cn/problems/report-contiguous-dates/solutions/1011320/fen-bu-zou-jie-jue-jian-dan-yi-dong-by-l-0ch5/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。