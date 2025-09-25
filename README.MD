# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Clone repo:
```
git clone https://github.com/Aleena-J/esep-grade-calculator.git
cd esep-grade-calculator
```

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile coverage.txt ./...
```

Usage:
Calculates a student's final course grade based on weighted averages from Assignemnts (50%) Exams (35%) and Essays (15%)
AddGrade allows for grade addition
GetFinalGrade returns the letter grade
    Ex. 	gradeCalculator := NewGradeCalculator()
	        gradeCalculator.AddGrade("open source assignment", 100, Assignment)
            gradeCalculator.GetFinalGrade()