package sortarraytan

func SearchMax(origin []int) (maxNum int, maxNumIndex int) {
	maxNum = origin[0]
	maxNumIndex = 0
	originLen := len(origin)
	for i := 1; i < originLen; i++ {
		if maxNum < origin[i] {
			maxNum = origin[i]
			maxNumIndex = i
		}
	}
	return
}

func SearchMin(origin []int) (minNum int, minNumIndex int) {
	minNum = origin[0]
	minNumIndex = 0
	originLen := len(origin)
	for i := 1; i < originLen; i++ {
		if minNum > origin[i] {
			minNum = origin[i]
			minNumIndex = i
		}
	}
	return
}

func SearchNumber(origin []int, number int) (NumberIndex []int) {
	for i := 0; i < len(origin); i++ {
		if number == origin[i] {
			NumberIndex = append(NumberIndex, i)
		}
	}
	return
}

func IsExist(origin []int, number int) bool {
	for i := 0; i < len(origin); i++ {
		if number == origin[i] {
			return true
		}
	}
	return false
}
