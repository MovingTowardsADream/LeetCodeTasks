SELECT a.machine_id, ROUND(CAST(AVG(b.timestamp - a.timestamp) AS numeric), 3) as processing_time
FROM Activity as a
LEFT JOIN Activity as b ON (a.machine_id = b.machine_id)
WHERE a.activity_type = 'start' AND b.activity_type = 'end'
GROUP BY a.machine_id
ORDER BY processing_time DESC