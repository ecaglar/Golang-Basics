package main

import "fmt"

/*Note1:
you cannot mix different types such as uint + uint32

Note2:
You cannot mix types if they are alias.

type T1 = T2

T1 is a alias (and same type) of T2 so you can mix them such as;

T1 + T2
byte + uint8
rune + int32
*/

/////////////NUMERIC TYPES/////////////////////////

var unsigned_int_8_var uint8	//8-bit unsigned integer (also 16,32,64 bit)


var signed_int_8_var int8	//8-bit signed integer (also 16,32,64 bit)
var signed_int_32_var int32


var float_32_var float32
var float_64_var float64

var byte_var byte //uint8 (they are alias so you can mix them)
var rune_var rune //int32 (they are alias so you can mix them) rune represents a unicode code point

var var1 uint //either 32 or 64 bit
var var2 int  //either 32 or 64 bit

type myType = uint8		//myType is a alias type of uint8. So they are same type
var custom_type myType


//////////////////////////////////////////////////

func  main() {
	fmt.Println("Type of uint8 is: %T", unsigned_int_8_var);
	fmt.Println("Type of float32 is: %T", float_32_var);

	fmt.Println("Sum of byte + uint8", byte_var  + unsigned_int_8_var)
	fmt.Println("Sum of byte + uint8", byte_var  + unsigned_int_8_var)
	fmt.Println("Sum of byte + uint8", custom_type  + unsigned_int_8_var) //legal!! as they are alias
	// fmt.Println("Sum of rune + int32", unsigned_int_8_var + signed_int_32_var) you cant do that. different types

}
