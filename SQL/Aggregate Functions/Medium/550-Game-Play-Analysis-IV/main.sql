select round(count(*)::decimal / (select count(distinct(player_id)) from Activity), 2) as fraction from (
    select player_id, min(event_date) as event_date from Activity
    group by player_id
) f
where exists(
    select * from Activity as s
    where f.player_id = s.player_id and s.event_date = f.event_date + 1
)