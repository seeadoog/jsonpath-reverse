package jsonref

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestLookup(t *testing.T) {
	var jsons = `
[{
	"bege":{
		"cp":[1,2,3],
		"cps":[
			{
			   "cp":[1,2,3],
				"cp1":[{"haha":"haha"}]
			}
			
		]
	}
}]

`
	//var i interface{}


	var i interface{}
	err := json.Unmarshal([]byte(jsons),&i)
	log.Println(err)
	fmt.Println(Lookup("$[0].bege.cps[0].cp1[0].haha",i))
	fmt.Println(Lookup("$[0].bege.cps[0].cp[0]",i))
	fmt.Println(Lookup("$[0].bege.cps[0].cp1[0].d",i))
}
