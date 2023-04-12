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
	n := readInt1()
	socks := readIntList()
	eachSocks := make(map[int]int)

	for _, sock := range socks {
		eachSocks[sock]++
	}

	ans := n - n
	for _, sock := range eachSocks {
		ans += sock / 2
	}
	fmt.Println(ans)
}

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
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[3])
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
	if s == "" {
		s = "0"
	}
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
func nextInt() int {
	sc.Scan()
	ret, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return ret
}
