package main

var justString string

func someFunc() {
  v := createHugeString(1 &lt;&lt; 10)
  justString = string([]byte(v[:100]))
}

func main() {
  someFunc()
}
