package main

import (
  "bytes"
  "fmt"
  "log"
  "math"
  "os/exec"
  "strconv"
  "strings"
)

func main() {
  cmd := exec.Command("ioreg", "-rc", "AppleSmartBattery")
  var out bytes.Buffer
  cmd.Stdout = &out

  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }

  lines := strings.Split(out.String(), "\n")

  var max string
  var cur string

  for _, element := range lines {
    if strings.Contains(element, "CurrentCapacity") {
      cur = strings.TrimSpace(element)
    }

    if strings.Contains(element, "MaxCapacity") {
      max = strings.TrimSpace(element)
    }
  }

  maxNum, err := strconv.Atoi(strings.Split(max, " ")[2])
  if err != nil {
    log.Fatal(err)
  }

  curNum, err := strconv.Atoi(strings.Split(cur, " ")[2])
  if err != nil {
    log.Fatal(err)
  }

  charge := float64(curNum) / float64(maxNum)
  chargeThreshold := int(math.Ceil(charge * 10))

  totalSlots := 10
  filledCount := int(math.Ceil(float64(chargeThreshold) * 1.0))
  emptyCount := totalSlots - filledCount

  var buf bytes.Buffer
  var color string
  resetColor := "%{[00m%}"

  if filledCount > 6 {
    // Green
    color = "%{[32m%}"
  } else if filledCount > 4 {
    // Yellow
    color = "%{[33m%}"
  } else {
    color = "%{[31m%}"
  }

  buf.WriteString(color)
  for i := 0; i < filledCount; i++ {
    buf.WriteString("◼")
  }

  for i := 0; i < emptyCount; i++ {
    buf.WriteString("◻")
  }

  buf.WriteString(resetColor)
  fmt.Println(buf.String())
}
