package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	a := false
	fmt.Printf("Value %v  and type %s \n", a, reflect.TypeOf(a))
	b := strconv.FormatBool(a)
	fmt.Printf("Value %v  and type %s \n", b, reflect.TypeOf(b))
	c := 10.5
	fmt.Printf("Value %v  and type %s \n", c, reflect.TypeOf(c))
	d := int(c)
	fmt.Printf("Value %v  and type %s \n", d, reflect.TypeOf(d))
	e := float32(c)
	fmt.Printf("Value %v  and type %s \n", e, reflect.TypeOf(e))
	f := complex(e, e)
	fmt.Printf("Value %v  and type %s \n", f, reflect.TypeOf(f))
	g := rune('a')
	fmt.Printf("Value %v  and type %s \n", g, reflect.TypeOf(g))
	h := byte('a')
	fmt.Printf("Value %v  and type %s \n", h, reflect.TypeOf(h))
	i := int8(10)
	fmt.Printf("Value %v  and type %s \n", i, reflect.TypeOf(i))
	j := int16(10)
	fmt.Printf("Value %v  and type %s \n", j, reflect.TypeOf(j))
	k := int32(10)
	fmt.Printf("Value %v  and type %s \n", k, reflect.TypeOf(k))
	l := int64(10)
	fmt.Printf("Value %v  and type %s \n", l, reflect.TypeOf(l))
	m := uint(10)
	fmt.Printf("Value %v  and type %s \n", m, reflect.TypeOf(m))
	n := uint8(10)
	fmt.Printf("Value %v  and type %s \n", n, reflect.TypeOf(n))
	o := uint16(10)
	fmt.Printf("Value %v  and type %s \n", o, reflect.TypeOf(o))
	p := uint32(10)
	fmt.Printf("Value %v  and type %s \n", p, reflect.TypeOf(p))
	q := uint64(10)
	fmt.Printf("Value %v  and type %s \n", q, reflect.TypeOf(q))
	r := float64(10)
	fmt.Printf("Value %v  and type %s \n", r, reflect.TypeOf(r))
	s := complex128(10)
	fmt.Printf("Value %v  and type %s \n", s, reflect.TypeOf(s))
	t := string("Hello")
	fmt.Printf("Value %v  and type %s \n", t, reflect.TypeOf(t))
	u := uintptr(10)
	fmt.Printf("Value %v  and type %s \n", u, reflect.TypeOf(u))
	v := interface{}(10)
	fmt.Printf("Value %v  and type %s \n", v, reflect.TypeOf(v))
	w := interface{}("Hello")
	fmt.Printf("Value %v  and type %s \n", w, reflect.TypeOf(w))
	x := interface{}(10.5)
	fmt.Printf("Value %v  and type %s \n", x, reflect.TypeOf(x))
	y := interface{}(false)
	fmt.Printf("Value %v  and type %s \n", y, reflect.TypeOf(y))
	z := interface{}(complex(10, 10))
	fmt.Printf("Value %v  and type %s \n", z, reflect.TypeOf(z))
	aa := interface{}([]int{1, 2, 3})
	fmt.Printf("Value %v  and type %s \n", aa, reflect.TypeOf(aa))
	bb := interface{}(map[string]int{"a": 1, "b": 2})
	fmt.Printf("Value %v  and type %s \n", bb, reflect.TypeOf(bb))
	cc := strconv.FormatInt(10, 2)
	fmt.Printf("Value %v  and type %s \n", cc, reflect.TypeOf(cc))
	dd := strconv.FormatFloat(10.5, 'f', -1, 64)
	fmt.Printf("Value %v  and type %s \n", dd, reflect.TypeOf(dd))
	ee := strconv.FormatBool(false)
	fmt.Printf("Value %v  and type %s \n", ee, reflect.TypeOf(ee))
	ff := strconv.FormatUint(10, 2)
	fmt.Printf("Value %v  and type %s \n", ff, reflect.TypeOf(ff))
	gg, _ := strconv.ParseBool("false")
	fmt.Printf("Value %v  and type %s \n", gg, reflect.TypeOf(gg))
	hh, _ := strconv.ParseInt("10", 10, 64)
	fmt.Printf("Value %v  and type %s \n", hh, reflect.TypeOf(hh))
	ii, _ := strconv.ParseFloat("10.5", 64)
	fmt.Printf("Value %v  and type %s \n", ii, reflect.TypeOf(ii))
	jj, _ := strconv.ParseBool("false")
	fmt.Printf("Value %v  and type %s \n", jj, reflect.TypeOf(jj))
	kk, _ := strconv.ParseUint("10", 10, 64)
	fmt.Printf("Value %v  and type %s \n", kk, reflect.TypeOf(kk))
	ll := strconv.Quote("Hello")
	fmt.Printf("Value %v  and type %s \n", ll, reflect.TypeOf(ll))
	mm := strconv.QuoteToASCII("Hello")
	fmt.Printf("Value %v  and type %s \n", mm, reflect.TypeOf(mm))
	nn := strconv.QuoteRune('a')
	fmt.Printf("Value %v  and type %s \n", nn, reflect.TypeOf(nn))
	oo, _ := strconv.Unquote("\"Hello\"")
	fmt.Printf("Value %v  and type %s \n", oo, reflect.TypeOf(oo))
	pp, _, _, _ := strconv.UnquoteChar("'a'", 'a')
	fmt.Printf("Value %v  and type %s \n", pp, reflect.TypeOf(pp))
	qq, _, _, _ := strconv.UnquoteChar("'a'", 'b')
	fmt.Printf("Value %v  and type %s \n", qq, reflect.TypeOf(qq))
	rr := strconv.Itoa(1000)
	fmt.Printf("Value %v  and type %s \n", rr, reflect.TypeOf(rr))
	ss := strconv.AppendInt([]byte{}, 10, 10)
	fmt.Printf("Value %v  and type %s \n", ss, reflect.TypeOf(ss))
	tt := strconv.AppendQuote([]byte{}, "Hello")
	fmt.Printf("Value %v  and type %s \n", tt, reflect.TypeOf(tt))
	uu := strconv.AppendQuoteRune([]byte{}, 'a')
	fmt.Printf("Value %v  and type %s \n", uu, reflect.TypeOf(uu))
	vv := strconv.AppendBool([]byte{}, false)
	fmt.Printf("Value %v  and type %s \n", vv, reflect.TypeOf(vv))
	ww := strconv.AppendUint([]byte{}, 10, 10)
	fmt.Printf("Value %v  and type %s \n", ww, reflect.TypeOf(ww))
}
