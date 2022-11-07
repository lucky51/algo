-- 坚定的友谊
-- 比较难理解
# Write your MySQL query statement below
with t as (
    (select user1_id user_id,user2_id friend_id
    from Friendship
    union all
    select user2_id,user1_id
    from Friendship)
    order by 1,2
    )
select t1.user_id user1_id
    ,t2.user_id user2_id
    ,count(distinct t1.friend_id) common_friend 
from t t1,t t2
where (t1.user_id,t2.user_id) in (select * from t)
and t1.friend_id = t2.friend_id
and t1.user_id < t2.user_id
group by 1,2
having count(distinct t1.friend_id) >=3

#作者：Lyon
#链接：https://leetcode.cn/problems/strong-friendship/solutions/907157/ming-bie-ming-ye-shi-yi-chong-xue-wen-qi-vhn5/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。