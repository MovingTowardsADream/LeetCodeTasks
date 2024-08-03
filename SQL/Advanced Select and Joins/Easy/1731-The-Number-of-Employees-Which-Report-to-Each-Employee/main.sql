with group_by_reports as (
    select reports_to, avg(age)::integer as average_age, count(*) as reports_count from Employees
    group by reports_to
)
select employee_id, name, reports_count, average_age from group_by_reports as gbr
join Employees as e on e.employee_id = gbr.reports_to
order by employee_id