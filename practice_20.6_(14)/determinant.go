package main

import (
	"fmt"
)
  const rows = 3
  const cols = 3

func determinant (detMatrix [rows][cols] int) int {
  
  onePart := (detMatrix [0][0] * detMatrix [1][1] * detMatrix [2][2]) +
             (detMatrix [0][1] * detMatrix [1][2] * detMatrix [2][0]) +
             (detMatrix [0][2] * detMatrix [1][0] * detMatrix [2][1])
  
  
  secondPart := (detMatrix [0][2] * detMatrix [1][1] * detMatrix [2][0]) +
                (detMatrix [0][0] * detMatrix [1][2] * detMatrix [2][1]) +
                (detMatrix [0][1] * detMatrix [1][0] * detMatrix [2][2])
  rsl := onePart - secondPart
  
  return rsl
}

func main() {
  matrix := [rows][cols] int {
    {3,1,7},
    {6,5,0},
    {2,8,4},
  }
  fmt.Println(determinant(matrix))
 }