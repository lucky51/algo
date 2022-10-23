-- 游戏玩法分析 I

SELECT player_id,MIN(event_date) as  first_login  FROM Activity  GROUP by player_id 