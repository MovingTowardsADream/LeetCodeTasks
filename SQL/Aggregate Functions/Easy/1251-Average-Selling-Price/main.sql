SELECT Prices.product_id, COALESCE(ROUND((SUM(price * units) / SUM(units)::decimal), 2), 0) as average_price FROM UnitsSold
RIGHT JOIN Prices ON UnitsSold.product_id = Prices.product_id
WHERE start_date <= purchase_date AND end_date >= purchase_date
GROUP BY Prices.product_id