# basic fastapi CRUD operations
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI()

# --- Model ---
class Item(BaseModel):
    id: int
    name: str
    price: float
    in_stock: bool = True

# --- "Database" (just a dict for now) ---
db: dict[int, Item] = {}

# --- CREATE ---
@app.post("/items", response_model=Item)
def create_item(item: Item):
    if item.id in db:
        raise HTTPException(status_code=400, detail="Item already exists")
    db[item.id] = item
    return item

# --- READ (all) ---
@app.get("/items", response_model=list[Item])
def list_items():
    return list(db.values())

# --- READ (one) ---
@app.get("/items/{item_id}", response_model=Item)
def get_item(item_id: int):
    if item_id not in db:
        raise HTTPException(status_code=404, detail="Item not found")
    return db[item_id]

# --- UPDATE ---
@app.put("/items/{item_id}", response_model=Item)
def update_item(item_id: int, updated: Item):
    if item_id not in db:
        raise HTTPException(status_code=404, detail="Item not found")
    db[item_id] = updated
    return updated

# --- DELETE ---
@app.delete("/items/{item_id}")
def delete_item(item_id: int):
    if item_id not in db:
        raise HTTPException(status_code=404, detail="Item not found")
    del db[item_id]
    return {"message": "Item deleted"}
