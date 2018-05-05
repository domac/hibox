package main

import (
	r "./ref"
)

func main() {
	u := &r.User{}
	u.Name = "dom"
	u.Age = 18
	r.FiltName(u, "can you give me some beer?")
	r.FiltNameWithReuseOffset(u, "can you give me some beer?")
	r.FiltNameWithCache(u, "can you give me some beer?")
}
