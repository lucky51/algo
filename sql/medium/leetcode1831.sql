-- 每天的最大交易

select transaction_id
FROM (
        SELECT
            *,
            dense_rank() over(
                partition by date_format(day, '%y%m%d')
                ORDER BY amount desc
            ) as rn
        from Transactions
    ) as t
where t.rn = 1
ORDER BY transaction_id