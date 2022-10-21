-- 求关注者的数量

SELECT   f.user_id ,count(DISTINCT f.follower_id ) as followers_count FROM Followers f GROUP BY f.user_id  order by f.user_id 
