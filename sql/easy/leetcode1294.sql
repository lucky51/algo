-- 不同国家的天气类型

# Write your MySQL query statement below
SELECT
    Countries.country_name, (
        case
            when avg(Weather.weather_state) <= 15 then 'Cold'
            when avg(Weather.weather_state) >= 25 then 'Hot'
            else 'Warm'
        end
    ) as weather_type
FROM Weather
    JOIN Countries ON Weather.country_id = Countries.country_id
WHERE
    Weather.day BETWEEN '2019-11-01' AND '2019-11-30'
GROUP BY Weather.country_id