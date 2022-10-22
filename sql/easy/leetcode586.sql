-- 订单最多的客户 

SELECT  top 1 customer_number 
 FROM Orders group by customer_number
   order by COUNT(customer_number) desc