//匿名函数
package main
import (
	"fmt"
)
func main(){
	//直接调用
	res1 := func (a int,b int) int {
		return a + b
	}(3,5)
	fmt.Println("res1=",res1)
	//用变量进行调用
	a := func (a int ,b int) int{
		return a - b
	}
	res2 := a(10,3)
	fmt.Println("res2=",res2)
}