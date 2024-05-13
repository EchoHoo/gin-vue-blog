package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	s := HashPwd("123456")
	fmt.Println(s)
	fmt.Println(CheckPwd(s, "123456"))
}
