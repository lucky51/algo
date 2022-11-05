-- 求团队人数

select
    e.employee_id,
    tb.cnt as team_size
from Employee e
    JOIN (
        select
            team_id,
            count(team_id) cnt
        from Employee
        GROUP BY
            team_id
    ) as tb on e.team_id = tb.team_id