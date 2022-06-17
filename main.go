package main

import (
	"fmt"
	"strings"
)

func main() {

	delimiter := strings.Repeat("-", 70)

	//func copy(dst, src []Type) int
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	//print initial data

	fmt.Println("Print initial data")
	fmt.Println("Info about SRC")
	printSliceParameters(src)
	fmt.Println("Info about DST")
	printSliceParameters(dst)
	fmt.Println(delimiter)

	//Example 1: simple copy execution, both slices have equal length

	fmt.Println("Example 1: simple copy execution, both slices have equal length")
	runCopyFuncAndPrintResult(src, dst)
	fmt.Println(delimiter)

	//Example 2: Copy from a long slice to a short slice

	fmt.Println("Example 2: Copy from a long slice to a short slice")
	src = append(src, 6)
	dst = make([]int, 5)
	runCopyFuncAndPrintResult(src, dst)
	fmt.Println(delimiter)

	//Example 3: Copy from a short slice to a long slice
	fmt.Println("Example 3: Copy from a short slice to a long slice")
	dst = make([]int, 8)
	runCopyFuncAndPrintResult(src, dst)
	fmt.Println(delimiter)

	//Example 4: Copy from an array to a slice
	fmt.Println("Example 4: Copy from an array to a slice")
	srcArray := [5]int{1, 2, 3, 4, 5}
	dst = make([]int, 5)
	runCopyFuncAndPrintResult(srcArray[:], dst)
	fmt.Println(delimiter)

	//Example 5: Copy from a slice to itself
	fmt.Println("Example 5: Copy from a slice to itself")
	runCopyFuncAndPrintResult(src, src[3:])
	fmt.Println(delimiter)

	//Example 6:  Copy from a string to a slice
	fmt.Println("Example 6:  Copy from a string to a slice")
	srcString := "Slice Copy Function"
	dstByte := make([]byte, len(srcString))
	copy(dstByte, srcString)
	fmt.Printf("Slice: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(dstByte), cap(dstByte), dstByte, dstByte, dstByte == nil)
	fmt.Println("First 5 symbols in destination slice - ", string(dstByte[0:6]))
	fmt.Println(delimiter)
}

//------------------------------------------------------------------------------------------------------------------------------------------------------------
func printSliceParameters(slice []int) {
	fmt.Printf("Slice: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(slice), cap(slice), slice, slice, slice == nil)
}

//------------------------------------------------------------------------------------------------------------------------------------------------------------

func runCopyFuncAndPrintResult(src []int, dst []int) {
	fmt.Println("Info about SRC before copying")
	printSliceParameters(src)
	fmt.Println("Info about DST before copying")
	printSliceParameters(dst)
	numberOfElementsCopied := copy(dst, src)
	fmt.Println("Info about SRC after copying")
	printSliceParameters(src)
	fmt.Println("Info about DST after copying")
	printSliceParameters(dst)
	fmt.Printf("Number of elements have been copied: %v\n", numberOfElementsCopied)

}
