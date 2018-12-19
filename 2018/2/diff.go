/*
	diff

	given a list of character strings we want to find two which differ by only 
	one letter, and return the characters that are common between those two

	Example Input:
		abcde
		fghij
		klmno
		pqrst
		fguij
		axcye
		wvxyz
	
	The IDs abcde and axcye are close, but they differ by two characters 
	(the second and fourth). However, the IDs fghij and fguij differ by exactly 
	one character, the third (h and u). Those must be the correct boxes.

	What letters are common between the two correct box IDs? (In the example above, 
	this is found by removing the differing character from either ID, producing fgij.)

*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "hash/fnv"
)

// We know that only one position will be differnt between the target pair. Iterate through
// each position, assuming that is the one that differs, and the target pair with be the two
// with the same sum
func main(){

	// open the input file
	file, err := os.Open("input3.txt")
	if err != nil{
		log.Fatal(err)
	}
	// close the file when we're done
	defer file.Close()

	// use scanner to go through the file line by line
	scanner := bufio.NewScanner(file)

	// find out how many positions we have to check
	scanner.Scan()
	length := len(scanner.Text())
	file.Seek(0,0)
	scanner = bufio.NewScanner(file)

	for i := 0; i < length; i++ {
		// assume position i is the one that differs, and sum all the other letters
		var m map[int]int
		// map[sum] = list entry
		m = make(map[int]int)
		entry := 0
		scanner = bufio.NewScanner(file)

		for scanner.Scan(){
			// compute hash of this entry, minus the rune at index i
			h := hash(scanner.Text()[:i] + scanner.Text()[i+1:])
			// see if we already found this sum
			v,present := m[h]
			if present {
				// found our pair!
				fmt.Printf("Found pair! index %v\n\t%v\n\t%v\n", i, v, scanner.Text())
				// print common text (this string minus char at position i)
				fmt.Printf("Found pair! Common: %v%v", scanner.Text()[:i], scanner.Text()[i+1:])
				return

			} else {
				m[h] = entry
			}

			// increment which entry we're working on
			entry++
		}
		// reset our file pointer
		file.Seek(0,0)
	}
	fmt.Println("No match found :(")
}

func hash(s string) int {
        h := fnv.New32a()
        h.Write([]byte(s))
        return int(h.Sum32())
}