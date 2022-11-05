-- 计算布尔表达式的值

SELECT
    e.left_operand,
    e.operator,
    e.right_operand, (
        case
            when e.operator = '=' then if(
                v.value = v1.value,
                'true',
                'false'
            )
            when e.operator = '>' then if(
                v.value > v1.value,
                'true',
                'false'
            )
            else if(
                v.value < v1.value,
                'true',
                'false'
            )
        end
    ) as value
FROM Expressions e
    JOIN Variables v on e.left_operand = v.name
    JOIN Variables v1 on e.right_operand = v1.name