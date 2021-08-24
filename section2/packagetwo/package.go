package packagetwo

import "fmt"

var PackageVar = "[package level variable in package pagetwo]"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
