-- 销售员


SELECT name FROM SalesPerson where name not in (
SELECT DISTINCT  sp.name FROM  SalesPerson sp LEFT JOIN  Orders  o  on sp.sales_id  = o.sales_id  LEFT  JOIN Company  c on o.com_id  = c.com_id
where c.name  = 'RED' 

)
