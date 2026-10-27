package main
import "fmt"
func main(){
	var i int = 10
	j := 20
	var k = 30
	var a int //默认值为0
	fmt.Println("i=",i,"\tj=",j,"\tk=",k,"\ta=",a)
	var b,c int = 1,2
	var str,d="abc",4
	e,f := "def",2.5
	fmt.Printf("b=%d,c=%d,str=%s,d=%d,e=%s,f=%f",b,c,str,d,e,f)
}