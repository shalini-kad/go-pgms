
package main

import "fmt"
import "math"


type shape interface {
    area() float64
    
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}


func (r rect) area() float64 {
    return r.width * r.height
}


func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}



func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    fmt.Printf("%.2f\n",r.area())
    fmt.Printf("%.2f\n",c.area())
}
