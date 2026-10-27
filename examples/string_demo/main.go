package main
import (
	"fmt"
	"project01/model"
)
func main(){
	var str string = `abcd`
	str = `defg`//这样可以
	//str[0]='z'   //这样不行
	fmt.Println(str)
	fmt.Println("name",model.Name)
}