-- 按年度列出销售总额


-- 这道题条件判断比较复杂，先略过
select
    s.product_id,
    p.product_name,
    y.year report_year,
    s.average_daily_sales * (
        if(
            year(s.period_end) > y.year,
            y.days_of_year,
            dayofyear(s.period_end)
        ) - if(
            year(s.period_start) < y.year,
            1,
            dayofyear(s.period_start)
        ) + 1
    ) total_amount
from Sales s
    inner join (
        select
            '2018' year,
            365 days_of_year
        union all
        select
            '2019' year,
            365 days_of_year
        union all
        select
            '2020' year,
            366 days_of_year
    ) y on year(s.period_start) <= y.year
    and year(s.period_end) >= y.year
    inner join Product p on p.product_id = s.product_id
order by
    s.product_id,
    y.year #作者：jcj228675
    #链接：https://leetcode.cn/problems/total-sales-amount-by-year/solutions/818507/9xing-dai-ma-fei-unionfang-fa-by-jcj2286-hvg6/
    #来源：力扣（LeetCode）
    #著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

    # https://leetcode.cn/problems/total-sales-amount-by-year/solutions/1667582/fei-di-gui-by-luckysi-kai-xin-6yx7/?orderBy=most_votes