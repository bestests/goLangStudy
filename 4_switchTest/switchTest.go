package switchTest

func SwitchTest1(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return false
	default:
		return false
	}
}

func SwitchTest2(age int) bool {
	// case에 조건 넣을 수 있음
	switch {
	case age < 18:
		return false
	case age >= 18:
		return true
	case age >= 50:
		return false
	}
	return false
}

func SwitchTest3(age int) bool {
	// switch 문에 변수 선언해 사용 가능
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func SwitchTest4(age int) bool {
	// switch 문에 변수 선언 후 case 조건 처리 가능
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return false
}
