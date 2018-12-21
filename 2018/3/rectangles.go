/*

	Rectangles

	We're given a list of rectangles (position and size) and need to figure out 
	how much area is covered by more than one.

	Use two arrays, one for covered area and one for overlap area. For each shape,
	fill in the cover array and, for any area that is already filled in, fill in
	the corresponding area on the overlap array. Then just measure the area filled
	on the overlap array.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"regexp"
	"strconv"
)

// struct representing a rectangle
type rect struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func main() {
	// first load the input shapes into memory
	var shapes []rect
	var island map[uint]bool
	island = make(map[uint]bool)

	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Each line looks like: #1 @ 179,662: 16x27
		// first split off the front part
		parts := strings.Split(scanner.Text(), "@")
		
		// parts looks like [#1   179,662: 16x27]
		// use regex to extract all numbers from this
		re := regexp.MustCompile("[0-9]+")
		data := re.FindAllString(parts[1], -1)
		// nums looks like [179 662 16 27]
		// validate
		if len(data) == 4{
			// convert to int
			var nums [4]int
			for i,v := range data {
				nums[i], _ = strconv.Atoi(v)
			}
			// add to shapes list
			shapes = append(shapes, rect{nums[0], nums[0] + nums[2] - 1, nums[1], nums[1] + nums[3] - 1})
			// add to island map
			island[uint(len(shapes))] = true
		} else {
			log.Print("Error parsing line: %v", scanner.Text())
		}
	}

	// we now have an array of shapes, start adding to cover and overlap arrays
	// start with a default size of 1000x1000 and throw error if that's too small
	const size = 1100
	var arrCover [size][size][]uint

	for i, shape := range shapes {
		if shape.x2 > size || shape.y2 > size {
			fmt.Printf("Shape %v exceeds array bounds!", i)
			os.Exit(1)
		}
		// fmt.Printf("Shape %v (%v,%v)(%v,%v)\n", i, shape.x1, shape.y1, shape.x2, shape.y2)
		for x := shape.x1; x <= shape.x2; x++ {
			for y := shape.y1; y <= shape.y2; y++ {
				arrCover[y][x] = append(arrCover[y][x], uint(i+1))
			}
		}
	}
	// fmt.Println(shapes)
	// fmt.Println(arrCover)
	// fmt.Println(arrOver)

	// now count how much area is in the overlap array
	// count := 0
	// for _,a := range arrCover {
	// 	str := ""
	// 	for _,v := range a {
	// 		str = str + strconv.Itoa(int(v))
	// 	}
	// 	fmt.Println(str)
	// }

	// Find the amount of overlap
	// for _,a := range arrCover {
	// 	for _,v := range a {
	// 		if v == ^uint(0) {
	// 			count++
	// 		}
	// 	}
	// }
	// fmt.Printf("Count: %v",count)
	
	// find a space with only on entry
	for _,a := range arrCover {
		for _,v := range a {
			if len(v) > 1 {
				// delete all from map
				for _,s := range v {
					delete(island, s)
				}
			}
		}
	}
	fmt.Println(island)

}