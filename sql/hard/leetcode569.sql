-- 员工薪水中位数
-- 没看懂
/*
对于一个 奇数 长度数组中的 中位数，大于这个数的数值个数等于小于这个数的数值个数。
根据上述的定义，我们来找一下 [1, 3, 2] 中的中位数。首先 1 不是中位数，因为这个数组有三个元素，却有两个元素 (3，2) 大于 1。3 也不是中位数，因为有两个元素小于 3。对于最后一个 2 来说，
大于 2 和 小于 2 的元素数量是相等的，因此 2 是当前数组的中位数。
总的来说，不管是数组长度是奇是偶，也不管元素是不是唯一，中位数出现的频率一定大于等于大于它的数和小于它的数的绝对值之差。这个规律是这道题的关键，可以通过下面这个搜索条件来过滤
作者：力扣 (LeetCode)
链接：https://leetcode.cn/problems/median-employee-salary/solutions/143337/yuan-gong-xin-shui-zhong-wei-shu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

SELECT
    Employee.Id, Employee.Company, Employee.Salary
FROM
    Employee,
    Employee alias
WHERE
    Employee.Company = alias.Company
GROUP BY Employee.Company , Employee.Salary
HAVING SUM(CASE
    WHEN Employee.Salary = alias.Salary THEN 1
    ELSE 0
END) >= ABS(SUM(SIGN(Employee.Salary - alias.Salary)))
ORDER BY Employee.Id
;

#作者：力扣 (LeetCode)
#链接：https://leetcode.cn/problems/median-employee-salary/solutions/143337/yuan-gong-xin-shui-zhong-wei-shu-by-leetcode/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。