package main

import (
	"fmt"
	"flag"
	"bytes"
	"time"
)

var characters = []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

func main() {
	var needle = flag.String("s", "foo", "the string to be cracked")
	var started_at = time.Now()
	flag.Parse()

	fmt.Println("Finding needle:", *needle)

	var haystack = make([]int, len(*needle))
	var hay string

	for (hay != *needle) {
		var added bool = false
		

		// check whethere there are any ints that are likely to go above the number of designated characters
		// and if there are, increment the next position by 1 reset the overflowing to 0
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == len(characters)-1 {
				haystack[i] = 0

				haystack[i+1] = haystack[i+1] + 1
				added = true

				break;
			}
		}

		// if nothing was overflowing, simply add more hay to the haystack so we can find the needle
		if added == false {
			haystack[0] = haystack[0]+1
		}

		// set the haystack string to hay
		hay = make_hay(haystack)

		fmt.Println("Made hay:", hay)
	}

	fmt.Println("----------^---------------------")
	fmt.Println("Found the needle in the haystack")
	fmt.Println("--------------------------------")
	fmt.Println("It took",time.Since(started_at))
}

func make_hay(a []int) string {
	var buffer bytes.Buffer

	for i := 0; i < len(a); i++ {
		buffer.WriteString(characters[a[i]])
	}

	return buffer.String()
}