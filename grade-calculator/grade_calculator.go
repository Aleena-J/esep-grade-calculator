package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades,
		Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.getGrades(Assignment))
	exam_average := computeAverage(gc.getGrades(Exam))
	essay_average := computeAverage(gc.getGrades(Essay))

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func (gc *GradeCalculator) getGrades(gt GradeType) []Grade {
	var filteredGrades []Grade
	for _, grd := range gc.grades {
		if grd.Type == gt {
			filteredGrades = append(filteredGrades, grd)
		}
	}
	return filteredGrades
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, grd := range grades {
		sum += grd.Grade
	}

	return sum / len(grades)
}
