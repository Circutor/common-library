// Copyright (c) 2021 Circutor S.A. All rights reserved.

package main

import "fmt"

var version string

const projectName = "common-library"

func main() {
	fmt.Printf("Started %s v%s\n", projectName, version)
}
