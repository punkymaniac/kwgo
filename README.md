# kwgo

kwgo is a Go client library for accessing the [klocwork API](https://docs.roguewave.com/en/klocwork/10-x/klocworkinsightwebapicookbook)


## Installation
```
go get github.com/punkymaniac/kwgo
```

## Usage
```go
import "github.com/punkymaniac/kwgo"

kw := kwgo.NewKwClient("url", "user", "ltoken")

data, err := kw.Version()
if err != nil {
    fmt.Println(err)
}

fmt.Println(data)
```

## LICENSE

MIT

