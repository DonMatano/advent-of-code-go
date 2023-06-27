package advent_of_code_2020

func GetSantaFloorFromCode(data string) int {
	chars := []rune(data)
	floor := 0
	for _, c := range chars {
		if "(" == string(c) {
			floor += 1
		}
		if ")" == string(c) {
			floor -= 1
		}
	}
	return floor
}

func CheckFirstBasementPosition(data string) int {
	chars := []rune(data)
	var firstBasementPosition int
	floor := 0

	for i, c := range chars {
		if "(" == string(c) {
			floor += 1
		}
		if ")" == string(c) {
			floor -= 1
			if floor < 0 {
				firstBasementPosition = i + 1
				break
			}
		}
	}
	return firstBasementPosition
}
