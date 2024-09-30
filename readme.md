#### josnref for go

**usage**

 jsonpath-reverse  is a tool used for creating jsonpath reversed. with this tool ,you can create a json string by specific query string like "$.user.name" 
 and set a value like  'bob' to it. The tool will generate a json string such as follow:
 ```text
{
    "user":{
        "name":"bob"
    }
}
```  

**get start**

- install
```text
    go get -t github.com/seeadoog/jsonpath-reverse
``` 

- use 
```text
import (
	"fmt"
	"encoding/json"
	jsonref "github.com/skyniu/jsonpath-reverse"
)
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	jp, err := Compile("a.name[0]")
	if err != nil {
		panic(err)
	}

	var a any
	err = jp.Set(&a, 1)
	if err != nil {
		panic(err)
	}
	ns, _ := json.MarshalIndent(a, "", "    ")
	fmt.Println(string(ns))
}

```
the output json is :
```text
{
    "a": {
        "name": [
            1
        ]
    }
}
```
