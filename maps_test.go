package main

import (
	"testing"
	"fmt"
)

//Format of a testing function
// func TestXXXXXXXXXXXXXXXXXX(t *testing.T){

// }

func TestIfKeyExistsOrNot(t *testing.T){
	//1. We have to create a dummy Map. and a key..
	var testMap map[string]Employee
	testMap = make(map[string]Employee)

	testMap["Asu"] = Employee{
		27,
		560078,
		"Marathahali",
	}

	testMap["Biswa"] = Employee{
		28,
		560068,
		"JPNagar",
	}

	//Biswa and Asu
	isKeyExists := CheckIfKeyExists("Biswa1", testMap)
	if !isKeyExists{
		t.Errorf("Test Case 0: Expected output is %v but got %v",true ,isKeyExists)
	}else{
		fmt.Printf("Test Case 0: Success, Expected output and output recieved is same.")
	}

	delete(testMap, "Biswa")

	isKeyExists = CheckIfKeyExists("Biswa", testMap)
	if !isKeyExists{
		t.Errorf("Test Case 1: Expected output is %v but got %v",true ,isKeyExists)
	}else{
		fmt.Printf("Test Case 1: Success, Expected output and output recieved is same.")
	}


}

