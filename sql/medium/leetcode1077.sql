-- 项目员工II

-- 这道题可以使用子查询匹配或者是dense_rank窗口函数
-- 题解 https://leetcode.cn/problems/project-employees-iii/solutions/1500891/by-zg104-w7si/
select project_id, employee_id
FROM (
        select
            project_id,
            employee_id,
            dense_rank() over(
                partition by project_id
                order by
                    experience_years desc
            ) as rnk
        from (
                select
                    a.project_id,
                    a.employee_id,
                    b.experience_years
                from project a
                    join employee b using(employee_id)
            ) tmp
    ) as tb
    where rnk=1