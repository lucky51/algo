-- 每个产品在不同商店的价格 Rearrange Products Table  ,行转列的例子



SELECT  product_id , 'store1' store, store1 price  FROM Products where  store1  is not NULL 
UNION 
SELECT  product_id , 'store2' store, store2 price  FROM Products where  store2  is not NULL 
UNION
SELECT  product_id , 'store3'  store, store3 price  FROM Products where  store3  is not NULL 