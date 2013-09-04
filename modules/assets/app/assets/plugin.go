package assets

import (
  "github.com/robfig/revel"
)

func init() {
  revel.OnAppStart(func() {
    revel.INFO.Println("Setting up assets plugin...")
    revel.TemplateFuncs["stylesheet"] = func(stylesheetName string) string {
      return ""
    }
    revel.TemplateFuncs["javascript"] = func(scriptName string) string {
      return ""
    }
  })
}
