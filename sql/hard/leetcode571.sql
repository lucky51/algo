-- 给定数字的频率查询中位数


/*

使用sum over(order by ) 对数字个数进行正序和逆序累计， 当某一数字的 正序和逆序累计 均大于 整个序列的数字个数的一半 时即为中位数， 将最后选定的一个或两个中位数进行求均值即可。

作者：fugue
链接：https://leetcode.cn/problems/find-median-given-frequency-of-numbers/solutions/497993/sum-over-order-by-by-fugue-s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

# Write your MySQL query statement below
select avg(num) median
from (
        select
            num,
            sum(frequency) over(
                order by
                    num
            ) asc_accumu,
            sum(frequency) over(
                order by
                    num desc
            ) desc_accumu
        from numbers
    ) t1, (
        select sum(frequency) total
        from numbers
    ) t2
where
    asc_accumu >= total / 2
    and desc_accumu >= total / 2 
    #作者：fugue
    #链接：https://leetcode.cn/problems/find-median-given-frequency-of-numbers/solutions/497993/sum-over-order-by-by-fugue-s/
    #来源：力扣（LeetCode）
    #著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。