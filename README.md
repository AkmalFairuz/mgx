# MongoX (mgx)

A Golang MongoDB API to simplify the use of MongoDB in your projects.

## Installation

```bash
go get github.com/akmalfairuz/mgx
```

### Example Usage

```go
package main

import (
    "fmt"

    "github.com/akmalfairuz/mgx"
)

func main() {
    db, err := mgx.New("mongodb://localhost:27017", "test")
    if err != nil {
        panic(err)
    }
    r := db.Collection("users").Insert(mgx.M{
        "name": "John Doe",
        "age":  20,
    }, NewInsertOneOptions().BypassDocumentValidation().Comment("John Doe user"))
    if r.Err != nil {
        panic(r.Err)
    }
    fmt.Println(r.InsertedID)
}