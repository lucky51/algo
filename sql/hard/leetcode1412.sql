-- 查找成绩处于中游的学生




-- 这道题说的不是所有测验的中游，而是具体某一次测验
-- 按照题解上说，可以按照score升序，降序各自排一次

SELECT  t.student_id,s.student_name FROM (
    SELECT  *,if(dense_rank() over(partition by exam_id ORDER BY score DESC)=1,1,0) d1,
    if(dense_rank() over(partition by exam_id ORDER BY score)=1,1,0) d2  
    FROM exam
) as t JOIN student s USING(student_id)
GROUP BY t.student_id
HAVING sum(d1)=0 and sum(d2)=0
ORDER BY t.student_id