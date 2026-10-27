package main
import (
	"fmt"
)
func f (x int) int {
	if x == 1 || x == 2 {
		return 1
	} else {
		return f(x-1)+f(x-2)
	}
	
}
func main(){
	a := f(4)
	fmt.Println(a)
}