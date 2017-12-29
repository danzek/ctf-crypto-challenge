// generate 10 salted hashes to test cracking on
// assumes hashed string uses colon delimiter and is in order salt:hash
// path to word list is hardcoded (word list not checked into repo)

/*
Copyright © 2017 Dan O'Day (d@4n68r.com)

This work is free. You can redistribute it and/or modify it under the terms of the Do What The Fuck You Want To Public
License, Version 2, as published by Sam Hocevar. See the COPYING file for more details.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package wordlist

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"math/rand"
	"time"
	"path/filepath"
	"os"
	"sort"
	"bufio"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#@!$%^&*()-_+=[]{};'?/>.<,|\\")

func Generate10() {
	// seed PRNG
	rand.Seed(time.Now().UTC().UnixNano())

	// get path to word list
	wordlist, err := filepath.Abs("wordlist/words.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding word list: %q\n", err)
		return
	}

	tenRandomWords := getTenRandomWords(wordlist)
	fmt.Println("#\tWord\tSalt:Hash")
	for i, word := range tenRandomWords {
		salt := getSalt(rand.Intn(10)+10)
		saltedWord := []byte(fmt.Sprintf("%s:%s", salt, word))
		var hash string
		switch {
		case i < 3:
			hash = fmt.Sprintf("%x", md5.Sum(saltedWord))
		case i < 6:
			hash = fmt.Sprintf("%x", sha1.Sum(saltedWord))
		case i < 8:
			hash = fmt.Sprintf("%x", sha256.Sum256(saltedWord))
		default:
			hash = fmt.Sprintf("%x", sha512.Sum512(saltedWord))
		}
		fmt.Printf("%d\t%s\t%s:%s\n", i, word, salt, hash)
	}
}

func getTenRandomNumbers() []int {
	// choose 10 unique random numbers
	randomNumbers := make([]int, 0, 10)
	for i := 0; i < 10; i = len(randomNumbers) {
		prn := rand.Intn(10000)
		if !containsInt(randomNumbers, prn) {
			randomNumbers = append(randomNumbers, prn)
		}
	}
	return randomNumbers
}

func containsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func getTenRandomWords(path string) []string {
	randomNumbers := getTenRandomNumbers()
	sort.Ints(randomNumbers)
	randomWords := make([]string, 0, 10)

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error opening word list\n")
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	i, j := 0, 0
	for scanner.Scan() {
		if j < len(randomNumbers) && i == randomNumbers[j] {
			randomWords = append(randomWords, strings.TrimSpace(scanner.Text()))
			j++
		}
		i++
	}

	return randomWords
}

func getSalt(length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(s)
}
