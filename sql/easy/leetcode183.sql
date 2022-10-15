-- 未曾下单的客户  Customers Who Never Order

SELECT c.name as Customers FROM Customers c LEFT JOIN 
Orders o on c.id = o.customerId where o.customerId is null 