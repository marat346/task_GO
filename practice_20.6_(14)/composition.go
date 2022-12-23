package main

import (
	"fmt"
)
const rows = 3
const fiveCols = 5
const fourCols = 4

func composition (matrix1 [rows][fiveCols] int , matrix2 [fiveCols][fourCols] int) 
     [rows] [fourCols]int {
  var rsl [rows] [fourCols] int
  
  for i:= 0;i < rows;i++{
    for j:= 0; j < fourCols;j++ {
      for f:= 0; f < fiveCols;f++ {
           rsl [i][j] = rsl [i][j] + matrix1[i][f] * matrix2[f][j]
        }
      }
    }
  return rsl
  }
  
func main() {
  matrix1 := [rows][fiveCols] int {
    {1,2,3,4,5},
    {1,2,3,4,5},
    {1,2,3,4,5},
  }
  matrix2 := [fiveCols][fourCols] int {
    {1,2,3,4},
    {1,2,3,4},
    {1,2,3,4},
    {1,2,3,4},
    {1,2,3,4},
  }
  fmt.Println(composition(matrix1,matrix2))
 }