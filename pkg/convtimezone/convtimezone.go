package convtimezone

import (
	"example/github.com/k-avy/gotimezoneapi/pkg/timejspasing"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	From := r.URL.Query().Get("from")
	To := r.URL.Query().Get("to")

	utc := time.Now().UTC()
	t := utc.String()
	fmt.Fprintf(w, "Time (UTC) : %s\n", t)
	ele := strings.Fields(t)
	e := strings.Split(ele[1], ":")
	sec := strings.Split(e[2], ".")[0]
	h, _ := strconv.Atoi(e[0])
	m, _ := strconv.Atoi(e[1])
	s, _ := strconv.Atoi(sec)

	off := timejspasing.GetOffset(From)
	if off == float32(int(off)) {
		h = (h + int(off)) % 24
	} else {
		m = m + 30
		h = (h + int(off) + int(m/60)) % 24
		m = m % 60
	}

	fmt.Fprintf(w, "Time= %d:%d:%d \n", h, m, s)

	tt := utc.String()
	el := strings.Fields(tt)
	em := strings.Split(el[1], ":")
	se := strings.Split(em[2], ".")[0]
	hh, _ := strconv.Atoi(em[0])
	mm, _ := strconv.Atoi(em[1])
	ss, _ := strconv.Atoi(se)

	off1 := timejspasing.GetOffset(To)
	if off1 == float32(int(off1)) {
		hh = (hh + int(off1)) % 24
	} else {
		mm = mm + 30
		hh = (hh + int(off1) + int(mm/60)) % 24
		mm = mm % 60
	}

	fmt.Fprintf(w, "Converted Time= %d:%d:%d \n", hh, mm, ss)

}
