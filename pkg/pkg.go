package pkg

import (
	"fmt"
	"strconv"
)

func I2AA(i int)string{
	s := "0" + strconv.Itoa(i)
	return s[len(s)-2:]
}

func Sec2HHMMSS(t int)string{
	mh := 0
	for t < 0{
		t += 3600
		mh++
	}
	h := int(t)/3600-mh
	m := (int(t)-h*3600)/60
	s := int(t)%60
	return I2AA(h)+":"+I2AA(m)+":"+I2AA(s)
}

func HHMMSS2Sec(str string)int{
	var h, m, s int
	fmt.Sscanf(str, "%d:%d:%d", &h, &m, &s)
	return h*3600+m*60+s
}
