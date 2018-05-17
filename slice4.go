
package main

import "fmt"

func main() {

  
    s := make([]string, 3)
    fmt.Println("initial slice:", s)
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("slice after setting values:", s)
    fmt.Println("second index in slice:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("After appending:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("copy of slice:", c)
    l := s[2:5]
    fmt.Println("slice from elements 2:5:", l)


    l = s[:5]
    fmt.Println("slice till 5:", l)

  
    t := []string{"g", "h", "i"}
    fmt.Println("new slice:", t)

  
    
}
