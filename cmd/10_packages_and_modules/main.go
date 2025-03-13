package main

import (
	"fmt"
	"project_pkg/greet"
	m "project_pkg/math" // Import with renaming the package

	"github.com/enescakir/emoji"
)

func main() {
	/*
		Module
			- a collection of Go packages stored in a file tree with a go.mod file at its root

		syntax:
		go.mod: dependency manager
		module <name> => name of the module
		go <version>
		require <module-path> <version> => indicates dependencies
		replace <module-path> <version> => replace the local or fork version
		exclude <module-path> <version> => exlude the package; rarely used

		go.sum
			- include different checksums of the packages
			- help to build packages
			- not recommended to handle manually, go does automatically

		Versioning Module
		1: Module Path
		| go.mod
		| package.go
		| v2
			| go.mod
			| package.mod

		2: Branch
		.main = v0/v1
		| go.mod
		| package.go

		.v2 = v2
		| go.mod
		| package.go

		<module-path>/v2
	*/

	/*
		Packages
		 	- Every Go program is made up of packages.
			- Programs start running in package main.
	*/

	// Local package
	greet.Hello()

	xs := []float64{1, 1, 2, 3, 5}
	avg := m.Average(xs)
	fmt.Println(avg)

	// Remote package
	// syntax: go get <package-link>
	fmt.Println(emoji.WavingHand.Tone())

	// call the function with same package name (main)
	// While running use both files
	SamePackageName()
}
