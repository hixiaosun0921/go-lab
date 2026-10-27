package main
import(
	"fmt"
	"errors"
)
func test(name string)(err error){
	if name == "123"{
		return nil
	}else{
		return  errors.New("文件错误")
	}
}
func test1(){
	err := test("123")
	if err != nil{
		panic(err)
	}
	fmt.Println("后续执行")
}
func main(){
	test1()
	fmt.Println("后续执行1")
}