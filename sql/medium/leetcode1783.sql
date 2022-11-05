-- 大满贯数量

# Write your MySQL query statement below 行转列在计算
select 
    p.player_id,p.player_name,count(*) as grand_slams_count 
from
    players as p inner join 
    (select wimbledon from championships
    union all
    select fr_open from championships
    union all 
    select us_open from championships
    union all
    select au_open from championships) as t
on 
    p.player_id = t.wimbledon
group by 
    p.player_id
