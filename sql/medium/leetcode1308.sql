-- 不同性别每日分数总计
-- 这道题的意思是，统计某一天的某一个性别累加分数，不仅仅是那一天的
-- 用聚合函数sum，作为窗口函数 所代表的意思是截止某一天所累加的结果是多少，符合题意。
-- 还有一种做法是使用left join ，然后判断日期，在聚合到性别，比较复杂不容易理解
SELECT gender,day,sum(score_points) over (partition by gender ORDER BY day) total
from scores