-- 查询活跃业务

SELECT business_id
FROM Events AS e
JOIN (
    SELECT event_type, AVG(occurences) AS eventAvg
    FROM Events
    GROUP BY event_type
) AS e1 
ON e.event_type = e1.event_type
WHERE e.occurences > e1.eventAvg
GROUP BY business_id
HAVING COUNT(*) >= 2

#作者：力扣官方题解
#链接：https://leetcode.cn/problems/active-businesses/solutions/141909/cha-xun-huo-yue-ye-wu-by-leetcode-solution/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。