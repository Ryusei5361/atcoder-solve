package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

type bombLocation struct {
	locationR int
	locationC int
	level     int
}

func main() {
	r, c := readInt2()
	mat := make([][]string, 0)
	bombLocations := make([]bombLocation, 0)
	for i := 0; i < r; i++ {
		line := readNotSpaceStringList()
		mat = append(mat, line)
		for j := 0; j < c; j++ {
			// 爆弾の判定
			if line[j] != "." && line[j] != "#" {
				bombLocations = append(bombLocations, bombLocation{locationR: i + 1, locationC: j + 1, level: s2i(line[j])})
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for _, x := range bombLocations {
				dist := manhattan(i+1, x.locationR, j+1, x.locationC)
				if dist <= x.level && mat[i][j] != "." {
					mat[i][j] = "."
				}
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%s", mat[i][j])
		}
		fmt.Println()
	}
}

func manhattan(r1, r2, c1, c2 int) int {
	return int(math.Abs(float64(r1-r2)) + math.Abs(float64(c1-c2)))
}
