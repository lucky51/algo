-- 页面推荐

select
    DISTINCT li.page_id as recommended_page
from (
        select
            DISTINCT if(
                f.user1_id = 1,
                f.user2_id,
                f.user1_id
            ) friendid
        from Friendship f
        where
            f.user1_id = 1
            or f.user2_id = 1
    ) as tb
    join Likes li on tb.friendid = li.user_id
where li.page_id not IN (
        select
            DISTINCT page_id
        from Likes
        where user_id = 1
    )