package main

import (
	"log"
	"os"

	"github.com/lukechampine/fastxor"
)

func main() {
	in1, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	in2, err := os.Open(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}
	fi, err := in1.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sz := fi.Size() / 16
	// fmt.Println(fi.Size())
	// fi, err = in2.Stat()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	o, err := os.Create(os.Args[3])
	if err != nil {
		log.Println(err)
		return
	}
	defer in1.Close()
	defer in2.Close()
	defer o.Close()

	buf1 := make([]byte, 16)
	buf2 := make([]byte, 16)
	buf3 := make([]byte, 16)
	var i int64
	for i = 0; i < sz; i++ {
		in1.Read(buf1)
		in2.Read(buf2)
		fastxor.Block(buf3, buf1, buf2)

		_, err = o.Write(buf3)
		if err != nil {
			log.Println(err)
			return
		}
	}
	var v int
	if v = int(fi.Size() - sz*16); v != 0 {
		buf1 = buf1[0:v]
		buf2 = buf2[0:v]
		buf3 = buf3[0:v]

		in1.Read(buf1)
		in2.Read(buf2)

		fastxor.Bytes(buf3, buf1, buf2)
		_, err = o.Write(buf3)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
