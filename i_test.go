package golangeg

import (
	"sync"
	"testing"
)

func TestGetFile(t *testing.T) {
	getFile()
	ReadFile("info.go")
	//WriteFile("testfile.sql","你好,世界!")
	//WriteFileAppend("./testfile.sql","Hello World")
	group := sync.WaitGroup{}
	//	var mutex sync.Mutex

	//	mutex.Lock()
	group.Add(2)
	go func() { WriteFileAppend("testfile.sql", "你好,世界!"); group.Done() }()

	go func() { WriteFileAppend("./testfile.sql", "Hello World"); group.Done() }()

	group.Wait()
}
