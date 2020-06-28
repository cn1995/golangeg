package golangeg

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func getFile() (s string) {
	open, err := os.Open("README.md")
	if err != nil {
		fmt.Println(err)
	}
	defer open.Close()
	var bc []byte = make([]byte, 1024)
	for true {
		/// 读取文件末尾时 read=0 ,e=io.EOF
		read, e := open.Read(bc)
		if e != io.EOF {
			s = s + string(bc[:read])
			fmt.Println(read)
			//	log.Fatal(s)
		} else {
			break
		}

	}
	fmt.Println(s)
	return
}

func ReadFile(s string) {
	file, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println(err)
	}
	s2 := string(file)

	fmt.Println(s2)
}
func WriteFile(f string, s string) {
	bytes := []byte(s)

	ioutil.WriteFile(f, bytes, 0661)
}

func WriteFileAppend(f string, s string) {
	//	bytes := []byte(s)

	file, err := os.OpenFile(f, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(s)
	defer file.Close()

}
