-- 餐馆营业额变化增长

# Write your MySQL query statement below
SELECT
	a.visited_on,
	sum( b.amount ) AS amount,
	round(sum( b.amount ) / 7, 2 ) AS average_amount 
FROM
	( SELECT DISTINCT visited_on FROM customer ) a JOIN customer b 
 	ON datediff( a.visited_on, b.visited_on ) BETWEEN 0 AND 6 
WHERE
	a.visited_on >= (SELECT min( visited_on ) FROM customer ) + 6   --大于7的，
GROUP BY
	a.visited_on

#作者：小数志
#链接：https://leetcode.cn/problems/restaurant-growth/solutions/213501/dai-ma-bu-pai-ban-de-ti-jie-shui-neng-kan-de-dong-/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


select 
distinct a.visited_on,sum(b.amount) as amount,round(sum(b.amount)/7,2) as average_amount 
from Customer a
join Customer b
on datediff(a.visited_on,b.visited_on) <= 6 and datediff(a.visited_on,b.visited_on) >= 0
group by a.visited_on,a.customer_id     
having  count(distinct b.visited_on) = 7
order by visited_on 

#作者：居右
#链接：https://leetcode.cn/problems/restaurant-growth/solutions/1101930/can-guan-ying-ye-e-bian-hua-zeng-chang-c-vplu/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
#在group by的时候，由于前面的筛选条件是DATEDIFF <= 6 且>= 0，出来的表依然含有2019-01-01到06这些前面不足6天的情况。
#having count distinct相当于保证每个group by的小分组中都有7个数据，也就是保证从2019-01-07开始统计，把前面的数据全部排除了
#