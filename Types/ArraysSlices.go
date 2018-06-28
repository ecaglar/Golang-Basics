package main

import "fmt"

//Arrays and Slices are both numbered sequence of elements of a single type
//However size of arrays is fixed and constant so it does not change (expand)
//On contrary, slices size is dynamic and can change in need (double)

func main(){

	//lets make array
	var arr1 [5]int
	arr2 := [5]int{1,2,3,4,5}

	fmt.Println("Len of arr1 is : ",len(arr1), " content: ", arr1)
	fmt.Println("Len of arr2 is : ",len(arr2), " content: ", arr2)

	/*
	Len of arr1 is :  5  content:  [0 0 0 0 0]
	Len of arr2 is :  5  content:  [1 2 3 4 5]
	*/

	//arr2 = append(arr2,10 )
	//You cannot use append for arrays


	//lets make slices

	var names1 = make([]string, 3)  //OR
	names2 := make([]string, 3) //OR
	names3 := []string {"Tom", "Bill"}

	//Lets check types of slices
	fmt.Printf("Type of names1 : %T\n", names1)
	fmt.Printf("Type of names2 : %T\n", names2)
	fmt.Printf("Type of names3 : %T\n", names3)

	/*OUTPUT: So they are all string slices
	Type of names1 : []string
	Type of names2 : []string
	Type of names3 : []string*/

	fmt.Println("Content of names1 : ", names1)
	fmt.Println("Content  of names2 : ", names2)
	fmt.Println("Content  of names3 : ", names3)

	/*
	Content of names1 :  [  ]
	Content  of names2 :  [  ]
	Content  of names3 :  [Tom Bill]
	*/

	names1[0] = "Jenny"
	names1[1] = "Benny"
	names1[2] = "Kenny"

	fmt.Println("After modifying Content of names1 : ", names1)

	//After modifying Content of names1 :  [Jenny Benny Kenny]

	//names1[3] = "No-Jenny"
	//fmt.Println("After modifying 4th Content of names1 : ", names1)
	//ArraysSlices.go:44 +0x7aa 	//exit status 2
	//You cannot set 4th name as capasity is 3 so you should use append.

	names1 = append(names1,"Yes-Jenny" )
	fmt.Println("After appending Content of names1 : ", names1)
	//After appending Content of names1 :  [Jenny Benny Kenny Yes-Jenny]

	names1 = append(names1, "AnotherName1", "AnotherName2")
	fmt.Println("After appending multiple, Content of names1 : ", names1)
	//After appending multiple, Content of names1 :  [Jenny Benny Kenny Yes-Jenny AnotherName1 An
	//otherName2]
}
