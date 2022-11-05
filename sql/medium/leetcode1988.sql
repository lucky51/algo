-- 找出每所学校的最低分数要求
-- TODO： 没读懂的题目

select school_id, ifnull(min(score),-1) as score 
from schools left join exam on capacity >= student_count
group by school_id