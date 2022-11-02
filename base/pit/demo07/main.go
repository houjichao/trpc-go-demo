package main

import (
	"encoding/json"
	"fmt"
)

/*
	不导出的 struct 字段无法被 encode
	以小写字母开头的字段成员是无法被外部直接访问的，
	所以 struct 在进行 json、xml、gob 等格式的 encode 操作时，这些私有字段会被忽略，导出时得到零值
*/
func main() {
	person := Person{"jack",22}

	fmt.Printf("%#v\n", person) //main.Person{Name:"jack", age:22}

	fmt.Println(person) //{jack 22}

	encoded, _ := json.Marshal(person)
	fmt.Println(string(encoded))    // {"Name":"jack"}    // 私有字段 age 被忽略了

	var out Person
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) //main.Person{Name:"jack", age:0}

}
