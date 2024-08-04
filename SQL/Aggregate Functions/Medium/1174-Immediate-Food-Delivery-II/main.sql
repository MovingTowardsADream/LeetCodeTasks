select
    round(100 * count(*) filter (where s.order_date = f.customer_pref_delivery_date) / count(*)::decimal, 2) as immediate_percentage
from Delivery as f
join (
    select customer_id, min(order_date) as order_date from Delivery
    group by customer_id
) as s using (customer_id, order_date)