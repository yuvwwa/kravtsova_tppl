package main

import (
	"fmt"
	"log"
	
	"tppl/pascal/pascal"
)

func main() {
	// Test 1
	fmt.Println("Test 1: Empty program")
	test1 := `BEGIN
END.`
	vars1, err := pascal.Interpret(test1)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Variables: %v\n\n", vars1)

	// Test 2
	fmt.Println("Test 2: Arithmetic")
	test2 := `BEGIN
	x:= 2 + 3 * (2 + 3);
	y:= 2 / 2 - 2 + 3 * ((1 + 1) + (1 + 1))
END.`
	vars2, err := pascal.Interpret(test2)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Variables: %v\n\n", vars2)

	// Test 3
	fmt.Println("Test 3: blocks")
	test3 := `BEGIN
	y := 2;
	BEGIN
		a := 3;
		a := a;
		b := 10 + a + 10 * y / 4;
		c := a - b
	END;
	x := 11
END.`
	vars3, err := pascal.Interpret(test3)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Variables: %v\n", vars3)
}