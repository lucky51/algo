-- 排名靠前的旅行者

SELECT  u.name ,ISNULL(sum(r.distance),0) as travelled_distance   
 FROM Rides r Right join Users u on r.user_id = u.id 
 GROUP BY  r.user_id  ,u.name 
 ORDER BY travelled_distance desc,name ASC