select
    query_name,
    round((sum(rating/position::decimal)/count(*)), 2) as quality,
    round(100 * ((count(*) filter (where rating < 3))/count(*)::decimal), 2) as poor_query_percentage
from Queries
where query_name is not NULL
group by query_name