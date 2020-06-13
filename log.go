package honlog

import (
	"sync"
)
//var DefaultLogger *Logger
type Logger struct {
	data []string
	size int
	capacity int 	// capacity == 2*size
	//readHead int
	writeHead int

	lock sync.Mutex	//Write lock. We would not protect read.
}

const defaultSize = 1024

func NewLoggerWithSize(size int) *Logger {
	return &Logger {
		size: size,
		capacity: 2*size,
		data : make([]string,0 , 2*size),
		//readHead: 0,
		writeHead: 0,
	}
}

// Create a Logger with default size 1024
func NewLogger() *Logger {
	return &Logger{
		size : defaultSize,
		capacity: 2*defaultSize,
		data: make([]string, 0, 2*defaultSize),
		//readHead: 0,
		writeHead: 0,
	}
}

func (l *Logger) GetSize() int{
	return l.size
}

func (l *Logger) Add(str string) {
	l.data[l.writeHead] = str
	l.writeHead += 1
}

func (l *Logger) OutputChan(c chan<- string) {
	readHead := l.writeHead - l.size
	if readHead <0 {
		readHead += l.capacity
	}
	for  i:=0; i<l.size; i++ {
		str := l.data[(readHead+i)%l.capacity]
		c <- str
	}
	close(c)
}
