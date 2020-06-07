package main

import "os"
import "fmt"
import "strconv"
import "simplemath"

var Usage = func() {
	fmt.Print("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe command are: \n\t add\t Addition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	fmt.Println(args)
	switch args[1] {
		case "add":
			if len(args) != 4 {
				fmt.Println("USAGE: calc add  <integer1> <integer2>")
				return
			}
			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: calc add  <integer1> <integer2>")
				return
			}
			ret := simplemath.Add(v1, v2)
			fmt.Println("Result : ", ret)
		case "sqrt":
			if len(args) != 3 {
				fmt.Println("USAGE: sqrt <integer>")
				return
			}
			v3, err3 := strconv.Atoi(args[2])
			if err3 != nil {
				fmt.Println("USAGE: sqrt <integer>")
				return
			}
			ret := simplemath.Sqrt(v3)
			fmt.Println("Result : ", ret)
		default:
			fmt.Println("This Is a default")
			Usage()
	}
}
