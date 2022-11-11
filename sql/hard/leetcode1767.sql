-- 寻找没有被执行的任务对


with recursive a as (
    SELECT task_id,1 as subtask_id,subtasks_count FROM Tasks
    UNION ALL
    SELECT t.task_id,t.subtask_id+1 as c,t.subtasks_count FROM  a t WHERE  subtask_id <t.subtasks_count
)
select * from  Executed WHERE (task_id ,subtask_id) not IN (
    SELECT task_id,subtask_id FROM a
)