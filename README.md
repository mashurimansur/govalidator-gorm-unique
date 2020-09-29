# govalidator-gorm-unique

Unique validator extensions for [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator). Inspired by Laravel's unique validation rule & [go-unique-validator](https://github.com/ramadani/go-unique-validator).

## Installation

Before use this validator, you need to install [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator) first, and then install this package.

```cmd
go get github.com/mashurimansur/govalidator-gorm-unique
```

## Usage

Import this package to your code

```go
import uniquevalidator "github.com/mashurimansur/govalidator-gorm-unique"
```

Create db instance

```go
db, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname?parseTime=true")
if err != nil {
  log.Fatal(err)
}
defer db.Close()
```

Add as custom rule to govalidator

```go
uniqueRule := uniquevalidator.NewUniqueRule(db, "unique")
govalidator.AddCustomRule("unique", uniqueRule.Rule)
```

#### Example Rule

Format: `unique:table,column,except,idColumn`

To check if attribute is unique:

```go
rules := govalidator.MapData{
	"email": []string{"required", "email", "unique:users,email"},
}
```

Forcing A Unique Rule To Ignore A Given ID:

```go
rules := govalidator.MapData{
	"email": []string{"required", "email", "unique:users,email,id,123"},
}
```

### Example Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uniquevalidator "github.com/mashurimansur/govalidator-gorm-unique"
	"github.com/thedevsaddam/govalidator"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"email": []string{"required", "email", "unique:users,email"},
	}

	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		RequiredDefault: true,  // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func main() {
	db, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	uniqueRule := uniquevalidator.NewUniqueRule(db, "unique")
	govalidator.AddCustomRule("unique", uniqueRule.Rule)

	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}
```

Send request to the server using curl or postman: `curl GET "http://localhost:9000?email=your@email.com"`

**Response**

```json
{
    "validationError": {
        "email": [
            "The email has already been taken"
        ]
    }
}
```

## **License**
The **govalidator-gorm-unique** is an open-source software licensed under the [MIT License](LICENSE).
