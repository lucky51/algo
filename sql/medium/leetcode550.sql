-- 游戏玩法分析IV
-- 注意题目的意思是首次登陆，要考虑只有首次之后连续的才算
--  3-1，3-3，3-4  比如这种情况就不符合，所以后边必须要查每个player_id的最小值筛选一下结果
SELECT
    ROUND(
        count(DISTINCT a.player_id) / (
            select
                count(DISTINCT player_id)
            FROM Activity
        ), 2
    ) AS fraction
FROM Activity a
    JOIN Activity b ON a.player_id = b.player_id and a.event_date = b.event_date -1
    where (a.player_id,a.event_date) in (
        select player_id,min(event_date) FROM Activity
        GROUP BY player_id
    )
