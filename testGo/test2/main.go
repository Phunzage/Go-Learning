package main

/*
#include <stdio.h>
void printSum(int a, int b) {
  printf("c:%d+%d=%d",a,b,a+b);
}
*/
import "C"

func main() {
	C.printSum(C.int(1), C.int(2))
}
