-- 两人之间的通话次数

SELECT 
 c.person1,c.person2,
 count(*) call_count,
 sum(c.duration) total_duration
 from  (
    SELECT 
        IIF(from_id <to_id,from_id,to_id) person1,
        IIF(from_id<to_id,to_id,from_id) person2,
        duration
     from Calls  
) c

GROUP BY c.person1,c.person2