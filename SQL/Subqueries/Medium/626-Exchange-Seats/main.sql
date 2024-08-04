select
    case when id % 2 = 0 then id-1
         when row_num = 1 and (select count(*) as l from Seat) % 2 = 1 then id
         else id+1
        end as id,
    student
from (
         select *, ROW_NUMBER() OVER (ORDER BY id desc) AS row_num from Seat
     ) as f
order by id