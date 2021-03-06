package jsonref

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func TestLoad2(t *testing.T) {
	var m = map[string]interface{}{}
	Marshal("$.zhansan",m,User{"zhangsan",20})
	Marshal("$.zhan[0]",m,1)
	Marshal("$.zhan[1]",m,2)
	log.Println(Marshal("$.class[5].name",m,"biaoge"))
	Marshal("$.class[0]",m,User{"lisi",11})
	Marshal("$.class[1]",m,User{"wangwu",18})
	Marshal("$.class[2]",m,User{"dajj",18})
	log.Println(Marshal("$.class[3].age",m,23))
	Marshal("$.group[5].age",m,12)
	Marshal("$.group[5].son.son.name",m,"bgnb")
	Marshal("$.group[5].son.son.age",m,33)
	Marshal("$.nii.sss.ggg.hhh.jjj.kkk.ll.sss.mmm.ggg",m,23)
    Marshal("$.nii.sss.ggg.hhh.jjj.kkk.ll.sss.mmm.ggg.ff",m,23)
    Marshal("$.nii.sss.ggs[1].hhh[0].jjj[0].kkk[1].ll[2].sss.mmm.ggg.ff[1].ss[0]",m,"12")
	s,_:=json.Marshal(m)
	fmt.Println(string(s))
}

func Test_yy(t *testing.T)  {

	var m = map[string]interface{}{}

	Marshal("$.result",m,map[string]interface{}{
		"status2":2,
	},)

	Marshal("$.result.status",m,1)

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
	//s:=time.Now()
	marshal1()
	//fmt.Println("time",time.Since(s).Nanoseconds())
}

func marshal1()  {
	tmp,err:=Marshals([]QueryProp{
		//{"$.biaoge.name.sss[0].sdg","biaoge"},
		{"$.result.status",1},
		{"$.result",map[string]interface{}{
			"status":2,
		}},
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
	s,_:=json.Marshal(tmp)
	fmt.Sprintf(string(s))
}
func Test_bench(t *testing.T)  {
	for i:=0;i<1000000;i++{
		marshal1()
	}
}

func TestNI(t *testing.T) {
	var m = map[string]interface{}{}
	log.Println(Marshal("$.s[0].a",m,1))
	log.Println(Marshal("$.s[1].a",m,2))
	log.Println(Marshal("$.s[2].a",m,3))
	log.Println(Marshal("$.s[5].a",m,3))
	s,_:=json.Marshal(m)
	fmt.Println(string(s))
}

func TestInterface(t *testing.T) {
	var i interface{}
	inter(&i)
	log.Println(i)
}

func inter(v interface{})  {
	i:=v.(*interface{})
	*i= map[string]interface{}{
		"kk":"dsf",
	}
	var b = *v.((*interface{}))
	b.(map[string]interface{})["sdf"]="sdf"
}

func TestMarshals2(t *testing.T) {
	var i interface{}
	marshalInterface("$",&i,map[string]interface{}{
		"keys":"234",
	})
	marshalInterface("$.mapss",&i,map[string]interface{}{
		"key":"2341",
	})
	marshalInterface("$.sdfs",&i,map[string]interface{}{
		"key":"2342",
	})
	marshalInterface("$.abc[0].bcd[1]",&i,map[string]interface{}{
		"key":"2343",
	})
	marshalInterface("$.abc[0].bcd[0]",&i,map[string]interface{}{
		"key":"2344",
	})
	b,_:=json.Marshal(i)
	fmt.Println(string(b))
}

func TestRoot(t *testing.T) {
	var j interface{}
	for i:=0;i<1000000;i++{
		marshalInterface("$.a.b.c",&j,map[string]interface{}{
			"hha":"hha",
		})
	}

	//marshalInterface("$.ha",&i,1)
	b,_:=json.Marshal(j)
	fmt.Println(string(b))
}

func TestArr(t *testing.T)  {

	var i interface{}
	marshalInterface("$[0].abc.gf[2]",&i,&User{})
	marshalInterface("$[0].abc.gf[1]",&i,&User{})
	marshalInterface("$[1].abc.sss[0].fff",&i,&User{"sdf",3})
	marshalInterface("$[5].abc.sss[0].fff",&i,&User{"sdf",3})
//	marshalInterface("$[1]",&i,&User{},-1)
	//marshalInterface("$[1].asss",&i,&User{},-1)
	b,_:=json.Marshal(i)
	fmt.Println(string(b))
}
func TestMap(t *testing.T) {
	var i interface{}
	marshalInterface("$.abcd",&i,[]string{"123","456"})
	marshalInterface("$.abc.gf[1]",&i,&User{})
	marshalInterface("$.abc.sss[0].fff",&i,&User{"sdf",3})
	marshalInterface("$.abc.sss[0].fff2",&i,&User{"sdf",3})
	//	marshalInterface("$[1]",&i,&User{},-1)
	//marshalInterface("$[1].asss",&i,&User{},-1)
	b,_:=json.Marshal(i)
	fmt.Println(string(b))

}

func TestSwitchJson(t *testing.T) {
data:=`
{
        "rlt": [
            {
                "sid": "F002_001.wav"
            },
            {
                "age": [
                    {
                        "middle": "0.3180",
                        "child": "0.4887",
                        "old": "0.1933",
                        "age_type": "1"
                    }
                ]
            },
            {
                "gender": [
                    {
                        "female": "0.5933",
                        "male": "0.4067",
                        "gender_type": "0"
                    }
                ]
            }
        ]
    }

`
    var datas interface{}
    json.Unmarshal([]byte(data),&datas)
    var res interface{}
    err:=SwitchJson([]SwitchExp{
    	{"$.a[0]","$.rlt[1].age[0]"},
    	{"$.a[1]","$.rlt[2].gender[0]"},
    	{"$.a[2]","$.rlt[0].sid"},
	},&res,datas)
    log.Println(err)
	//age,_:=Lookup("$.rlt[1].age[0]",its)
	//marshalInterface("$.result.age",&i,age)
	//gender,_:=Lookup("$.rlt[2].gender[0]",its)
	//marshalInterface("$.result.gender",&i,gender)
	b,_:=json.Marshal(res)
	fmt.Println(string(b))

}

func TestLookup4(t *testing.T) {
	var i interface{}
	marshalInterface("$.a.b",&i,1)
	marshalInterface("$.a.c.c",&i,2)

	log.Println(Lookup("$.a.b",i))
	log.Println(Lookup("$.a.c.c",i))
}

func TestMarshal(t *testing.T) {
	var i interface{}
	marshalInterface("$",&i,[]interface{}{"123456","123456"})
	marshalInterface("$[2]",&i,"ddd")
	b,_:=json.Marshal(i)
	log.Println(string(b))
}

type Value interface {}

type Map map[string]Value



func TestAbstractController_ArgsError(t *testing.T) {
	s:=`	
{
    "field": "$auf", 
    "op": "IN", 
    "val": [
        "audio/L16;rate=16000", 
        "audio/L16;rate=16k"
    ], 
    "result": {
        "match": {
            "field": "$rate", 
            "op": "SET", 
            "val": "16000"
        }
    }
}

`
	var m = Map{}
	json.Unmarshal([]byte(s),&m)
	log.Println(m)
}