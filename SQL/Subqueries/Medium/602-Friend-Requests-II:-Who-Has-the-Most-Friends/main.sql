with tmp as (
    select id, coalesce(f.num, 0) + coalesce(s.num, 0) as num from (
       select accepter_id as id, count(*) as num from RequestAccepted
       group by accepter_id
    ) as f
    full join (
        select requester_id as id, count(*) as num from RequestAccepted
        group by requester_id
    ) as s using (id)
)
select * from tmp where num = (select max(num) from tmp)