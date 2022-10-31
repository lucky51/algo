-- 应该被禁止的Leetflex账户

SELECT DISTINCT a.account_id
FROM LogInfo a
    JOIN LogInfo b on a.account_id = b.account_Id and a.ip_address <> b.ip_address and ( (
            a.login BETWEEN b.login and b.logout
        ) OR (
            a.logout BETWEEN b.login and b.logout
        )
    )