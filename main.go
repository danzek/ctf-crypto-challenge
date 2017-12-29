// detect type of encryption used given salted hash and decrypt using provided salt
// path to word list must be provided as command-line parameter

// todo -- allow command-line parameter to specify order of salt:hash (is salt first or second?)

package main

import (
	//"crypto/md5"
	//"crypto/sha1"
	//"crypto/sha256"
	//"crypto/sha512"
	"fmt"
	"strings"
	// "path/filepath"
	"os"
	"crypto_ctf_challenge/wordlist"
)

func main() {
	// establish path to word list
	if len(os.Args) > 2 {
		wordListPath := os.Args[1]

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

	// populate map of salted hashes (key is hash, value is salt)
	hashSaltMap := make(map[string]string)
	for _, sh := range os.Args[2:] {
		splitString := strings.Split(sh, ":")
		if len(splitString) == 2 {
			hashSaltMap[splitString[1]] = splitString[0]
		} else {
			fmt.Fprintf(os.Stderr, "Invalid salt:hash -- %s\n", sh)
		}
	}
}

func showUsage() {
	fmt.Fprint(os.Stderr, "Usage: ./crack /path/to/word/list.txt salt:hash [salt:hash ...]\n")
}
