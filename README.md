CTF Crypto Challenge
====================

Given a salt, salted hash, and word list, determine which word from the word list corresponds to the salted hash (i.e., crack the password).

Download
--------

Clone the repo if you plan to do development, or simply download it using `go get`:

    go get github.com/danzek/ctf-crypto-challenge

Usage
-----

    go run main.go /path/to/word/list.txt salt:hash [salt:hash ...]

To generate 10 sample salted hashes from the word list (assumes word list is named `words.txt` and is in the `wordlist` folder):

    go run main.go crackme test

Of course you can compile it for your platform with `go build` and run whatever you name the executable as well.

Why?
---

I once helped a colleague with a specific CTF problem involving receiving 10 salted hashes (with salts provided) and having to decrypt them using a supplied word list in a short time frame. At the time, we decided to use Python but it was too slow to solve the challenge in the time provided. Now over a year later I'm learning Go and figured this might be a fun project to return to for educational purposes. If it helps you, even better!

License
-------

MIT License

Copyright &copy; 2017 Dan O'Day (d@4n68r.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
