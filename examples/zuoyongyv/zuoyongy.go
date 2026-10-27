package main
import (
	"fmt"
)
var name = "tom"
//name ="123" 错误
func test1(){
	fmt.Println(name)
}
func test2(){
	var name = "jack"
	fmt.Println(name)
}
func main(){
	fmt.Println(name)
	test1()
	test2()
	test1()
}