-- 页面推荐



select * from (
    select  DISTINCT  if(f.user1_id=1,f.user2_id,f.user1_id)  friendid
from Friendship f 
) as tb  join  Likes li on tb.friendid = li.user_id
where  li.page_id not IN (
 select DISTINCT  page_id from Likes where user_id =1
) 