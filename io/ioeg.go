package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	//doIO()
	//ConsoleEcho()
	//ConsoleEcho2()
	//fmt.Println("\u03C0")
	//doReadAt()
	//doWriteAt()
	//doReadFrom()
	//doByteIO()
	//doCopy()
	//copyarr()
	//doTee()
	doScan()
}

func doIO() {

	fmt.Println("请输入你的名字")
	from, _ := ReadFrom(os.Stdin, 10)
	fmt.Println(string(from))
	a, _ := os.Open("testfile.sql")
	readFrom, _ := ReadFrom(a, 6)
	fmt.Println(string(readFrom))
	bytes, _ := ReadFrom(strings.NewReader("abcde"), 4)
	fmt.Println(string(bytes))
	b3 := []byte{0b11101111, 0b10000111, 0b10101111}
	fmt.Println(string(b3))
}

func ReadFrom(reader io.Reader, n int) ([]byte, error) {
	bytes := make([]byte, n)

	read, err := reader.Read(bytes)

	if err != nil {
		println(err)
	}

	if read > 0 {
		bytes = bytes[:read]
	}
	return bytes, err

}

func ConsoleEcho() {
	stdin := os.Stdin

	bytes := make(chan []byte, 10)
	i := make([]byte, 128)
	go func() {
		for {
			fmt.Print(string(<-bytes))
		}

	}()
	for {

		read, _ := stdin.Read(i)
		//	fmt.Print(i,read)
		if read > 0 {
			bytes <- i[:read]
		}

	}

}

func ConsoleEcho2() {

	var (
		Name, Age string
	)

	//	bytes :=make([]byte,100)
	fmt.Scan(&Name, &Age)

	fmt.Println(Name, Age)

}
func doReadAt() {
	reader := strings.NewReader("12345你好!")

	bytes := make([]byte, reader.Len())
	at, _ := reader.ReadAt(bytes, 8)
	fmt.Printf("%s,%d", bytes, at)

}
func doWriteAt() {
	create, e := os.Create("w.tmp")
	if e != nil {
		panic(e)
	}
	defer create.Close()

	create.Write([]byte{48, 'a', 'A'})
	at, _ := create.WriteAt([]byte(`asdf1234`), 2)
	fmt.Println(at)

}

func doReadFrom() {
	open, err := os.Open("w.tmp")
	if err != nil {
		panic(err)
	}
	create, e := os.Create("copy.tmp")
	if e != nil {
		panic(e)
	}
	writer := bufio.NewWriter(create)

	writer.ReadFrom(open)
	writer.Flush()
	defer create.Close()
	defer open.Close()

}

func doByteIO() {
	b := byte(0)
	fmt.Scanf("%c\n", &b)
	b2 := new(bytes.Buffer)

	err := b2.WriteByte(b)
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := b2.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	}
}
func doCopy() {

	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF -- bye")

}
func copyarr() {
	var c = make([]int, 100)
	var d = make([]int, 10)

	rand.Seed(time.Now().Unix())
	for i := 0; i < len(c); i++ {
		c[i] = rand.Intn(1000)
	}
	for i := 0; i < len(c)-1; i++ {
		for j := 0; j < len(c)-i-1; j++ {
			if c[j] > c[j+1] {
				c[j], c[j+1] = c[j+1], c[j]
			}
		}
	}

	copy(d, c)

	fmt.Println(c)
	fmt.Println(d)
	for i := 0; i < len(d)-1; i++ {
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
		fmt.Println(d[len(d)-i-1])
	}
	fmt.Println(d)

}

func doTee() {
	//r := strings.NewReader("some io.Reader stream to be read\n")
	//r=
	var buf bytes.Buffer
	tee := io.TeeReader(os.Stdin, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(tee)
	printall(&buf)
}
func doScan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
