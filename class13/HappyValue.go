package class13

import "math"

type Employee struct {
	HappyValue   int
	Subordinates []Employee
}

type Info struct {
	LaiMaxHappy int
	BuMaxHappy  int
}

func GetHappyValue(head Employee) int {
	hv := processHV(head)
	return int(math.Max(float64(hv.LaiMaxHappy), float64(hv.BuMaxHappy)))
}

func processHV(head Employee) Info {
	if len(head.Subordinates) == 0 {
		return Info{head.HappyValue, 0}
	}

	lai := head.HappyValue
	bu := 0

	for _, v := range head.Subordinates {
		hv := processHV(v)
		lai += hv.LaiMaxHappy
		bu += int(math.Max(float64(hv.BuMaxHappy), float64(hv.LaiMaxHappy)))
	}
	return Info{lai, bu}
}
