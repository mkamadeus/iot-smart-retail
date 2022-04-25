import random
import pandas as pd
import psycopg2
import json
from faker import Faker
import uuid
import secrets

df = pd.read_excel("transactions.xlsx")
print(df.head())

#1
customer_ids = df['Customer_ID'].unique()
customer_dictionary = {cid: str(uuid.uuid4()) for cid in customer_ids}
df['Customer_ID'] = df['Customer_ID'].map(customer_dictionary)

#2
df['No_Transaksi'] = df['No_Transaksi'].apply(lambda x: str(uuid.uuid4()))

#3
df['Tanggal'] = pd.to_datetime(df['Tanggal']).dt.strftime('%Y-%m-%dT%H:%M:%SZ')

#4
df['Items'] = df['Items'].apply(lambda x: x.split(', '))

result = df.to_numpy()
fake = Faker()

# items
itemset = []
for items in result[:, 3]:
	itemset.extend(items)
itemset = list(set(itemset))
itemset = list(map(lambda item: {"id": str(uuid.uuid4()), "name": item}, itemset))
itemset = {s["name"]: s for s in itemset}

# users
userset = list(map(lambda user: {
	"id": user, 
	"name": fake.name(),
	"card_id": secrets.token_hex(12),
	"birthdate": fake.iso8601(),
	"gender": fake.bothify(letters='MF', text='?')
}, result[:, 0]))

# transaction
print(result[0])
txnset = list(map(lambda dt: {
	"id": dt[1],
	"user_id": dt[0],
}, result[:, 0:2]))

# order
orderset = []
for row in result:
	rowresult = []
	for item in row[3]:
		item_id = itemset[item]["id"]
		rowresult.append({
			"transaction_id": row[1], 
			"item_id": item_id,
			"count": random.randint(1,100)
		})
	orderset.append(rowresult)

with open("secrets.json", "r") as f:
	CONFIG = json.load(f)

conn = psycopg2.connect("dbname=%s user=%s password=%s host=%s" % (CONFIG["db_name"], CONFIG["db_user"], CONFIG["db_pass"], CONFIG["db_host"]))

cur = conn.cursor()
try:

	# users
	print('users')
	for u in userset:
		cur.execute("INSERT INTO users (id, name, card_id, birthdate, gender) VALUES (%s, %s, %s, %s, %s)", (u['id'], u['name'], u['card_id'], u['birthdate'], u['gender']))

	conn.commit()

	# items
	print('items')
	for i in itemset.values():
		cur.execute("INSERT INTO items (id, name) VALUES (%s, %s)", (i['id'], i['name']))

	conn.commit()

	# transactions
	print('txns')
	for t in txnset:
		cur.execute("INSERT INTO transactions (id, user_id) VALUES (%s, %s)", (t['id'], t['user_id']))

	conn.commit()

	# orders
	print('orders')
	for o in orderset:
		for oi in o:
			cur.execute("INSERT INTO orders (transaction_id, item_id, count) VALUES (%s, %s, %s)", (oi['transaction_id'], oi['item_id'], oi['count']))

	conn.commit()

except Exception as e:
	print(e)
finally:
	cur.close()