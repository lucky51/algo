-- 合作过至少三次的演员和导演

SELECT  actor_id ,director_id FROM ActorDirector GROUP BY actor_id,director_id 
HAVING count(actor_id) >2