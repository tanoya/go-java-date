package date

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	d := New()
	var formater string = "yyyy-MM-dd HH:mm:ss"
	ret := d.Java().String(formater)
	fmt.Print(ret)
}

func TestString01(t *testing.T) {
	d := New()
	var formater string = "yyyy-MM-dd HH:mm:ssSSS"
	ret := d.Java().String(formater)
	fmt.Print(ret)
}
