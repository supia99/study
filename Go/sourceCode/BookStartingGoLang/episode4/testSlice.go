package main

import (
	"fmt"
	"sort"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func testSlice() {
	var sli []int = make([]int, 1)
	var seq = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("type %T %T\n", sli, seq)
	fmt.Println("sli len:", len(sli))
	fmt.Println("seq len:", len(seq))
	fmt.Println("cap", cap(sli), " ", cap(seq))
	fmt.Println("len", len(sli), " ", len(seq))
	reverse(seq[:])
	fmt.Println("reverse:", seq[:])
	// var copySeq [10]int
	// copy(copySeq[0:4], seq[:])
	bufSeq := seq[1:3]
	fmt.Println(len(bufSeq), " ", cap(bufSeq))
}

func testReverse() {
	// squ := []string{"bc", "vv", "va"}
	squInt := []int{3, 6, 2}

	fmt.Println("^^^testReverse^^^")
	// sort.Sort(sort.IntSlice(squ))
	// fmt.Println(squ)
	sort.Sort(sort.Reverse(sort.IntSlice(squInt)))
	fmt.Println(squInt)

}

func testSliceAppend() {
	fmt.Println("\n----testSliceAppend---")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%10v %3d %3d %8v %8v\n", "primes", len(primes), cap(primes), &primes[0], primes)

	var s []int = primes[3:6]
	fmt.Printf("%10v %3d %3d %8v %8v\n", "s", len(s), cap(s), &s[0], s)
	s2 := append(s, 15) //別メモリを確保する
	fmt.Printf("%10v %3d %3d %8v %8v\n", "s2 append", len(s2), cap(s2), &s2[0], s2)
	fmt.Println("primes:", primes)

	s = primes[0:3]
	fmt.Println(" s:", &s[0])
	s3 := append(s, 16) //primes配列内を書き込む
	fmt.Printf("%10v %3d %3d %8v %8v\n", "s3 append", len(s3), cap(s3), &s3[0], s3)
	fmt.Println("primes:", primes)
	//fmt.Println(append(primes,15))

	p := &primes[0]
	fmt.Printf("%10v %3v\n", "point", *p)
	p++
}

func main() {
	testSlice()
	// testReverse()
	testSliceAppend()
}
