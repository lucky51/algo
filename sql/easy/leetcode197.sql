-- 上升的温度

SELECT  w2.id  FROM Weather w1 left join Weather w2 on w1.recordDate=dateadd(dd,-1,w2.recordDate) where w2.temperature > w1.temperature