package main

import (
	"fmt"
	"os"
)

type Students struct {
	id      string
	name    string
	address string
	job     string
	reason  string
}

func main() {
	students := []Students{
		{
			id:      "1",
			name:    "Kazuha",
			address: "Inazuma",
			job:     "Software Engineer",
			reason:  "I want to learn a new programming language.",
		},
		{
			id:      "2",
			name:    "Xiao",
			address: "Liyue",
			job:     "Data Scientist",
			reason:  "I heard Go is good for concurrency.",
		},
		{
			id:      "3",
			name:    "Scara",
			address: "Sumeru",
			job:     "Web Developer",
			reason:  "I'm interested in building scalable web applications.",
		},
	}

	args := os.Args
	studentID := args[1]
	
	for _, student := range students {
		if student.id == studentID {
			fmt.Printf("Name: %+s\n", student.name)
			fmt.Printf("Address: %+s\n", student.address)
			fmt.Printf("Job: %+s\n", student.job)
			fmt.Printf("Reason: %+s\n", student.reason)
			return
		}
	}

	fmt.Println("Student not found.")
}
