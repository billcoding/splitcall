package splitcall

import "testing"

func TestCall(t *testing.T) {
	var list []interface{}
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 4)
	list = append(list, 4)
	list = append(list, 4)
	list = append(list, 4)
	list = append(list, 4)
	Call(list, func(subList []interface{}) {
		t.Log(len(subList))
	}, 4)

}
