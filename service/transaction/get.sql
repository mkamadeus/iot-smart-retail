SELECT 
	transactions.id as id, 
	users.id as user_id, 
	jsonb_agg(jsonb_build_object(
		'id', orders.item_id,
		'name', items.name,
		'count', orders.count
	)) as items
FROM transactions
INNER JOIN users ON users.id = transactions.user_id
INNER JOIN orders ON transactions.id = orders.transaction_id
INNER JOIN items ON orders.item_id = items.id
WHERE transactions.transaction_id = ?
GROUP BY transactions.id, users.id;