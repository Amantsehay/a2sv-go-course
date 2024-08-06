package main

import (
	"fmt"
	"log"

)

type Student struct {
	Name string
	Subjects map[string]float64
}

func calculateAverage(grades map[string]float64) float64{
	var sum float64
	for _, grade := range grades{
		sum += grade
	}
	return sum / float64(len(grades))
}

func main(){

	var student Student
	fmt.Print("Enter student's name: ")
	fmt.Scanln(&student.Name)

	var numSubjects int
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numSubjects)

	student.Subjects = make(map[string]float64)

	for i := 0; i < numSubjects; i++{
		var subject string
		fmt.Printf("Enter subject %d: ", i+1)
		fmt.Scanln(&subject)

		var grade float64
		fmt.Printf("Enter grade for subject %s: ", subject)
		_, err := fmt.Scanln(&grade)
		if err != nil{
			log.Fatal(err)
		}

		student.Subjects[subject] = grade
	}
	var average float64 = calculateAverage(student.Subjects)
	fmt.Printf("The average grade for %s is %.2f\n", student.Name, average)


}