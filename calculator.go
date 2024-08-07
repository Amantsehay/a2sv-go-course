package main

import (
	"fmt"
	"log"
)

type Student struct {
	Name     string
	Subjects map[string]float64
}

func calculateAverage(grades map[string]float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var student Student
	fmt.Print("Enter student's name: ")
	fmt.Scanln(&student.Name)

	var numSubjects int
	for {
		fmt.Print("Enter number of subjects: ")
		_, err := fmt.Scanln(&numSubjects)
		if err != nil || numSubjects <= 0 {
			fmt.Println("Please enter a positive number for the number of subjects.")
		} else {
			break
		}
	}

	student.Subjects = make(map[string]float64)

	for i := 0; i < numSubjects; i++ {
		var subject string
		fmt.Printf("Enter subject %d: ", i+1)
		fmt.Scanln(&subject)

		var grade float64
		for {
			fmt.Printf("Enter grade for subject %s (0 to 100): ", subject)
			_, err := fmt.Scanln(&grade)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			if grade < 0 || grade > 100 {
				fmt.Println("Grade must be between 0 and 100.")
			} else {
				break
			}
		}

		student.Subjects[subject] = grade
	}

	average := calculateAverage(student.Subjects)
	fmt.Printf("The average grade for %s is %.2f\n", student.Name, average)
}
