package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$kX.Xi000du6lpY31/5BSbuEktfs4gx7aqBIKVr/908iKiHajkITnm", "123456"))

}
