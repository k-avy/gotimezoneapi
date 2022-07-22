package convzone

import (
	"example/github.com/k-avy/gotimezoneapi/pkg/timejspasing"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ConvertTime(w http.ResponseWriter, r *http.Request) {

	From := r.URL.Query().Get("from")
	To := r.URL.Query().Get("to")
	ele := strings.Split(From, "_")
	t := ele[0]
	Fr := ele[1]
	e := strings.Split(t, ":")

	h, _ := strconv.Atoi(e[0])
	m, _ := strconv.Atoi(e[1])
	s, _ := strconv.Atoi(e[2])

	off := timejspasing.GetOffset(Fr)
	if off == float32(int(off)) {
		h = ((h - int(off)) + 24) % 24
	} else {
		m = m - 30
		if m < 0 {
			h = ((h - int(off) - 1) + 24) % 24
			m = m + 60
		} else {
			h = ((h - int(off)) + 24) % 24

		}
	}

	of := timejspasing.GetOffset(To)
	if of == float32(int(of)) {
		h = (h + int(of)) % 24
	} else {
		m = m + 30
		h = (h + int(of) + int(m/60)) % 24
		m = m % 60
	}

	fmt.Fprintf(w, "Converted Time= %d:%d:%d \n", h, m, s)

}
