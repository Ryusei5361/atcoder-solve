package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	_, p, q, r, s := readInt5()
	a := readIntList()
	pp := a[:p-1]
	pq := a[p-1 : q]
	qr := a[q : r-1]
	rs := a[r-1 : s]
	ss := a[s:]

	var result []int
	result = append(result, pp...)
	result = append(result, rs...)
	result = append(result, qr...)
	result = append(result, pq...)
	result = append(result, ss...)
	for _, x := range result {
		fmt.Printf("%d ", x)
	}

}

const MaxInt = int(^uint(0) >> 1)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt1() int {
	sc.Scan()
	return s2i(sc.Text())
}

func readInt2() (int, int) {
	sc.Scan()
	lines := strings.Split(sc.Text(), " ")
	return s2i(lines[0]), s2i(lines[1])
}

func readInt3() (int, int, int) {
	sc.Scan()
	lines := strings.Split(sc.Text(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2])
}

func readInt5() (int, int, int, int, int) {
	sc.Scan()
	lines := strings.Split(sc.Text(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2]), s2i(lines[3]), s2i(lines[4])
}

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func readIntList() []int {
	sc.Scan()
	intList := make([]int, 0)
	text := strings.Split(sc.Text(), " ")
	for _, s := range text {
		intList = append(intList, s2i(s))
	}
	return intList
}

func readSpaceStringList() []string {
	sc.Scan()
	stringList := make([]string, 0)
	text := strings.Split(sc.Text(), " ")
	for _, s := range text {
		stringList = append(stringList, s)
	}
	return stringList
}

func readNotSpaceStringList() []string {
	sc.Scan()
	stringList := make([]string, 0)
	text := strings.Split(sc.Text(), "")
	for _, s := range text {
		stringList = append(stringList, s)
	}
	return stringList
}

func sliceStringUnique(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func sliceIntUnique(target []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// 昇順にソート
func sortIntListAsc(intList []int) []int {
	sort.Slice(intList, func(i, j int) bool {
		return intList[i] < intList[j]
	})
	return intList
}

// 降順にソート
func sortIntListDes(intList []int) []int {
	sort.Slice(intList, func(i, j int) bool {
		return intList[i] > intList[j]
	})
	return intList
}

// String -> Int
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

// Int -> String
func i2s(i int) string {
	return strconv.Itoa(i)
}

func sumOfInts(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxOfInts(a []int) int {
	res := -MaxInt
	for _, v := range a {
		res = max(res, v)
	}
	return res
}

func minOfInts(a []int) int {
	res := MaxInt
	for _, v := range a {
		res = min(res, v)
	}
	return res
}
