CTF Crypto Challenge
====================

Given a salt, salted hash, and word list, determine which word from the word list corresponds to the salted hash (i.e., crack the password).

Download
--------

You can clone the repo if you plan to do development, or simply download it using `go get`:

    go get github.com/danzek/ctf-crypto-challenge

Usage
-----

    go run main.go /path/to/word/list.txt salt:hash [salt:hash ...]

To generate 10 sample salted hashes from the word list (assumes word list is named `words.txt` and is in the `wordlist` folder):

    go run main.go crackme test

Why?
---

I once helped a colleague with a specific CTF problem involving receiving 10 salted hashes (with salts provided) and having to decrypt them using a supplied word list in a short time frame. At the time, we decided to use Python but it was too slow to solve the challenge in the time provided. Now over a year later I'm learning Go and figured this might be a fun project to return to for educational purposes. If it helps you, even better!

License
-------

Copyright &copy; 2017 Dan O'Day (d@4n68r.com)

This work is free. You can redistribute it and/or modify it under the terms of the Do What The Fuck You Want To Public License, Version 2, as published by Sam Hocevar. See the COPYING file for more details.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

