package main
import (
	"fmt"
)
func main(){
	var n int = 4
	for i := 1 ; i  <= n ; i ++ {
		for k := n-i ; k > 0 ; k --{
			fmt.Print(" ")
		}
		for j := 1 ; j <= i * 2 - 1 ; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}