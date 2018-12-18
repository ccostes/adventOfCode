/*
	Input: 	text file containing one integer per line
	Output: sum of all lines, along with the first repeated 
			intermediate value
*/

package main
	
import (
    "bufio"
    "strconv"
    "fmt"
    "log"
    "os"
)

func main(){
	// open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// close the file when we're done
	defer file.Close()

	// use scanner to read file line by line and add to sum
	var sum, repeat int64 = 0, 0	// intermediate value and repeated value
	var foundRepeat bool = false
	var m map[int64]int

	m = make(map[int64]int)

	for i := 0; !foundRepeat && i < 1000; i += 1 {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			i, err := strconv.ParseInt(scanner.Text(), 0, 64)
			if err != nil {
				log.Fatal(err)
			}
			// calculate new intermediate value
			sum += i
			// fmt.Printf("%d ",sum)

			// try to add it to the map and see if it's a repeat
			if !foundRepeat {
				_, present := m[sum]
				// fmt.Printf("present: %v\n", present)
				if present {
					// found repeat
					foundRepeat = true
					repeat = sum
				} else {
					// add new value to the map
					m[sum] = 1
				}
			}
		}
		// reset file to beginnnig if we haven't found anything
		file.Seek(0,0)
	}

	// fmt.Println(m)

	fmt.Printf("Sum: %d\n", sum)
	if foundRepeat {
		fmt.Printf("Repeat: %d\n", repeat)
	} else {
		fmt.Println("No repeat\n")
	}
}