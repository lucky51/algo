-- 矩形面积


-- 没看懂这道题
SELECT
    t1.id AS 'p1',
    t2.id AS 'p2',
    ABS(t2.x_value - t1.x_value) * ABS(t2.y_value - t1.y_value) AS 'area'
FROM
    Points AS t1
INNER JOIN Points AS t2 ON t1.id < t2.id
WHERE (t2.x_value - t1.x_value) != 0
AND (t2.y_value - t1.y_value) != 0
ORDER BY area DESC, t1.id, t2.id
