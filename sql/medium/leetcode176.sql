-- 第二高薪水


SELECT e3.salary AS SecondHighestSalary
FROM 
    (SELECT 2 AS n) tb1
    LEFT JOIN 
    (SELECT e.salary,
        row_number() over(order by e.salary desc) rn
    FROM 
        (SELECT top 2 e1.salary
        FROM Employee e1
        GROUP BY  e1.salary
        ORDER BY  e1.salary desc) AS e ) AS e3
        ON tb1.n =e3.rn


        ----------

  # Write your MySQL query statement below

 select (   select  distinct tb.salary   from (
            select e.salary, dense_rank() over(order by e.salary desc) rn from Employee  e

    )
     tb where tb.rn =2 ) as SecondHighestSalary