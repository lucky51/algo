-- 学生们参加各科测试的次数

select
    s.student_id,
    s.student_name,
    sj._subject_name,
    count(e.subject_name) as attended_exams
from
    Students s -- 笛卡尔积获取生成所有学生的课程
    CROSS JOIN Subjects sj
    LEFT JOIN Examinations e ON s.student_id = e.student_id
    AND sj.subject_name = e.subject_name
GROUP BY
    s.student_id,
    sj.subject_name
ORDER BY s.student_id, sj.subject_name