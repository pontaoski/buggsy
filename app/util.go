package app

import "encoding/json"

func debugDump(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "\t")
	println(string(data))
}
