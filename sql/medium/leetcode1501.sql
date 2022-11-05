-- 可以放心投资的国家

select cy.name as country
from Calls c
    JOIN Person p on c.caller_id = p.id or c.callee_id = p.id
    JOIN Country cy on left(p.phone_number, 3) = cy.country_code
GROUP BY cy.name
HAVING avg(c.duration) > avg( (
            select avg(duration)
            from Calls
        )
    )