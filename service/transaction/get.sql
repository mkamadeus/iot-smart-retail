SELECT 
	transactions.id as id, 
	users.id as user_id, 
	jsonb_agg(jsonb_build_object('items', orders.item_id)) as items
FROM transactions
INNER JOIN users ON users.id = transactions.user_id
INNER JOIN orders ON transactions.id = orders.transaction_id
GROUP BY transactions.id, users.id
LIMIT 1;