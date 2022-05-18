package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	pointer()

	segitiga(5)
	genPass("abcdefg", "low")
	genPass("abcdefg", "medium")
	genPass("abcdefg", "strong")
	hasil := sumFilm(8)
	fmt.Print(hasil)
}

// Mencetak segitiga terbalik
func segitiga(rows int) {
	if rows == 0 {
		fmt.Println("Error")
	} else {
		var i, j int
		i = rows
		for ; i >= 1; i-- {
			space := 1
			for ; space <= rows-i; space++ {
				fmt.Print(" ")
			}
			j = i
			for ; j <= 2*i-1; j++ {
				fmt.Printf("*")
			}
			j = 0
			for ; j < i-1; j++ {
				fmt.Printf("*")
			}
			fmt.Println("")
		}
	}
}

// Generate Password sesuai input dan level
func genPass(password string, level string) {
	rand.Seed(time.Now().Unix())
	if len(password) < 6 {
		fmt.Println("Password too short")
	} else if level == "low" {
		fmt.Println(strings.Title(password))
	} else if level == "medium" {
		rangeLower := 0
		rangeUpper := 99
		num := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		pass := strings.Title(password)
		concated := pass + strconv.Itoa(num)
		fmt.Println(concated)
	} else if level == "strong" {
		rangeLower := 0
		rangeUpper := 99
		num := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		pass := strings.Title(password)

		var symbol = []rune("~!@#$%^&*()_-+=")
		s := symbol[rand.Intn(len(symbol))]
		genSymbol := string(s)

		concated := pass + genSymbol + strconv.Itoa(num)
		fmt.Println(concated)
	} else {
		fmt.Println("error")
	}
}

// Membandingkan durasi film dengan lama terbang
func sumFilm(lamaTerbang int) string {
	arrFilm := []int{1, 7, 3, 4, 8, 9}
	result := ""
	x := 0
	for i := 0; i < len(arrFilm); i++ {
		if x < len(arrFilm)-1 {
			x += 1
		} else {
			x += 0
		}

		if calculate := arrFilm[i] + arrFilm[x]; calculate == lamaTerbang {
			result = strconv.Itoa(arrFilm[i]) + " dan " + strconv.Itoa(arrFilm[x])
		}
	}
	return result
}

// Contoh Golang pointer
func pointer() {
	var x int = 123
	var p *int
	p = &x

	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
}
