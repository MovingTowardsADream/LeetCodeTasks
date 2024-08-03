with count_users as (
    select count(*) from Users
)
select
    contest_id,
    round(100 * (count(*)::decimal / (SELECT * FROM count_users)), 2) as percentage
from Register as r
join Users on r.user_id = Users.user_id
group by contest_id
order by percentage DESC, contest_id