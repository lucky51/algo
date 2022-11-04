-- 找到连续区间的开始和结束数字
-- 这里使用 rownumber进行编号，如果连续相同的话和log_id的差值是一样的，那么就可以进行分组，每一组中的最小值和最大值就是我们想要的答案，即start_id ,end_id
SELECT  min(log_id) start_id, max(log_id) end_id FROM
(
   SELECT log_id, log_id - row_number() over() as   diff  FROM Logs
)  t
GROUP BY diff
ORDER BY diff

