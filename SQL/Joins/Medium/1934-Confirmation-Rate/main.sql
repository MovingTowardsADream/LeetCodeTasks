select
    user_id,
    round((count(*) filter (where action = 'confirmed'))::decimal / count(*), 2) as confirmation_rate
from Signups
left join Confirmations using(user_id)
group by user_id