package form2json

import (
	"log"
	"testing"
)

type dss struct {
	BoardID  uint64 `json:"board_id"`
	Category string `json:"category"`
}
type dd interface{}
type ds []interface{}
type dm map[string]interface{}

func TestUnmarshal(t *testing.T) {
	s := "attachments%5B0%5D%5Bembeded%5D=0&attachments%5B0%5D%5Bfid%5D=&attachments%5B0%5D%5Bsize%5D=0&attachments%5B0%5D%5Burl%5D=http%3A%2F%2Foncloud.quanshi.com%3A80%2Fucfserver%2Fhddown%3Ffid%3DMTk4MDI5My84L-aOqOiNkCAyLnBuZ15eXnRhbmdoZGZzXl5eYjA0ZDIxMjJiNWM5MWIzOTNlNDM1NzBmNTE0NThmOGVeXl50YW5naGRmc15eXjc4NjQxMA%24%26u%3D1980293&attachments%5B1%5D%5Bembeded%5D=1&attachments%5B1%5D%5Bfid%5D=11&attachments%5B1%5D%5Bsize%5D=false&attachments%5B1%5D%5Burl%5D=http%3A%2F%2F111oncloud.quanshi.com%3A80%2Fucfserver%2Fhddown%3Ffid%3DMTk4MDI5My84L-aOqOiNkCAyLnBuZ15eXnRhbmdoZGZzXl5eYjA0ZDIxMjJiNWM5MWIzOTNlNDM1NzBmNTE0NThmOGVeXl50YW5naGRmc15eXjc4NjQxMA%24%26u%3D1980293&board_id=6786&category=bbs&comment_enabled=1&content=%3Cp%3E%3Cspan+style%3D%22font-size%3A+14px%3B%22%3E%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A%3C%2Fspan%3E%3Cbr%2F%3E%3C%2Fp%3E&description=%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A&object_type=bbs&publish_scope%5Bgroup_ids%5D=&publish_scope%5Buser_ids%5D%5B0%5D=1382766&publish_scope%5Buser_ids%5D%5B1%5D=false&title=%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A"
	//	s := "lists%5B0%5D%5Bid%5D=111&lists%5B0%5D%5Bname%5D=fffff&lists%5B1%5D%5Bid%5D=222&lists%5B1%5D%5Bname%5D=eee&lists%5B3%5D%5Bid%5D=333&lists%5B3%5D%5Blist%5D%5Bid%5D=11&lists%5B3%5D%5Blist%5D%5Blistsss%5D%5B0%5D=1&lists%5B3%5D%5Blist%5D%5Blistsss%5D%5B1%5D=2&lists%5B3%5D%5Blist%5D%5Blistsss%5D%5B2%5D=3&lists%5B3%5D%5Blist%5D%5Blistsss%5D%5B3%5D=4&lists%5B3%5D%5Blist%5D%5Blistsss%5D%5B4%5D=5&s=sss&ss%5B0%5D=12&ss%5B1%5D=33&ss%5B2%5D=44&ss%5B3%5D=55"
	//s = "a=010&b=0.10&c=0.0001&d=1000.11&e=0&k=111&f=1&g=9&h=9.00&i=true&j=false&p=909090.55"
	var ssss dss
	ss, _ := Unmarshal(s, &ssss)
	v := 0999.9
	log.Println(v)
	log.Println(ss)
	//log.Println(ssss)

}
