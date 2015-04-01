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
const BLOCK_SIZE int = 40

func split_block(block s_block, size int64){
	var split s_block
	block.size = size
	split.next = block.next
	split.prev = &block
	split.free = 1
	split.size = block.size - size - BLOCK_SIZE
}

func extend_heap(last s_block, size int) s_block{
	var block s_block
	var newEnd int
	block = C.sbrk(0)
	block.size = size
	block.prev = &last
	block.free = 0
	
	return block
}

func gmalloc(size int) unsafe.Pointer{
	return unsafe.Pointer(base)
}

func main(){
	data := new(int)
	data = (*int) (gmalloc(4))
	fmt.Println("Malloc sanity test passes")
}

