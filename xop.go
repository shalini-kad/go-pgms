// Permutation Generator in Golang
package main
  
import (
	"fmt"  
	"strings"
	"strconv"
)   
 func join(ins []rune, c rune) (result []string) {
    for i := 0; i <= len(ins); i++ {
        result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
    }
    return
}
  
func permutations(testStr string) []string {
    var n func(testStr []rune, p []string) []string
    n = func(testStr []rune, p []string) []string{
        if len(testStr) == 0 {
            return p
        }else {
            result := []string{}
            for _, e := range p {
                result = append(result, join([]rune(e), testStr[0])...)
            }
            return n(testStr[1:], result)
        }
    }
  
    output := []rune(testStr)
    return n(output[1:], []string{string(output[0])})
}
func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}


func main() {  
	var str string
	var result []string
	var final []string
	str="*********"
	for i:=1;i<=9;i++{
		str=strings.Replace(str,"*",strconv.Itoa(i),1) 
		d := permutations(str)
		result = removeDuplicatesUnordered(d)
		final=append(final,result...)
		for j:=0;j<len(result);j++{
			result[j]=strings.Replace(result[j],"*","* ",-1) 
			result[j]=strings.Replace(result[j],"1","x1 ",-1) 
			result[j]=strings.Replace(result[j],"2","o1 ",-1)
			result[j]=strings.Replace(result[j],"3","x2 ",-1) 
			result[j]=strings.Replace(result[j],"4","o2 ",-1) 
			result[j]=strings.Replace(result[j],"5","x3 ",-1)
			result[j]=strings.Replace(result[j],"6","o3 ",-1)
			result[j]=strings.Replace(result[j],"7","x4 ",-1)
			result[j]=strings.Replace(result[j],"8","o4 ",-1) 
			result[j]=strings.Replace(result[j],"9","x5 ",-1)    
			fmt.Println(result[j])
		}
		fmt.Println("\n")
		}
		
	
	fmt.Println("Length=",len(final))
}
