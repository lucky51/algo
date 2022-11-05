-- 每位学生的最高成绩

select
    t.student_id,
    t.course_id,
    t.grade
from (
        select
            student_id,
            course_id,
            grade,
            row_number() over(
                partition by student_id
                ORDER BY
                    grade desc,
                    course_id
            ) as rn
        from Enrollments
    ) as t
where t.rn = 1
ORDER BY t.student_id