-- 统计各专业学生人数

select
    Department.dept_name,
    count(Student.dept_id) as student_number
from Department
    LEFT JOIN Student on Student.dept_id = Department.dept_id
GROUP BY Department.dept_id
ORDER BY student_number desc