select name, bonus from Employee as e
left join Bonus as b on e.empId = b.empId
where bonus < 1000 OR bonus is null