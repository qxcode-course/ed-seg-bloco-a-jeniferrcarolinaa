//usando if/else em Go:
//if <condição>{
//codigo
//}else if <condicao>{
//codigo
//}else{
//codigo
//}

package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scan(&name)
	fmt.Scan(&age)

	if age < 12 {
		fmt.Printf("%s eh crianca\n", name)
	} else if age < 18 {
		fmt.Printf("%s eh jovem\n", name)
	} else if age < 65 {
		fmt.Printf("%s eh adulto\n", name)
	} else if age < 1000 {
		fmt.Printf("%s eh idoso\n", name)
	} else {
		fmt.Printf("%s eh mumia\n", name)
	}

}
