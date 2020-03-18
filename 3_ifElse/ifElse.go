package ifElse

func CanIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func CanIDrink2(age int) bool {
	if age < 18 {
		return false
	}

	return true
}

func CanIDrink3(age int) bool {
	// if문 안에 변수 선언 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}
