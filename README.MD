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

Usage:<br/>
Calculates a student's final course grade based on weighted averages from Assignemnts (50%) Exams (35%) and Essays (15%)<br/>
AddGrade allows for grade addition<br/>
GetFinalGrade returns the letter grade<br/>
```
gradeCalculator := NewGradeCalculator()
gradeCalculator.AddGrade("open source assignment", 100, Assignment)
gradeCalculator.GetFinalGrade()
```
