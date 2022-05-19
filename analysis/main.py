from fastapi import FastAPI
from fastapi.responses import JSONResponse
from mlxtend.frequent_patterns import fpgrowth
from mlxtend.preprocessing import TransactionEncoder
from pydantic import BaseSettings
import pandas as pd
import requests
import json

class Settings(BaseSettings):
    service_url: str = "http://localhost:8080"

app = FastAPI()
settings = Settings()

@app.get("/")
async def root():
    response = requests.get(f'{settings.service_url}/transactions')
    txns = json.loads(response.text)
    
    orders = []
    for txn in txns:
        order = list(map(lambda item: item["name"], txn["items"]))
        orders.append(order)

    te = TransactionEncoder()
    te_ary = te.fit(orders).transform(orders)
    txn_df = pd.DataFrame(te_ary, columns=te.columns_)

    result = fpgrowth(txn_df, min_support=0.3, use_colnames=True, max_len=2)
    
    basket = []
    for _, row in result.iterrows():

        basket.append({
            "support": row["support"],
            "itemsets": list(row["itemsets"])
        })

    return JSONResponse(content=basket)

    
