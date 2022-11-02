-- 最大数量高于平均水平的订单

-- 这道题是理解题，本意是  这些订单的最大quantity 要严格大于所有订单的平均quantity。
-- https://leetcode.cn/problems/orders-with-maximum-quantity-above-average/discussion/comments/1249588
SELECT order_id 
FROM OrdersDetails 
GROUP BY order_id
HAVING max(quantity) >  ALL (SELECT AVG(quantity) FROM OrdersDetails GROUP BY order_id)