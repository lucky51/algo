-- 仓库经理

SELECT
    w.name WAREHOUSE_NAME,
    SUM(
        p.Width * p.Length * p.Height * w.units
    ) VOLUME
FROM Warehouse w
    JOIN Products p ON w.product_id = p.product_id
Group by w.name