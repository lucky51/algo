-- 树节点

SELECT  

select distinct self.id, 
case 
     --没有根节点就是父节点
    when self.p_id is null then 'Root'
    -- 有孩子结点
    when child.id is not null then 'Inner'
    when child.id is null then 'Leaf'
end  type
 from tree self left join tree child on self.id = child.p_id;