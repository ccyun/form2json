package form2json

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strings"
)

type f2j struct {
	input []input
}
type input struct {
	keys  []string
	value string
}

type ders []interface{}
type derm map[string]interface{}

//Unmarshal form转json
func Unmarshal(s string, str interface{}) (string, error) {
	var (
		err error
		vs  url.Values
	)
	F := new(f2j)
	vs, err = url.ParseQuery(s)
	if err != nil {
		return "", err
	}
	for k, r := range vs {
		k = strings.Replace(k, "]", "", -1)
		F.input = append(F.input, input{
			keys:  strings.Split(k, "["),
			value: r[0],
		})
	}
	data := F.unmarshal([]string{}, 0)
	d, _ := json.Marshal(data)
	jsonStr := regexp.MustCompile(`\"([0-9]|[0-9]+(\.)+[0-9]+|[1-9]+[0-9]+|true|false)\"`).ReplaceAllString(string(d), "$1")
	if str != nil {
		err = json.Unmarshal([]byte(jsonStr), str)
	}
	return jsonStr, err
}

func (F *f2j) unmarshal(prefix []string, skep int) interface{} {
	data1 := ders{}
	data2 := make(derm)
	nextSkep := skep + 1
	isMap := false
	for _, r := range F.input {
		if len(r.keys) >= nextSkep {
			if strings.Join(r.keys[0:skep], ".") == strings.Join(prefix, ".") {
				isNum := false
				if regexp.MustCompile(`^\d+$`).FindString(r.keys[skep]) != "" {
					isNum = true
				}
				if len(r.keys) == nextSkep {
					if isNum {
						data1 = append(data1, r.value)
					} else {
						data2[r.keys[skep]] = r.value
					}
				} else {
					if isNum {
						isMap = true
					}
					data2[r.keys[skep]] = F.unmarshal(r.keys[0:nextSkep], nextSkep)
				}
			}
		}
	}
	if len(data1) > 0 {
		return data1
	}
	if len(data2) > 0 {
		if isMap {
			for _, v := range data2 {
				data1 = append(data1, v)
			}
			return data1
		}
		return data2
	}
	return nil
}
