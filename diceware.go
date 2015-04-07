package main

//go:generate rice embed-go

import (
	"bufio"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/GeertJohan/go.rice"
)

func main() {
	delim := flag.String("d", " ", "delimitor - the thing between the words")
	count := flag.Int("c", 8, "count - how many words")
	flag.Parse()

	fmt.Println(diceware(*delim, *count))
}

func diceware(delim string, count int) string {
	words := []string{}

	for i := 0; i < count; i++ {
		words = append(words, word())
	}

	return strings.Join(words, delim)
}

func word() string {
	search := ""

	// Roll 5 dice
	for i := 0; i < 5; i++ {
		j, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			panic(err)
		}

		search += strconv.Itoa(int(j.Uint64()) + 1)
	}

	return findWord(search)
}

func findWord(search string) string {
	box := rice.MustFindBox("box")

	f, err := box.Open("diceware.wordlist.asc")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, search+"\t") {
			return t[6:]
		}
	}
	panic(errors.New("not found"))
}
