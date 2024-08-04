with count_products as (
    select count(*) as num from Product
)
select customer_id from Customer
group by customer_id
having count(distinct(product_key)) = (select num from count_products)