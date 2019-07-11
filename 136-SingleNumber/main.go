package main

import (
    "fmt"
    "reflect"
    "strconv"
)
func singleNumber(nums []int) int {
    m := make(map[int]int)
    
    for _,num := range nums {
        _, ok := m[num]
        if ok {
            delete(m,num)
        } else{
            m[num] = 1
        }
    }
    
    keys := reflect.ValueOf(m).MapKeys()
    
    resultString := fmt.Sprintf("%v", keys[0])
	
	result,_ := strconv.Atoi(resultString)

	return result
}

func main() {
	fmt.Println(singleNumber([]int{3,5,5}))
}
