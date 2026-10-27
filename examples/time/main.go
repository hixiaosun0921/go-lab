package main
import(
	"fmt"
	"time"
)
func main(){
	//现在时间
	now := time.Now()
	fmt.Printf("现在时间：%v,数据类型%T\n",now,now)

	//通过now获得年月日……
	fmt.Printf("年=%v\n",now.Year())

	//格式化日期时间
}