package main

import "fmt"

const size = 10

func InsertionSort(array [size]int) [size] int{
	
  length := len(array)

	for i := 1; i < length; i++ {
		j := i
		   for j > 0 {
			     if array[j-1] > array[j] {
				    array[j-1], array[j] = array[j], array[j-1]
			    }
			j = j - 1
		}
	}
  return array
}

func main(){
  b := [size] int {5,4,3,2,1,10,9,8,7,6}
  fmt.Println("Исходный массив :\n",b)
  fmt.Println("Cортировка вставками :\n",InsertionSort(b))
}