-- 至少有5名直接下属的经理

select e1.name
from Employee e
    JOIN Employee e1 on e.managerId = e1.id
GROUP BY
    e.managerId,
    e1.name
HAVING count(e.managerid) > 4