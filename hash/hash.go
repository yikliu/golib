package hash

import (
	//array "github.com/yikliu/golib/arraylist"
)

type hashset struct {
	size int,
	arr []interface{}
}

const DEFAULT_SIZE = 100

func New() {
	arr := [DEFAULT_SIZE]interface{}
	return &hashset{
		size : DEFAULT_SIZE,
		arr = &arr
	}
}

func hashFunc (key int) (slot int) {
	
}

func (hs *hashset) Insert(key int) error {
	slot := hashFunc(key)
	
}



