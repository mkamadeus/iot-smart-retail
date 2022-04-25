import random
import pandas as pd
import numpy as np
from faker import Faker
import uuid
import secrets

df = pd.read_excel("transactions.xlsx")

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
	"birhtdate": fake.iso8601(),
	"gender": fake.bothify(letters='MF', text='?')
}, result[:, 0]))

# transaction
txnset = list(map(lambda dt: {
	"id": dt[0],
	"user_id": dt[1],
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
# TODO: input items to db



print(itemset)
print(userset[0])
print(txnset[0])
print(orderset[0])