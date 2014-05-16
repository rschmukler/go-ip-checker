package ipchecker

import (
  "strings"
  "time"
  "io/ioutil"
  "net/http"
)

type IPChecker struct {
  C chan string
  quit chan bool
  d time.Duration
  running bool
}

// Checks the current public IP and returns it
func Check() string {
  resp, err := http.Get("http://ipv4.icanhazip.com")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  ip, _ := ioutil.ReadAll(resp.Body)
  return strings.Trim(string(ip), "\n")
}

// Returns a new IP Checker that will check at a specified duration. Must be started
func NewIPChecker(d time.Duration) *IPChecker {
  return &IPChecker{
    make(chan string),
    make(chan bool),
    d,
    false,
  }
}

// Shotcut to make an IP checker and get the channel. Useful if you won't need to stop the timer.
func Poll(every time.Duration) chan string {
  checker := NewIPChecker(every)
  checker.Start()
  return checker.C
}


// Starts the IP Checker checking for an ip and sends results on the channel
func (i *IPChecker) Start() {
  if(i.running) {
    return
  }
  i.running = true
  var oldIp string
  checkIp := func() {
    if ip := Check(); ip != oldIp {
      i.C <- ip
    }
  }
  go func() {
    go checkIp()
    for {
      select {
      case <- time.Tick(i.d):
        go checkIp()
      case <-i.quit:
        break
      }
    }
  }()
}

// Stops the IP Checker from checking for updates
func (i *IPChecker) Stop() {
  if(i.running) {
    close(i.quit)
    return
  }
  i.running = false
}

