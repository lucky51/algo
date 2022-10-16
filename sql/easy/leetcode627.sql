-- 变更性别 ,这里使用了 IFF函数，也可以使用 when case 

UPDATE  Salary SET sex = IIF(sex='m','f','m')