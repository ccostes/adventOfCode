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
	var sum int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		i, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		
		sum += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %d\n", sum)
}