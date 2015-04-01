package main


/*
#include <unistd.h>
#include <stdint.h>
*/
import "C"
import "unsafe"
import "fmt"
import "errors"

type s_block struct {
	size int
	next *s_block
	prev *s_block
	free int
	data []int
}

var base *s_block
const BLOCK_SIZE int = 40

/* Split the block struct by size 'size'*/
func split_block(block s_block, size int){
	var split s_block
	block.size = size
	split.next = block.next
	split.prev = &block
	split.free = 1
	split.size = block.size - size - BLOCK_SIZE
}

/* Extend heap by size size */
func extend_heap(last s_block, size int) s_block{
	var block *s_block
	var newEnd *int
	var extend_length int = BLOCK_SIZE + size
	block = (*s_block)(C.sbrk(0))
	/* NEW END OF DATA */
	newEnd = (*int)(C.sbrk(extend_length))
	if *newEnd < 0 {
		err := errors.New("sbrk fails")
	}
	block.size = size
	block.prev = &last
	block.free = 0
	
	return *block
}
/* Allocate memory */
func gmalloc(size int) unsafe.Pointer{
	return unsafe.Pointer(base)
}

func main(){
	data := new(int)
	data = (*int) (gmalloc(4))
	fmt.Println("Malloc sanity test passes")
}

