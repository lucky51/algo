-- 第N高的薪水

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
        select distinct t.salary from  (
            select salary,dense_rank() over(ORDER by salary desc) rn from Employee
        ) as t where t.rn =N
      
  );
END