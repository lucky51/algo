-- 股票的资本损益

SELECT s.stock_name ,SUM(IIF(s.operation='Buy',-s.price,s.price)) capital_gain_loss 
  FROM Stocks s  GROUP BY s.stock_name 