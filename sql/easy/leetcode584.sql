-- 寻找推荐人  Find customer Referee

SELECT c.name FROM customer c where c.referee_id <>2 or c.referee_id is NULL