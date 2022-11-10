-- 同一天的第一个电话和最后一个电话

with
    a as (
        select caller_id,
        recipient_id,
        call_time
        FROM Calls
        UNION ALL
        SELECT
            recipient_id caller_id,
            caller_id recipient_id,
            call_time
        FROM Calls
    )
SELECT
    DISTINCT a.caller_id user_id
FROM (
        SELECT
            caller_id,
            recipient_id,
            call_time,
            dense_rank() over(
                partition by caller_id,
                date_format(call_time, '%y-%m-%d')
                ORDER BY call_time
            ) rn
        from a
    ) a
    JOIN (
        SELECT
            *,
            dense_rank() over(
                partition by caller_id,
                date_format(call_time, '%y-%m-%d')
                ORDER BY call_time DESC
            ) rn
        from
            a
    ) b ON a.caller_id = b.caller_id
    and a.recipient_id = b.recipient_id
    and a.rn = 1
    and b.rn = 1