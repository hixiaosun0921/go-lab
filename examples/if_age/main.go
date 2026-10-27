package main
import (
	"fmt"
)
func main(){
	var age int 
	fmt.Println("输入年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("年龄大于18")
	} else {
		fmt.Println("年龄<18")
	}
}