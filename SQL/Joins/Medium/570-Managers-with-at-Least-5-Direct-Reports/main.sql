with managers as (
    select managerId from Employee
    group by managerId
    having count(*) >= 5
)
select name from managers
join Employee on managers.managerId = Employee.id