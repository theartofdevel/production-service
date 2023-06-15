package bytes

import (
	"encoding/json"
	"fmt"
)

func MapToByteArr(m map[string]interface{}) (value []byte, err error) {
	empData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return empData, nil
}

func ByteArrToMap(j string) (value map[string]interface{}, err error) {
	jsonMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(j), &jsonMap)
	if err != nil {
		return nil, err
	}

	dumpMap("", jsonMap)

	return jsonMap, nil
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}
