package main
import (
	"fmt"
)
func addupper() func(int) int {
	var n int = 10
	return func (x int) int{
		n = n + x
		return n
	}
}
func main(){
	f := addupper()
	fmt.Println(f(1))//11
	fmt.Println(f(1))//12
	fmt.Println(f(2))//14
}