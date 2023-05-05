/* !!
 * File: main.go
 * File Created: Thursday, 30th September 2021 10:21:52 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 30th September 2021 10:21:52 am
 
 */

package main

import (
	"fmt"

	"github.com/TrHung-297/fountain/baselib/validator"
)

type A struct {
	Age int
}

func main() {
	arrS := make([]string, 0)
	for i := 0; i < 10; i++ {
		arrS = append(arrS, fmt.Sprintf("%v", i))
	}
	fmt.Println(validator.IsUniqueIterator(arrS))

	arrN := make([]int, 0)
	for i := 0; i < 10; i++ {
		arrN = append(arrN, i)
	}
	arrN = append(arrN, 5)
	fmt.Println(validator.IsUniqueIterator(arrN))

	arrA := make([]A, 0)
	for i := 0; i < 10; i++ {
		arrA = append(arrA, A{i})
	}
	arrA = append(arrA, A{0})
	fmt.Println("arrA: ", validator.IsUniqueIterator(arrA))

	arrAP := make([]*A, 0)
	for i := 0; i < 10; i++ {
		arrAP = append(arrAP, &A{i})
	}
	arrAP = append(arrAP, &A{0})
	fmt.Println("arrAP: ", validator.IsUniqueIterator(arrAP))

	mInt := make(map[int]int)
	for i := 0; i < 10; i++ {
		mInt[i*i] = i
	}
	mInt[10] = 9
	fmt.Println(validator.IsUniqueIterator(mInt))

	mA := make(map[string]*A)
	for i := 0; i < 10; i++ {
		mA[fmt.Sprintf("%v", i*i)] = &A{i}
	}
	mA["10"] = &A{10}
	fmt.Println(validator.IsUniqueIterator(mA))
}
