package main

import "fmt"

//Note:
//String types are immutable so they are constant you cannot change its value unless you generate a new string

var str string  = "This is a string" //set to ""

var l = len(str)

//for below also check array and slide scripts
var sliced_1 = str[0:5]
var sliced_2 = str[:5]
var sliced_3 = str[5:]


//more advanced below

//Golang is unicode UTF-8 encoding scheme
//Unicode is 1 to 4 bytes. First 1 byte same  as ASCII

var unicode_ch1 rune = 'ぼ'
var unicode_ch2 rune = '⌘'


func main(){
	fmt.Println("String itself: ", str)
	fmt.Println("Length of String : ", l)
	fmt.Println("String sliced_1[0:5]: ", sliced_1)
	fmt.Println("String sliced_2[:5]: ", sliced_2)
	fmt.Println("String sliced_3[0:]: ", sliced_3)

	fmt.Printf("Unicode representation of ぼ : %U\n", unicode_ch1)
	fmt.Printf("Unicode representation of ⌘ : %U\n", unicode_ch2)

	// str[0] = 'S' you cant to that as it is immutable so it is constant
}

