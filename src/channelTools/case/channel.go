package _case

import (
	"demogo/src/tools"
	"fmt"
)

var ()

func GetCon() {

	path := "D:/demogo/src/resources/config.yaml"
	res := tools.ConfigDbConnUrl(path)
	fmt.Println(res)

}
