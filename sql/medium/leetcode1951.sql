-- 查询具有最多共同关注者的所有两两结对组

-- 两两结对组 可以考虑 ，自连接 笛卡尔积，然后找到 关注者相同的

SELECT t.user1_id, t.user2_id
FROM (
        SELECT
            r1.user_id as user1_id,
            r2.user_id as user2_id,
            rank() over(
                ORDER BY count(1) desc
            ) rn
        from relations as r1
            inner JOIN relations as r2 on r1.follower_id = r2.follower_id
        where
            r1.user_id < r2.user_id
        GROUP BY
            r1.user_id,
            r2.user_id
    ) as t
where t.rn = 1