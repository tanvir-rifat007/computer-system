#include<stdio.h>
#include<stdlib.h>

int bitCount(unsigned int x){

  int count = 0;

  while(x!=0){
    x &= (x - 1); // Clear the least significant bit set
    count++;
  }

  return count;


}

int main(){

  unsigned int x = 25;

  int count = bitCount(x);

  printf("The number of bits set in %u is: %d\n", x, count);

  return 0;








}