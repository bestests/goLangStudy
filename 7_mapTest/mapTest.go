package mapTest

import (
	"fmt"
)

func MapTest () {
	/* Map */

	// map 선언 - 변수명 := map[key type]value type{key:value,....}
	mapTest := map[string]string{"name": "song", "age": "10"}

	fmt.Println(mapTest)

	// map 사용 - 변수명[key]
	fmt.Println(mapTest["name"])

	mapTest["name"] = "test"

	fmt.Println(mapTest["name"])

	// for문으로 map 돌리기
	for key, value := range mapTest {
		fmt.Println("key : '", key, "', value : '", value, "'")
	}
}