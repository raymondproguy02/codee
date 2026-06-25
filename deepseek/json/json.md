# JSON (JavaScript Object Notation) 
Json: is a text format for storing and exchanging data.
How it looks:

``` json 
{
  "name": "Proguy",
  "age": 20,
  "hobbies": ["smiling", "coding"]
}
```

In Go, JSON is everywhere:
APIs send/recive JSON
Configuration files (config.json)
Storing data in databases
Communication between services

The two Core Operations
Enoding (Go -- JSON): Converting a Go struct into JSON string
Decoding (JSON -- Go): Converting a JSON string into a struct

First example "main.go" same location as this.