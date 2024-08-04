select max(salary) filter (where salary <> (select max(salary) from Employee)) as SecondHighestSalary
from Employee