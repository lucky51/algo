-- 文章浏览I

SELECT  DISTINCT v1.author_id as id FROM Views v1  where v1.author_id  = v1.viewer_id   order by v1.author_id