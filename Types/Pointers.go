package main

import "fmt"

//Pointers are memory adresses to actual variables
//you can access to a variable by a memory address reference

func main(){

	i := 5
	/**
	there is a memory location where value 5 is being stored. (Lets say 0X123F) We also assign a
	symbolic name i to this memory address. So that, whenever we use name i, then
	we reach same memory address. There is another way to access that memory location.
	by pointers.
	 */

	 fmt.Println("value of i  : " , i)

	var mem_address_i1 *int = &i
	mem_address_i2 := &i
	 /**
	 Now we have memory address of i (which is 0X123F) & operator is being used to get memory address
	 Then we can access that memory location (value of i) by both mem_address_i1 and mem_address_i2
	  */

	  *mem_address_i1 = 10
	fmt.Println("value of i  after mem_address_i1 : " , i)

	*mem_address_i2 = 20
	fmt.Println("value of i  after mem_address_i2 : " , i)

	/*
	value of i  :  5
	value of i  after mem_address_i1 :  10
	value of i  after mem_address_i2 :  20
	*/


}
