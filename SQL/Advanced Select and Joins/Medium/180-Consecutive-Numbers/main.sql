WITH cte AS (
    SELECT
        id,
        num,
        max(id) OVER (ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS max_id,
            min(id) OVER (ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS min_id,
            max(num) OVER (ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS max_num,
            min(num) OVER (ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS min_num
    FROM (select * from Logs order by id) as a
)
SELECT distinct(num) AS ConsecutiveNums
FROM cte
WHERE max_num = min_num AND max_id = min_id + 2;