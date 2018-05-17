package main

import "fmt"

func main() {
	count:=0
	sum:=0

	   for i:=2; count<=50; i++{


        isPrime:=true

        for j:=2; j<i; j++{

            if (i % j == 0 ){

                isPrime = false
            }
        }

        if isPrime == true {

			sum=sum+i
			count++
        } 
	}
	fmt.Println("The sum of first 50 numbers is :",sum)


}
		
		