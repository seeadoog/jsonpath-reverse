package jsonref

import (
	"errors"
)

func Lookup(query string,val interface{}) (interface{},error )  {
	tokens,err:=tokenize(query)
	if err !=nil{
		return nil,err
	}
	//log.Println(tokens)
	var cp = val
	for k,v:=range tokens{
		if k==0{
			continue
		}
		field,idx,err:=yyp(v)
		//log.Println(yyp(v))
		if err !=nil{
			return nil,err
		}
		if k< len(tokens)-1{
			if idx == TYPE_KEY{
				cpm,ok:=cp.(map[string]interface{})
				if !ok{
					return nil,errors.New("cannot convert to map")
				}
				cp = cpm[field]
			}else{
				if field==""{
					cps,ok:=cp.([]interface{})
					if !ok{
						return nil,errors.New("cannot convert to interface{}")
					}
					cp = cps[idx]
					continue
				}
				cpm,ok:=cp.(map[string]interface{})
				if !ok{
					return nil,errors.New("cannot convert to map")
				}
				//log.Println(cpm[field],reflect.TypeOf(cpm[field]))
				 cps,ok:=cpm[field].([]interface{})
				 if !ok{
				 	return nil,errors.New("caonnot convert to []map-")
				 }
				 cp =cps[idx]
			}
		}else{
			if idx == TYPE_KEY{

				cpm,ok:=cp.(map[string]interface{})
				if !ok{
					return nil,errors.New("cannot convert to map.")
				}
				return cpm[field],nil
			}else{
				if field==""{
					cps,ok:=cp.([]interface{})
					if !ok{
						return nil,errors.New("cannot convert to interface{}")
					}
					return cps[idx],nil
				}
				cpm,ok:=cp.(map[string]interface{})
				if !ok{
					return nil,errors.New("cannot convert to map.")
				}
				cps,ok:=cpm[field].([]interface{})
				if !ok{
					return nil,errors.New("cannot convert to []interface{}")
				}
				return cps[idx],nil
				return nil,nil
			}
		}

	}
	return nil,nil
}