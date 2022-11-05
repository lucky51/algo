-- 市场分析I

SELECT   u.user_id buyer_id,u.join_date,isnull(UserBuy.cnt,0) orders_in_2019  
FROM Users u
LEFT join (
    select buyer_id,count(order_id) cnt from Orders where order_date BETWEEN '2019-01-01' AND '2020-01-01'
     GROUP BY buyer_id

)UserBuy  on u.user_id = UserBuy.buyer_id
