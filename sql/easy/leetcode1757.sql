-- 可回收且低脂的产品 Recyclable and Low Fat Products

SELECT p.product_id from Products p where p.low_fats='Y' and p.recyclable ='Y'