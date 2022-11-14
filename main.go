package main
 
import "fmt"
 
func main() {
  
  one := 0
  second := 0
  third := 0 

  for i := 0; i <= 1000; i++{
    fmt.Println(i)
    one+= i
    if one % 10 != 0 {
      one++
      continue
     }
    fmt.Println("Наполнена one на 10")
    
    if second % 100 != 0 {
       second++
         continue
     }
    fmt.Println("Наполнена one на 100")
    
    if third % 1000 != 0 {
      third++
      continue
    }
    fmt.Println("Наполнена one на 1000")
  }
}