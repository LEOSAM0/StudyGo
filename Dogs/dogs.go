package dogs

type Dog struct {
	Name       string
	JudgeGrade int
	IsPolitely bool
}

// *dogs.Dog
func (D *Dog) IfPolitely() int {
	if D.IsPolitely {
		D.JudgeGrade += 1
		return D.JudgeGrade
	}
	return D.JudgeGrade
}
