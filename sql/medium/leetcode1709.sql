-- 访问日期之间最大的空档期

SELECT
    user_id,
    max(
        datediff(next_day, visit_date)
    ) as biggest_window
FROM(
        SELECT
            user_id,
            visit_date,
            lead(visit_date, 1, '2021-1-1') over(
                partition by user_id
                ORDER BY visit_date
            ) next_day
        from UserVisits
    ) tmp
GROUP BY user_id
ORDER BY user_id