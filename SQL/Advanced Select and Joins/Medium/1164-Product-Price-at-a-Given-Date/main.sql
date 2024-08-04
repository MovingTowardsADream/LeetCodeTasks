select product_id, coalesce(price, 10) as price from (select distinct(product_id) from Products) as f
left join (
    select product_id, price from (
        select
            product_id,
            new_price as price, row_number() over (partition by product_id order by change_date desc) as first_row
        from Products
        where change_date <= '2019-08-16'
    ) as t
    where first_row = 1
) as s using (product_id)