package main
import (
	"fmt"
)
func main(){
	var num int = 3
	fmt.Println("地址",&num)
	var ptr *int = &num
	*ptr = 5
	fmt.Println("num的值",num)
	num_n := 3
	fmt.Println("num_n的值",num_n)
}