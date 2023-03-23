package main

import (
  "fmt"
"strconv"
  "strings")

type Storage interface{
  Add (int) bool
  Size (int)
}