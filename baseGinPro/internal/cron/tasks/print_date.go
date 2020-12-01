package tasks

import (
	"baseGinPro/pkg/define"
	"fmt"
	"time"
)

func PrintDate(){
	fmt.Printf("%s\n",time.Now().Format(define.YMD))
}
