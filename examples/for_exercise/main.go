package main
import (
	"fmt"
)
//有100000元
//>50000,交5%
//<=50000，交1000
func main(){
	count := 0
	var m float64 = 100000
	for {
		if(m<1000){ 
			break
		}
		count ++
		if(m>50000){
			m -= m*0.05
		} else{
			m -= 1000
		}
	}
	fmt.Println("次数",count)
}