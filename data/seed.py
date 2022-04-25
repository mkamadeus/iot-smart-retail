import psycopg2
import json

with open("secrets.json", "r") as f:
	CONFIG = json.load(f)

conn = psycopg2.connect("dbname=%s user=%s password=%s" % (CONFIG["db_name"], CONFIG["db_user"], CONFIG["db_pass"]))

cur = conn.cursor()

# insert 
cur.execute("SELECT * FROM my_data")

# Retrieve query results
records = cur.fetchall()