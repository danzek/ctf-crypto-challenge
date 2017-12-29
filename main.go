// detect type of encryption used given salted hash and decrypt using provided salt
// path to word list must be provided as command-line parameter

/*
Copyright © 2017 Dan O'Day (d@4n68r.com)

This work is free. You can redistribute it and/or modify it under the terms of the Do What The Fuck You Want To Public
License, Version 2, as published by Sam Hocevar. See the COPYING file for more details.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/* todo:
    - allow command-line parameter to specify order of salt:hash (is salt first or second?)
    - accept salted hashes in file parameter where each salted hash is on a new line in the file
    - concurrently crack passwords
*/

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strings"
	"os"
	"crypto_ctf_challenge/wordlist"
	"bufio"
)

var hashSizes = map[int]string {
	md5.Size: "MD5",
	sha1.Size: "SHA1",
	sha256.Size: "SHA256",
	sha512.Size: "SHA512",
}

func main() {
	// establish path to word list
	var wordListPath string

	if len(os.Args) > 2 {
		wordListPath = os.Args[1]

		// for testing
		if wordListPath == "crackme" {
			wordlist.Generate10()
			return
		}

		fileInfo, err := os.Stat(wordListPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with path to word list: %q\n", err)
			return
		} else if fileInfo.IsDir() {
			fmt.Fprint(os.Stderr, "Word list path must point to a file, not a directory.\n")
			showUsage()
		}
	} else {
		fmt.Fprint(os.Stderr, "Missing required argument(s).\n")
		showUsage()
	}

	// crack passwords
	// todo -- use goroutines / concurrency
	for _, sh := range os.Args[2:] {
		// iterate over word list and crack password
		if word, ok := crack(sh, wordListPath); ok {
			fmt.Printf("CRACKED: %s\n", word)
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", word)
		}
	}
}

func showUsage() {
	fmt.Fprint(os.Stderr, "Usage: go run main.go /path/to/word/list.txt salt:hash [salt:hash ...]\n")
}

func crack(saltedHash string, wordListPath string) (string, bool) {
	splitString := strings.Split(saltedHash, ":")
	if len(splitString) == 2 {
		salt := splitString[0]
		hash := splitString[1]

		if hashType, ok := hashSizes[len(hash)/2]; ok {
			// iterate over word list
			f, err := os.Open(wordListPath)
			if err != nil {
				fmt.Fprint(os.Stderr, "Error opening word list\n")
				panic(err)
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				word := strings.TrimSpace(scanner.Text())
				saltedWord := []byte(fmt.Sprintf("%s:%s", salt, word))

				// calculate hashResult
				var hashResult string
				switch hashType {
				case "MD5":
					hashResult = fmt.Sprintf("%x", md5.Sum(saltedWord))
				case "SHA1":
					hashResult = fmt.Sprintf("%x", sha1.Sum(saltedWord))
				case "SHA256":
					hashResult = fmt.Sprintf("%x", sha256.Sum256(saltedWord))
				case "SHA512":
					hashResult = fmt.Sprintf("%x", sha512.Sum512(saltedWord))
				}

				// compare hashResult to hash
				if hashResult == hash {
					return word, true
				}
			}
		} else {
			return "UNRECOGNIZED HASH TYPE", false
		}
	} else {
		return "INVALID SALT:HASH", false
	}

	return "WORD NOT FOUND", false
}
