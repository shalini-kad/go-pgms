// _Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["Asha"] = 78
    m["Ben"] =54

    fmt.Println("map:", m)

    fmt.Println("Marks of ben: ", m["Ben"])

    fmt.Println("len:", len(m))

    delete(m, "Asha")
    fmt.Println("Map after deleting:", m)

    _, prs := m["Asha"]
	fmt.Println("Is the value present in map:", prs)
	
    n := map[string]int{"India": 1, "Sri Lanka": 2}
    fmt.Println("New map:", n)
}
