-- 连续空余座位

SELECT DISTINCT a.seat_id
from Cinema a
    JOIN Cinema b on abs(a.seat_id - b.seat_id) = 1 and a.free = true and b.free = true
ORDER BY a.seat_id