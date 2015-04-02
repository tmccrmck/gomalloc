/* A simple test for gmalloc */

#include <stdio.h>
#include "_cgo_export.h"

int main(int argc, char **argv){
  int *data;
  data = (int*)gmalloc(4);

  printf("Malloc works with no error\n");
  return 0;
}

