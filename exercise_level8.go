package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

type specialAgent struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ByAge []specialAgent

type ByLast []specialAgent

func main() {
	/*
	   Exercise 1
	   Given the code below, marshal the []user to JSON.
	*/
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	// your code goes here
	xb, err := json.Marshal(users)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(string(xb))

	/*
	   Exercise 2
	   Given the code below, unmarshal the JSON into a Go data structure.
	*/
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	// your code goes here
	var people []user
	json.Unmarshal([]byte(s), &people)

	fmt.Println(people)

	fmt.Printf("\v")

	/*
	   Exercise 3
	   Given the code below, encode to JSON the []user sending the results to Stdout.
	*/
	sa1 := specialAgent{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	sa2 := specialAgent{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	sa3 := specialAgent{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	agents := []specialAgent{sa1, sa2, sa3}

	fmt.Println(agents)

	// your code goes here
	sb, e := json.Marshal(agents)
	if err != nil {
		fmt.Printf("Error: %v", e)
	}

	fmt.Println(string(sb))

	// json.NewEncoder(os.stdout) returns a *Encoder, which has a method Encode()
	json.NewEncoder(os.Stdout).Encode(sb)

	fmt.Printf("\v")

	/*
	   Exercise 4
	   Given the code below, sort the []int and []string for each person
	*/
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	// your code goes here

	// sort xi
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	// sort xs
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Printf("\v")

	/*
	   Exercise 5
	   1. Given the code in exercise 3, sort the []specialAgents by:
	       a. age
	       b. last
	   2. Also sort each []string "Sayings" for each user.
	   3. Print everything out using the range clause and in a way that is pleasant.
	*/

	fmt.Println(agents)

	// your code goes here
	sort.Sort(ByAge(agents))
	PrintArr(agents, "By age:")

	sort.Sort(ByLast(agents))
	PrintArr(agents, "By last:")
}

// These three funcs are used to sort by age
func (agents ByAge) Len() int {
	return len(agents)
}
func (agents ByAge) Swap(i, j int) {
	agents[i], agents[j] = agents[j], agents[i]
}
func (agents ByAge) Less(i, j int) bool {
	return agents[i].Age < agents[j].Age
}

// These three funcs are used to sort by last
func (agents ByLast) Len() int {
	return len(agents)
}
func (agents ByLast) Swap(i, j int) {
	agents[i], agents[j] = agents[j], agents[i]
}
func (agents ByLast) Less(i, j int) bool {
	return agents[i].Age < agents[j].Age
}

func PrintArr(agents []specialAgent, order string) {
	fmt.Println(order)
	for _, sa := range agents {
		fmt.Printf("\tFirst: %v\n", sa.First)
		fmt.Printf("\tLast: %v\n", sa.Last)
		fmt.Printf("\tAge: %v\n", sa.Age)
		fmt.Printf("\tSayings:\n")
		for _, v := range sa.Sayings {
			fmt.Printf("\t\t\t%v\n", v)
		}
	}
}
