package ipchecker

import (
  "time"
  "io/ioutil"
  "net/http"
)

var (
  quit    chan bool
  results chan string
)

func Check() string {
  resp, err := http.Get("http://ipv4.icanhazip.com")
  if(err != nil) {
    panic(err)
  }
  defer resp.Body.Close()
  ip, _ := ioutil.ReadAll(resp.Body)
  return string(ip)
}

func Poll(every time.Duration) chan string {
  if(quit != nil) {
    return results
  }
  var oldIp string
  results = make(chan string)
  checkIp := func() {
    if ip := Check(); ip != oldIp {
      results <- ip
    }
  }

  quit = make(chan bool)
  go func() {
    go checkIp()
    for {
      select {
      case <- time.Tick(every):
        go checkIp()
      case <-quit:
        break
      }
    }
  }()
  return results
}

func Stop() {
  close(quit)
}
