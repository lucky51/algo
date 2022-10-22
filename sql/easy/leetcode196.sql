-- 删除重复的电子邮箱 

delete from Person  where id in (
  select id from ( select   p.id,ROW_NUMBER() OVER(PARTITION BY p.email order by p.id ) rn from Person p  ) as tb where tb.rn !=1

)

