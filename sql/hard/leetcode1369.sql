-- 获取最近第二次的活动

-- 使用窗口函数判断条数为1 或者是 row number =2的

SELECT  username,activity,startdate,enddate from(
    SELECT *,row_number() over(partition by username ORDER BY enddate desc) rn ,
    count(1) over(partition by username) count1
) as t WHERE t.rn =2 or count1 =1