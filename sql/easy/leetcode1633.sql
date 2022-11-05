-- 各赛事的用户注册率

-- mssql 中提交不过

SELECT
    contest_id,
    round(
        count(user_id) / (
            SELECT count(1)
            from
                users
        ) * 100,
        2
    ) as percentage
FROM Register
GROUP BY contest_id
order by percentage desc, contest_id