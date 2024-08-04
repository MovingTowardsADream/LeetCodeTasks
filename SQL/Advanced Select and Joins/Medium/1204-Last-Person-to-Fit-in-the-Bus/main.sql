select person_name from (
    select *, row_number() over (order by sum_w desc) as rn from (
        select
            *,
            sum(weight) over (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as sum_w
        from (select * from Queue order by turn) as f
    ) as s
    where sum_w <= 1000
) as t
where rn = 1