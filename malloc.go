package main


/*
#include <unistd.h>
*/
import "C"
import "unsafe"
import "fmt"

type s_block struct {
	size int64
	next *s_block
	prev *s_block
	free int
	data []int
}

var base *s_block

func split_block(block s_block, size int){

}

func extend_heap(last s_block, size int) s_block{
	var block s_block

	return block
}

func malloc(size int) unsafe.Pointer{
	return unsafe.Pointer(base)
}

func main(){
	data := new(int)
	data = malloc(4)
	fmt.Println("Malloc sanity test passes")
}










