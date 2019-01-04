package jsonref

import (
	"testing"
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func TestLoad2(t *testing.T) {
	var m = map[string]interface{}{}
	Marshal("$.zhansan",m,User{"zhangsan",20})
	Marshal("$.class[0]",m,User{"lisi",11})
	Marshal("$.class[1]",m,User{"wangwu",18})
	Marshal("$.class[2]",m,User{"dajj",18})
	Marshal("$.group[5].name",m,"biaoge")
	Marshal("$.group[5].age",m,12)

	s,_:=json.Marshal(m)
	fmt.Println(string(s))
}

func Test_yy(t *testing.T)  {
	var m = map[string]interface{}{}
	Marshal("$.ssd.sd",m,1)
	Marshal("$.ssd.gf",m,1)
	Marshal("$.sdf.fg",m,1)
	Marshal("$.fg",m,1)
	Marshal("$.sdfg[1].gh",m,1)
	Marshal("$.sdfg[2].fg",m,1)
	Marshal("$.sdf2[2].as",m,1)
	Marshal("$.sdfl[0]",m,3)
	s,_:=json.Marshal(m)
	fmt.Println(string(s))
}

func TestLoad(t *testing.T) {
	var a []string
	var b =a
	b=append(b,"sdf")
	fmt.Println(a)
}