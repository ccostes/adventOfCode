/*
	checksum

	given a list of character strings we want to compute a checksum that is A * B 
	where A is the number of strings with exactly two of any letter, and B is the
	number of strings with exactly three of any letter. 

	Example Input:
		abcdef contains no letters that appear exactly two or three times.
		bababc contains two a and three b, so it counts for both.
		abbcde contains two b, but no letter appears exactly three times.
		abcccd contains three c, but no letter appears exactly two times.
		aabcdd contains two a and two d, but it only counts once.
		abcdee contains two e.
		ababab contains three a and three b, but it only counts once.
	
	Of these box IDs, four of them contain a letter which appears exactly twice, 
	and three of them contain a letter which appears exactly three times. 
	Multiplying these together produces a checksum of 4 * 3 = 12.

*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main(){
	var repTwo, repThree int

	// open the input file
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	// close the file when we're done
	defer file.Close()

	// use scanner to go through the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		// check to see if this 
		switch testRepeat(scanner.Text()) {
		case 1:
			// 2 and 3 repeats
			repTwo += 1
			repThree += 1
		case 2:
			// 2 repeats
			repTwo += 1
		case 3:
			repThree += 1
		}
	}

	// compute checksum
	fmt.Printf("repTwo: %d repThree: %d Checksum: %d\n", repTwo, repThree, repTwo * repThree)
}

func testRepeat(str string) int {
	// given a string, check for repeated characters
	// if exactly 2 return 2, if 3 return 3, if both return 1
	// else return 0
	var m map[rune]int
	m = make(map[rune]int)

	// store character count into map
	for _, char := range str{
		val, present := m[char]

		if present{
			// increment value
			m[char] = val + 1
		} else {
			m[char] = 1
		}
	}

	// iterate map to see if we have 2 and/or 3 repeats
	foundTwo, foundThree := false, false
	for _,v := range m{
		if v == 2{
			foundTwo = true
		}
		if v == 3 {
			foundThree = true
		}
		if foundTwo && foundThree {
			// can return without checking the rest
			// fmt.Printf("1:%v\n",str)
			return 1
		}
	}
	if foundTwo {
		// fmt.Printf("2:%v\n",str)
		return 2
	} else if foundThree {
		// fmt.Printf("3:%v\n",str)
		return 3
	} else {
		return 0
	}
}