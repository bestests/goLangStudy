package structs

import (
	"fmt"
)

// type struct 선언
type person struct {
	name    string
	age     int
	hobbies []string
}

// Structs type and structs test
func Structs() {
	/* Structs */
	hobbyArr := []string{"movie", "music"}

	// 순서대로 넣으면 매핑 됨
	person1 := person{"kim", 18, hobbyArr}

	fmt.Println(person1)

	hobbyArr = []string{"book", "walk"}

	// 순서 없이 변수명으로 매핑해서 넣을 수 있음.
	person2 := person{age: 20, hobbies: hobbyArr, name: "lee"}

	fmt.Println(person2)

	// Array, Slice에서 해당 타입으로 사용가능
	personArr := []person{person1, person2}

	fmt.Println(personArr)

	// Map도 가능
	personMap := map[string]person{"person1": person1, "person2": person2}

	fmt.Println(personMap)

	person3 := person{"park", 21, []string{"test1", "test2"}}

	fmt.Println(person3)

	// type변수명.변수명으로 데이터 접근 가능
	fmt.Println(person3.hobbies[0])
}
