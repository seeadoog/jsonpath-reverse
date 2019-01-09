package jsonref

import (
	"errors"
	"strconv"
	"strings"
	"fmt"
)

type QueryProp struct {
	Query string
	Value interface{}
}

const (
	TYPE_KEY=-1
)

// src must be map[string]interface{}
// Marshal() do not support expression start with  $,$[0] .
//if you need the function .please use MarshalInterface()
func Marshal(query string,src interface{}, value interface{})error{
	return marshal(query,src,value,0)
}

func marshal(query string,src interface{}, value interface{},start int) error {
	tks,err:=tokenize2(query)
	if err !=nil{
		return err
	}
	var cp = src
	return parserToken(tks,cp,value)
}

func Marshals(querys []QueryProp)(tmp map[string]interface{} ,err error){
	m:=map[string]interface{}{}
	for _,v:=range querys{
		if err:=Marshal(v.Query,m,v.Value);err!=nil{
			return nil,err
		}
	}
	return m,nil

}


func yyp(token string)(string,int,error){
	numidx_start:=0
	numidx_end:=0
	for k,v:=range token{
		t:=string(v)
		if t=="["{
			numidx_start=k
		}
		if t=="]"{
			numidx_end=k
		}
	}
	if numidx_end>0 && numidx_start>=0{
		num,err:=strconv.Atoi(token[numidx_start+1:numidx_end])
		if err !=nil{
			return "",TYPE_KEY,err
		}
		return token[:numidx_start],num,nil
	}
	return token,TYPE_KEY,nil
}


func tokenize2(query string) ([]string, error){
	return strings.Split(strings.Trim(query,"$."),"."),nil
}
//marshal and set the value to interface{}
//MarshalInterface() is power than Marshal().
// MarshalInterface() support expression such as $ ,$[0] which  Marshal() doesn't support
//attention that src must be *interface{}
func MarshalInterface(query string,src interface{}, value interface{}) error {
	return marshalInterface(query,src,value)
}

func marshalInterface(query string,src interface{}, value interface{}) error {
	//log.Println(reflect.TypeOf(src))
	tks,err:=tokenize2(query)
	if err !=nil{
		return err
	}
	var cp = src
	if cpi,ok:=src.(*interface{});ok{
		if strings.Trim(query,"$.")==""{
			*cpi=value
			return nil
		}
		if _,ok:=(*cpi).(map[string]interface{});ok{

		}else if _,ok:=(*cpi).([]interface{});ok{
			cp = cpi
			goto done
		}else{
			yp,idx,err:=yyp(tks[0])
			if err !=nil{
				return err
			}
			if idx==TYPE_KEY{
				*cpi = map[string]interface{}{}

			}else{
				if yp!=""{
					return errors.New("expression error: root array name must be empty,now is "+yp)
				}
				*cpi = make([]interface{},idx+1)
				cp = cpi
				goto done
			}
			//cp=*cpi

		}
		cp = *cpi

	}
done:
	return parserToken(tks,cp,value)

}
func parserToken(tks []string,cp,value interface{})error  {
	for k,v:=range tks{
		field,idx,err:=yyp(v)
		if err!=nil{
			return err
		}
		if k < len(tks) {
			//map
			if idx == TYPE_KEY {
				cpm, ok := cp.(map[string]interface{})
				if !ok {
					return errors.New(fmt.Sprintf("create field failed ,%s.parent cannot convert_ to map", field))
				}

				if k < len(tks)-1 {
					if cpm[field] == nil {
						cpm[field] = map[string]interface{}{}
					}
					cpm, ok = cpm[field].(map[string]interface{})
					if !ok {
						return errors.New(fmt.Sprintf("create field failed ,%s cannot convert_ to map", field))
					}
					cp = cpm
				} else {
					cpm[field] = value
				}
			} else { //array
				if field == "" && k ==0{
					cpi, ok := cp.(*interface{})
					if !ok {
						return errors.New("root is not pointer")
					}
					if _,ok:=(*cpi).([]interface{});!ok{
						return errors.New("root is not array")
					}
					if len((*cpi).([]interface{})) < idx+1 {
						for i := len((*cpi).([]interface{})); i < idx+1; i++ {
							*cpi = append((*cpi).([]interface{}), nil)
						}
					}
					if k < len(tks)-1 {
						for i := 0; i < idx+1; i++ {
							if (*cpi).([]interface{})[i] == nil {
								(*cpi).([]interface{})[i] = map[string]interface{}{}
							}
						}
						cp = (*cpi).([]interface{})[idx]
					} else {
						(*cpi).([]interface{})[idx] = value
					}
					//fmt.Println((*cpi).([]interface{}))
					continue
				}
				cpm, ok := cp.(map[string]interface{})
				if !ok {
					return errors.New("nil array child")
				}
				if cpm[field] == nil {
					cpm[field] = make([]interface{}, idx+1)
				}
				cps, ok := cpm[field].([]interface{})
				if !ok {
					return errors.New(fmt.Sprintf("create array failed ,%s cannot convert2 to array", field))
				}
				lenmap := len(cps)
				if lenmap < idx+1 {
					for i := lenmap; i < idx+1; i++ {
						cpm[field] = append(cpm[field].([]interface{}), nil)
					}
				}
				if k < len(tks)-1 {
					for i := 0; i < idx+1; i++ {
						if cpm[field].([]interface{})[i] == nil {
							cpm[field].([]interface{})[i] = map[string]interface{}{}
						}
					}
					cp = cpm[field].([]interface{})[idx]
				} else {
					cpm[field].([]interface{})[idx] = value
				}
			}
		}

	}
	return nil
}
//
func tokenize(query string) ([]string, error) {
	tokens := []string{}
	//	token_start := false
	//	token_end := false
	token := ""

	// fmt.Println("-------------------------------------------------- start")
	for idx, x := range query {
		token += string(x)
		// //fmt.Printf("idx: %d, x: %s, token: %s, tokens: %v\n", idx, string(x), token, tokens)
		if idx == 0 {
			if token == "$" || token == "@" {
				tokens = append(tokens, token[:])
				token = ""
				continue
			} else {
				return nil, fmt.Errorf("should start with '$'")
			}
		}
		if token == "." {
			continue
		} else if token == ".." {
			if tokens[len(tokens)-1] != "*" {
				tokens = append(tokens, "*")
			}
			token = "."
			continue
		} else {
			// fmt.Println("else: ", string(x), token)
			if strings.Contains(token, "[") {
				// fmt.Println(" contains [ ")
				if x == ']' && !strings.HasSuffix(token, "\\]") {
					if token[0] == '.' {
						tokens = append(tokens, token[1:])
					} else {
						tokens = append(tokens, token[:])
					}
					token = ""
					continue
				}
			} else {
				// fmt.Println(" doesn't contains [ ")
				if x == '.' {
					if token[0] == '.' {
						tokens = append(tokens, token[1:len(token)-1])
					} else {
						tokens = append(tokens, token[:len(token)-1])
					}
					token = "."
					continue
				}
			}
		}
	}
	if len(token) > 0 {
		if token[0] == '.' {
			token = token[1:]
			if token != "*" {
				tokens = append(tokens, token[:])
			} else if tokens[len(tokens)-1] != "*" {
				tokens = append(tokens, token[:])
			}
		} else {
			if token != "*" {
				tokens = append(tokens, token[:])
			} else if tokens[len(tokens)-1] != "*" {
				tokens = append(tokens, token[:])
			}
		}
	}
	// fmt.Println("finished tokens: ", tokens)
	// fmt.Println("================================================= done ")
	return tokens, nil
}