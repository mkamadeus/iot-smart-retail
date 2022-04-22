SELECT 
	transactions.id as id, 
	orders.id as order_id, 
	jsonb_agg(jsonb_build_object('items', order_items.order_id)) as items
FROM transactions
INNER JOIN orders ON orders.id = transactions.order_id
INNER JOIN order_items ON orders.id = order_items.order_id
GROUP BY transactions.id, orders.id;

8a6f1d1e-7bae-4135-9183-add34df63a61