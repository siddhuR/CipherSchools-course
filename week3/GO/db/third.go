package db

import "log"

var key int

func Getkey() {
	key = 9
}

type val struct {
	i int
}

func (v *val) Getval() int {
	v.i = 10
	return v.i
}

func get() {
	vv := val()
	vv.Getval()
	log.Println(vv.i)
	ww := &val()
	ww.Getval()
	log.Println(ww.i)
}

func Val() {
	panic("unimplemented")
}
