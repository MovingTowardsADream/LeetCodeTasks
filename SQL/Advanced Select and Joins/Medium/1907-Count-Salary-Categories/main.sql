with res as (
    select
        count(*) filter (where income < 20000) as "Low Salary",
            count(*) filter (where income >= 20000 and income <= 50000) as "Average Salary",
            count(*) filter (where income > 50000) as "High Salary"
    from Accounts
)
-- Преобразование строк в столбцы
SELECT 'High Salary' AS category, f."High Salary" AS accounts_count
FROM res as f
UNION ALL
SELECT 'Low Salary', s."Low Salary"
FROM res as s
UNION ALL
SELECT 'Average Salary', t."Average Salary"
FROM res as t