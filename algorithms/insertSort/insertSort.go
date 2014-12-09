package main

import "fmt"

func insertSort(inList []int) []int {
    outList := make([]int, 0)

    // for every item
    for j, k := range inList {
        if j == 0 {
	    outList = append(outList, k)
        } else {
	    // walk down outList
	    for i := 0; i < len(outList); i++ {
	   	if outList[i] > k {
	            // k belongs at position i
	            outList = append(outList, 0)
		    // copy(x, y) is bulltin function that copies the items from y
		    // into x
	            copy(outList[i + 1:], outList[i:])
		    outList[i] = k
		    break
		}
	    }
	}
    }
    return outList
}

func main() {
    inList := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
    fmt.Println(insertSort(inList))
}
