// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

func main() {
	var array = []byte{'x', 'b', 'c'}
	c1 := sha256.Sum256(array)
	c2 := sha256.Sum256(array)
	array[2] = '4'
	c3 := sha256.Sum256(array)
	fmt.Printf("%x\n%x\n%t\n%T\n%x\n", c1, c2, c1 == c2, c1, c3)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
