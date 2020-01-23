package main

import "fmt"

//Employee structs the employee data
type Employee struct{
	Age int
	ZipCode int
	Address string
}

var m map[string]Employee // Inits a map 'm'

//CheckIfKeyExists checks if a key is present or not
func CheckIfKeyExists(key string, m map[string]Employee) bool{
	_, exists := m[key]
	return exists
}

func main(){
	m = make(map[string]Employee)

	m["PoojaShetty"] = Employee{
		25, 
		560068,
		"Banaglare",
	}
	m["BanuPrakash"] = Employee{
		28, 560089, "Banashankari",
	}	

	fmt.Println(m)

	fmt.Println(m["BanuPrakash"])

	delete(m,"BanuPrakash")
	

	fmt.Println(m)
	
	fmt.Println("Check if Banu is presentor not")
	fmt.Println(CheckIfKeyExists("PoojaShetty", m))
	
}