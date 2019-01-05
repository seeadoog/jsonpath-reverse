package jsonref

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
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
	log.Println(Marshal("$.class[5].name",m,"biaoge"))
	log.Println(Marshal("$.class[5].age",m,23))
	Marshal("$.group[5].age",m,12)
	Marshal("$.group[5].son.son.name",m,"bgnb")
	Marshal("$.group[5].son.son.age",m,33)

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

func TestMarshals(t *testing.T) {
	s:=time.Now()
	marshal()
	fmt.Println("time",time.Since(s).Nanoseconds())
}

func marshal()  {
	_,err:=Marshals([]QueryProp{
		{"$.biaoge.name","biaoge"},
		//{"$.biaoge.say","bgnb"},
		//{"$.dajj.name","dajj"},
		//{"$.dajj.say","dajj niubi"},
		//{"$.group[0]","biaoge"},
		//{"$.group[1]","dajj"},
		//{"$.group[2]","hg"},
		////{"$.group[2].gg","hg"},
		//{"$.less[0].hgfs.had.pg[0].hhh[1].fs","hg"},
	})
	if err !=nil{
		fmt.Println(err)
		return
	}
	//s,_:=json.Marshal(tmp)
	//fmt.Sprintf(string(s))
}
func Test_bench(t *testing.T)  {
	for i:=0;i<1000000;i++{
		marshal()
	}
}

func TestNI(t *testing.T) {
	var m = map[string]interface{}{}
	log.Println(Marshal("$.s[0]",m,1))
	log.Println(Marshal("$.s[1]",m,2))
	log.Println(Marshal("$.s[2]",m,3))
	s,_:=json.Marshal(m)
	fmt.Println(string(s))
}