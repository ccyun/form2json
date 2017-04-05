package form2json

import (
	"log"
	"testing"
)

type dss struct {
	BoardID  uint64 `json:"board_id"`
	Category string `json:"category"`
}

func TestUnmarshal(t *testing.T) {
	s := "board_id=6786&attachments%5B0%5D%5Burl%5D=http%3A%2F%2Foncloud.quanshi.com%3A80%2Fucfserver%2Fhddown%3Ffid%3DMTk4MDI5My84L-aOqOiNkCAyLnBuZ15eXnRhbmdoZGZzXl5eYjA0ZDIxMjJiNWM5MWIzOTNlNDM1NzBmNTE0NThmOGVeXl50YW5naGRmc15eXjc4NjQxMA%24%26u%3D1980293&attachments%5B0%5D%5Bfid%5D=&attachments%5B0%5D%5Bsize%5D=0&attachments%5B0%5D%5Bembeded%5D=0&content=%3Cp%3E%3Cspan%20style%3D%22font-size%3A%2014px%3B%22%3E%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A%3C%2Fspan%3E%3Cbr%2F%3E%3C%2Fp%3E&title=%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A&object_type=bbs&publish_scope%5Bgroup_ids%5D=&publish_scope%5Buser_ids%5D%5B0%5D=1382766&publish_scope%5Buser_ids%5D%5B1%5D=1980293&description=%E6%B5%8B%E8%AF%95%E5%85%AC%E5%91%8A&comment_enabled=1&category=bbs"
	//s = "a=010&b=0.10&c=0.0001&d=1000.11&e=0&k=111&f=1&g=9&h=9.00&i=true&j=false&p=909090.55"
	var ssss dss
	ss, _ := Unmarshal(s, &ssss)

	log.Println(ss)
	log.Println(ssss)
}
