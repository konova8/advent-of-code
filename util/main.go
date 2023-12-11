package util

type Position struct {
	X int
	Y int
}

func StrReverse(s string) string {
	var ret string
	for _, v := range s {
		ret = string(v) + ret
	}
	return ret
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	return lcm(numbers[0], numbers[1], numbers[2:]...)
}

func GCD(numbers ...int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	D := gcd(numbers[0], numbers[1])
	newNumbers := []int{D}
	newNumbers = append(newNumbers, numbers[2:]...)
	return GCD(newNumbers...)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func ABS[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
