delete from Person as a using Person as b
where a.id > b.id and a.email = b.email