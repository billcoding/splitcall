package splitcall

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[split call]", log.LstdFlags)

func Interface(list []interface{}, fn func(subList []interface{}), per uint) {
	if list == nil {
		logger.Println("list is nil")
		return
	}
	size := len(list)
	if size <= 0 {
		logger.Println("list is empty")
		return
	}
	logger.Printf("list size is %d\n", size)
	logger.Printf("per size is %d\n", per)
	peer, index, have := int(per), 0, 0
	for {
		if size-have > peer {
			fn(list[index*peer : (index+1)*peer])
		} else {
			fn(list[index*peer:])
		}
		have += peer
		if have >= size {
			break
		}
		index++
	}
}

func String(list []string, fn func(subList []string), per uint) {
	if list == nil {
		logger.Println("list is nil")
		return
	}
	size := len(list)
	if size <= 0 {
		logger.Println("list is empty")
		return
	}
	logger.Printf("list size is %d\n", size)
	logger.Printf("per size is %d\n", per)
	peer, index, have := int(per), 0, 0
	for {
		if size-have > peer {
			fn(list[index*peer : (index+1)*peer])
		} else {
			fn(list[index*peer:])
		}
		have += peer
		if have >= size {
			break
		}
		index++
	}
}

func Int(list []int, fn func(subList []int), per uint) {
	if list == nil {
		logger.Println("list is nil")
		return
	}
	size := len(list)
	if size <= 0 {
		logger.Println("list is empty")
		return
	}
	logger.Printf("list size is %d\n", size)
	logger.Printf("per size is %d\n", per)
	peer, index, have := int(per), 0, 0
	for {
		if size-have > peer {
			fn(list[index*peer : (index+1)*peer])
		} else {
			fn(list[index*peer:])
		}
		have += peer
		if have >= size {
			break
		}
		index++
	}
}
