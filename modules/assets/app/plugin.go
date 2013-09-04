package app

import (
  "github.com/robfig/revel"
  "fmt"
)

func init() {
  revel.OnAppStart(func() {
    fmt.Println("Setting up assets plugin...")
  })
}
