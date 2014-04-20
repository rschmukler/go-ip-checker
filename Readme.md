#go-ip-checker

Public ip checker for go

Ip checking generously provided by [http://icanhazip.com](http://icanhazip.com)

## Example Usage

#### Check

```go
ip := ipchecker.Check()
fmt.Printf("My ip is: %s\n", ip)
```

#### Poll/Stop

```go
for {
  checkEvery := time.Duration(30) * time.Second
  select {
    case ip := <- ipchecker.Poll(checkEvery):
      fmt.Printf("Ip changed to: %s\n", ip)
    case <- time.After(time.Duration(30) * time.Second):
      ipchecker.Stop()
      break
  }
}
```

## Full Documentation

You can see the godocs
[here](http://godoc.org/github.com/rschmukler/go-ip-checker).
