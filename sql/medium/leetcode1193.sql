-- 每月交易I

SELECT
    FORMAT(trans_date, 'yyyy-MM') as month,
    country,
    count(*) trans_count,
    COUNT(IIF(state = 'approved', 1, NULL)) as approved_count,
    SUM(amount) as trans_total_amount,
    SUM(
        IIF(state = 'approved', amount, 0)
    ) as approved_total_amount
from Transactions
GROUP BY
    FORMAT(trans_date, 'yyyy-MM'),
    country