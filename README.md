# wipe for go

using reflect make alll field and sub struct value to default.

## get start:

`go get github.com/daidai21/wipe`, need `go1.14`+

```go
import (
    "github.com/daidai21/wipe"
    "sync"
)

type Req struct{}

func main() {

    pool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Creating a new Req")
            return new(Req)
        },
    }

    // borrow
    r := pool.Get().(Req)

    // using
    // ...

    // return
    wipe.Wipe(r)
    pool.Put(r)
}
```
