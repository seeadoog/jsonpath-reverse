package jsonref

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
)

type Lex struct {
	tokens []string
	lexss string
}

const (
	TYPE_KEY=-1
)

func MarshalToJson(query string,src map[string]interface{},dst interface{}) ([]byte,error) {
	err:= Marshal(query,src,dst)
	if err !=nil{
		return nil,err
	}

	return json.Marshal(src)

}

func Marshal(query string,src map[string]interface{},dst interface{}) error {
	tks,err:=tokenize(query)
	if err !=nil{
		return err
	}
	var cp map[string]interface{} = src
	for k,v:=range tks{
		if k==0{
			continue
		}
		field,idx,err:=yyp(v)
		if err!=nil{
			return err
		}
		if k< len(tks)-1{
			//field
			if idx ==TYPE_KEY{
				if cp[field]==nil{
					cp[field]= map[string]interface{}{}
				}
				cp = cp[field].(map[string]interface{})
			}else{   //array
				if cp[field]==nil{
					cp[field]=make([]map[string]interface{},idx+1)
				}
				//arrmp:=cp[field].([]map[string]interface{})
				lenmap:=len(cp[field].([]map[string]interface{}))
				if lenmap<idx+1{
					for i:=lenmap;i<idx+1;i++{
						cp[field]=append(cp[field].([]map[string]interface{}),map[string]interface{}{})
					}
				}
				for i:=0;i<idx+1;i++{
					if cp[field].([]map[string]interface{})[i]==nil{
						cp[field].([]map[string]interface{})[i]=map[string]interface{}{}
					}
				}
				cp = cp[field].([]map[string]interface{})[idx]
			}
		}else{
			if idx ==TYPE_KEY{
				cp[field]=dst
			}else{
				//todo
				if cp[field]==nil{
					cp[field]=make([]interface{},idx+1)
				}
				if len(cp[field].([]interface{}))<idx+1{
					for i:=len(cp[field].([]interface{}));i<idx+1;i++{
						cp[field]= append(cp[field].([]interface{}),1)
					}
				}
				cp[field].([]interface{})[idx]=dst

			}
		}


	}
	return nil

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
	if numidx_end>0 && numidx_start>0{
		num,err:=strconv.Atoi(token[numidx_start+1:numidx_end])
		if err !=nil{
			return "",TYPE_KEY,err
		}
		return token[:numidx_start],num,nil
	}
	return token,TYPE_KEY,nil
}


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