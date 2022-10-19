-- 进店却未进行过交易的顾客

SELECT   v.customer_id  , count( v.customer_id ) as count_no_trans   FROM  Visits v left join  Transactions t on v.visit_id = t.visit_id  
where t.visit_id is NULL  group by v.customer_id