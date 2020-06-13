package honlog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	tSize := 20
	logger := NewLoggerWithSize(tSize)
	t.Log("Write 10000 logs")
	for i:=0; i<10000; i++{
		logger.Add(fmt.Sprintf("Log: %d", i))
	}
	outChan := make(chan string)
	go func(){
		for {
			s, ok:= <- outChan
			if ok {
				fmt.Printf("Get: %s\n", s)
			} else {
				//fmt.Printf("Get: %s\n", s)
				t.Log("Channel close")
				return
			}
		}
	}()

	logger.OutputChan(outChan)

	testcb := func(s string) {
		fmt.Printf("Func Get: %s\n", s)
	}

	testdone := func() {
		fmt.Println("Func End")
	}

	logger.OutputFunc(testcb, testdone)
}
