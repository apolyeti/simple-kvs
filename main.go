package main

import (
	"fmt"

	"github.com/apolyeti/simple-kvs/kvs"
)

func main() {
	store := kvs.New()
	fmt.Println("Setting key1 to 1")
	store.Set("key1", 1)
	fmt.Println("Setting key2 to 2")
	store.Set("key2", 2)
	fmt.Println("Getting key1")
	value, err := store.Get("key1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
	fmt.Println("Getting key2")
	value, err = store.Get("key2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
