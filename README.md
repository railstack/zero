# Zero

Zero just provides some helpers for those Gophers prefer the `zero values` than touching the `sql.Null*` types when you have to work with some database tables with nullable fields.

Zero's main idea is using a function `COALESCE` that most popular databases support, as:

* [SQLite](https://sqlite.org/lang_corefunc.html#coalesce)
* [MySQL](https://dev.mysql.com/doc/refman/8.0/en/comparison-operators.html#function_coalesce)
* [PostgreSQL](https://www.postgresql.org/docs/current/static/functions-conditional.html#functions-coalesce-nvl-ifnull)

And I first got the inspiration from [here](https://github.com/go-sql-driver/mysql/issues/34#issuecomment-158391340).

## Usage

Now you must have known what the function `COALESCE` does. So what `zero` does is simple, it helps you write `COALESCE` function calls.

Create a sample table at first:


Let's show how use `zero` to work with a nullable field. Firstly we see the method call style:

```go
import (
    "fmt"
    "github.com/goonr/zero"
)


func main() {
    zr := zero.New("mysql") // or "postgres"
    sql := fmt.Sprintf("SELECT id, name, %v FROM users", zr.Int("age"))
    // here the "sql" will be `SELECT id, name, COALESCE(age, 0) AS age FROM users`
    /* and then do a query with the "sql" %/
}
```

here we create an object using the `New()` function.

Of course you can call the equivalent function without creating an object:

```go
import (
    "fmt"
    "github.com/goonr/zero"
)


func main() {
    sql := fmt.Sprintf("SELECT id, name, %v FROM users", zero.Int("age"))
    // here the "sql" will be `SELECT id, name, COALESCE(age, 0) AS age FROM users`
    /* and then do a query with the "sql" %/
}
```

but the function `DateTime()` is a little different due to it's database independent, so it need to pass the database name as the first parameter:

```go
/* ... */
zero.Time("mysql", "FEILD_NAME")
/* ... */
```

## Functions avaliable

* `String()` is for string type variables in Go
* `Int()` is for all the `int*` typed variables in Go
* `Float()` is for all the `float*` typed variables
* `Bool()` is for the `bool` typed variables
* `Time()` is for the `time.Time` typed variables


And for each of above there's a correnponding `TypeAs()` function that take another parameter as a `AS` alias name, for example:

```go
zero.StringAs("name", "last_name") // will return: COALESCE(name, "") AS last_name
```


