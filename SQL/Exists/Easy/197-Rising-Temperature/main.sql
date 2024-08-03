SELECT id FROM Weather as today
WHERE EXISTS (
    SELECT * FROM Weather as yesterday
    WHERE today.recordDate = yesterday.recordDate + 1 AND today.temperature > yesterday.temperature
)