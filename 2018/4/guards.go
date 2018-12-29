

package main

import (
	"os"
	"log"
	"errors"
	"fmt"
	"time"
	"strings"
	"bufio"
	"regexp"
	"strconv"
)

/*
	Input is an unsorted file of timestamped log entries. First need to 
	sort the entries, so as we read them in we'll add them to a binary 
	search tree.
*/

func main() {
	// dates look like this 1518-07-18 23:57
	const form = "2006-01-02 15:04"
	tree := &Tree{}

	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// outFile, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outFile.Close()
	// w := bufio.NewWriter(outFile)

	// read each entry in and add to tree
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// [1518-11-01 00:00] Guard #10 begins shift
		parts := strings.Split(scanner.Text(), "]")
		// fmt.Println(parts[0][1:])
		t, err := time.Parse(form, parts[0][1:])
		// fmt.Println(t)
		if err == nil {
			// fmt.Printf("Inserting %v : %v\n", t, parts[1][1:])
			tree.Insert(t,parts[1][1:])
		} else {
			fmt.Printf("time parse err: %v\n", err)
		}
	}

	// tree.Traverse(tree.root, PrintNode)
	// sorted entries can now be retrieved by traversing the tree
	// find the guard who has the most minutes asleep
	if tree.root == nil {
		fmt.Println("Empty tree")
		return
	}

	guardTotal := make(map[int]int)
	guardMap := make(map[int] [60]int)
	currGuard := -1
	var sleepTime time.Time

	f := func(n *Node) {
		// fmt.Printf("%v %v\n", n.key, n.value)
		// see what kind of log event
		switch n.value[0:5]{
		case "Guard":
			// new guard; use regex to extract all numbers
			re := regexp.MustCompile("[0-9]+")
			data := re.FindAllString(n.value, -1)
			newGuard, _ := strconv.Atoi(data[0])
			// fmt.Printf("Found new guard %v\n", newGuard)
			currGuard = newGuard
			// fmt.Fprintf(w, "#%v\n", newGuard)
		case "falls": 
			// update sleepTime
			sleepTime = n.key
			// fmt.Printf("Guard went to sleep\n")
			// fmt.Fprintf(w, "[%v] %v\n", n.key.Format(form), n.value)
		case "wakes":
			// fmt.Printf("Guard %v sleep: %v wakes: %v time: %v\n", currGuard, sleepTime, n.key, n.key.Sub(sleepTime))
			// fmt.Printf("adding %v minuts of sleep for guard %v\n", duration, currGuard)

			// update minutes guard was asleep
			for m := sleepTime.Minute(); m < n.key.Minute(); m++ {
				slc := guardMap[currGuard]
				slc[m] += 1
				guardMap[currGuard] = slc
			}
			// update guard total sleep
			duration := (int)(n.key.Sub(sleepTime).Minutes())
			guardTotal[currGuard] += duration
			// fmt.Fprintf(w, "%v,%v\n", currGuard, duration)
		}
	}


	// fExportTree := func(n *Node) {
	// 	fmt.Fprintf(w, "[%v] %v\n", n.key.Format(form), n.value)
	// }

	// // traverse the log entries and tally minutes of sleep for each guard
	// tree.Traverse(tree.root, fExportTree)
	// w.Flush()
	// return

	tree.Traverse(tree.root, f)
	// w.Flush()
	// fmt.Println(guardMap)

	// iterate map to find guard with most sleep
	maxGuard := -1
	maxMin := -1
	guardCount := 0
	for k,v := range guardTotal {
		guardCount++
		if v > maxMin {
			maxMin = v
			maxGuard = k
		}
	}
	fmt.Printf("Found %v guards. Guard %v had the most sleep (%v miutes)\n", guardCount, maxGuard, maxMin)

	// now need to find which minute the guard was asleep most
	slc := guardMap[maxGuard]
	// fmt.Println(slc)
	maxMin = -1
	maxVal := -1
	for m := 0; m < 60; m++ {
		if slc[m] > maxVal {
			maxMin = m
			maxVal = slc[m]
		}
	}
	fmt.Printf("Guard slept most during minute %v Part 1 result: %v\n", maxMin, maxGuard * maxMin)

	// find the guard who was asleep most during the same minute
	maxMin = -1
	maxVal = -1
	guardCount = 0
	for k,v := range guardMap {
		fmt.Println(k,v)
		// iterate each minute to find max
		for m := 0; m < 60; m++ {
			if v[m] > maxVal {
				maxMin = m
				maxVal = v[m]
				guardCount = k
			}
		}
	}

	fmt.Printf("Guard %v slept most frequently (%v times) on minute %v\nPart 2 result: %v\n", guardCount, maxVal, maxMin, guardCount * maxMin)
}


// implementation of a binary search tree with times as keys and strings as values

type Node struct{
	key time.Time
	value string
	left *Node
	right *Node
}

// function to insert a node into a BST
func (n *Node)Insert(key time.Time, value string) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	// figure out  if this goes into the left or right subtree
	if key == n.key {
		fmt.Printf("Error: duplicate key %v", key)
		return nil
	}
	if key.Before(n.key) {
		// put this in the left subtree
		if n.left == nil {
			// add as left child
			n.left = &Node{key,value,nil,nil}
			return nil
		}
		//recurse down left subtree
		return n.left.Insert(key,value)
	} else {
		// insert in right subtree
		if n.right == nil {
			// add as left child
			n.right = &Node{key,value,nil,nil}
			return nil
		}
		//recurse down left subtree
		return n.right.Insert(key,value)
	}
}

func PrintNode(n *Node) {
	fmt.Printf("%v : %v\n", n.key, n.value)
}

type Tree struct {
	root *Node
}

func (t *Tree)Insert(key time.Time, value string) error {
	if t.root == nil {
		t.root = &Node{key, value, nil, nil}
		return nil
	}
	return t.root.Insert(key,value)
}

func (t *Tree)Traverse(n *Node, f func(*Node)){
	if n == nil {
		return
	}
	t.Traverse(n.left,f)
	f(n)
	t.Traverse(n.right,f)
}