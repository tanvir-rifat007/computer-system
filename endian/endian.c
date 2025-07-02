// My device is arm64, so it is little endian
// and the output will be Little Endian also


#include<stdio.h>
#include<stdlib.h>


int main(){

    unsigned int x = 0x12345678;

    // so char in c used to be byte, so we can cast the integer to a char pointer

    char *ch = (char *)&x;
   
    // Right to left (little endian)   <-
    if(ch[0] == 0x78){

     printf("Little Endian\n");

    }
    // left to right -> Big indian
    else if(ch[0] == 0x12){
    
     printf("Big Endian\n");

    }else{

        printf("Unknown Endian\n");
    }


 return 0;
}
