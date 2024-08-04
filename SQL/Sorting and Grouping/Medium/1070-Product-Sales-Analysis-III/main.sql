select a.product_id, a.year as first_year, quantity, price from Sales as a
join (
    select product_id, min(year) as year from Sales
    group by product_id
) as b using(product_id, year)