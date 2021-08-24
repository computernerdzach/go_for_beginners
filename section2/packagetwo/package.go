package packagetwo

import "fmt"

var PackageVar = "[package level variable in package pagetwo]"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
