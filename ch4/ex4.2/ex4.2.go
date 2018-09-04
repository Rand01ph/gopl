package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashTyep = flag.String("type", "SHA256", "哈希算法")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch *hashTyep {
		case "SHA256":
			fmt.Printf("SHA256 %x\n", sha256.Sum256([]byte(arg)))
		case "SHA384":
			fmt.Printf("SHA384 %x\n", sha512.Sum384([]byte(arg)))
		case "SHA512":
			fmt.Printf("SHA512 %x\n", sha512.Sum512([]byte(arg)))
		case "MD5":
			fmt.Printf("MD5 %x\n", md5.Sum([]byte(arg)))
		default:
			fmt.Printf("SHA256 %x\n", sha256.Sum256([]byte(arg)))
		}
	}
}
