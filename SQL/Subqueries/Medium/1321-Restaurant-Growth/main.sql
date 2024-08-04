select
    visited_on,
    sum(amount) over (rows between 6 preceding and current row) as amount,
    round(sum(amount) over (rows between 6 preceding and current row)/7::decimal, 2) as average_amount
from (
    select visited_on, sum(amount) as amount from Customer
    group by visited_on
    order by visited_on
) as f
offset 6