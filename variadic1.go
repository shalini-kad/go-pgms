package main

import (  
    "fmt"
)

func avg(nums ...float64) { 
	var sum float64
	var average float64
	sum=0
    for _, v := range nums {
        sum+=v
	}
	average=sum/float64(len(nums))
    fmt.Printf("Avg =%.2f \n",average)
}
func main() {  
    avg(89, 89, 90, 95)
    avg(45, 56, 67, 45, 90, 109)

}