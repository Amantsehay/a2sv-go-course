package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)

	var student Student
	fmt.Print("Enter student's name: ")
	scanner.Scan()
	student.Name = scanner.Text()

	var numSubjects int
	for {
		fmt.Print("Enter number of subjects: ")
		scanner.Scan()
		numSubjects, _ = strconv.Atoi(scanner.Text())
		if numSubjects <= 0 {
			fmt.Println("Please enter a positive number for the number of subjects.")
		} else {
			break
		}
	}

	student.Subjects = make(map[string]float64)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter subject %d (e.g., 'Mathematics'): ", i+1)
		scanner.Scan()
		subject := scanner.Text()

		var grade float64
		for {
			fmt.Printf("Enter grade for subject %s (0 to 100): ", subject)
			scanner.Scan()
			gradeInput := scanner.Text()
			grade, err := strconv.ParseFloat(gradeInput, 64)
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
