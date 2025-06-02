package main

import (
	"encoding/binary"
	"fmt"
)

func main(){

  var num uint64 = 150


	encode(num);

	decode([]byte{0x96, 0x01})






}


func decode(nums []byte){

	// convert from hex to the decimal number:

	 num,n := binary.Uvarint(nums)

	 if n <= 0 {
		fmt.Println("Failed to decode varint")
		return
	}

	fmt.Printf("Decoded number: %d\n", num)



}

func encode(num uint64){



	// takes the most 10 bytes for varints

	buf:= make([]byte, binary.MaxVarintLen64)

	// Encode the number into varints format

	n:= binary.PutUvarint(buf,num)

	fmt.Printf("Variant encoding of decimal number %d : \n",n)

	for i:=0; i<n; i++{
		fmt.Printf("0x%02x ", buf[i])

	}

	fmt.Println();

}