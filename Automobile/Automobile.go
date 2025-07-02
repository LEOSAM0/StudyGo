package automobile

import "fmt"

type automobile struct {
	name         string
	isStarted    bool
	acceleration float64
	speed        float64
	safety       int
}

func NewAutomobile(Name string, IsStarted bool, Acceleration float64, Speed float64, Safety int) automobile {
	fmt.Println("Mark")
	if Name == "" {
		return automobile{}
	}

	fmt.Println("Started")
	if !IsStarted {
		return automobile{}
	}

	fmt.Println("Acceleration")
	if Acceleration <= 0 {
		return automobile{}
	}

	fmt.Println("Speed")
	if Speed < 0 || Speed > 250 {
		return automobile{}
	}

	fmt.Println("Safety")
	if Safety < 0 || Safety > 10 {
		return automobile{}
	}

	return automobile{
		name:         Name,
		isStarted:    IsStarted,
		acceleration: Acceleration,
		speed:        Speed,
		safety:       Safety,
	}
}

func (A *automobile) GetName() string {
	return A.name
}

func (A *automobile) ChangeStarted() {
	A.isStarted = true
}

func (A *automobile) ChangeAcceleration(chAc float64) {
	if chAc > 0 {
		A.acceleration = chAc
	}
}

func (A *automobile) ChangeSpeed(sp float64) {
	if sp >= 0 || sp < 250 {
		A.speed = sp
	}
}

func (A *automobile) ChangeSafety(Safe int) {
	if Safe >= 0 || Safe <= 10 {
		A.safety = Safe
	}
}
