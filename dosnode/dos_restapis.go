package dosnode

import (
	"net/http"
	"strconv"
)

func (d *DosNode) Status(w http.ResponseWriter, r *http.Request) {
	html := "Dos Client is working "
	w.Write([]byte(html))
}

func (d *DosNode) Balance(w http.ResponseWriter, r *http.Request) {
	html := "Balance "
	result := d.chain.Balance()
	html = html + result.String()
	w.Write([]byte(html))
}

func (d *DosNode) Groups(w http.ResponseWriter, r *http.Request) {
	html := "Group Number "
	result := d.dkg.GetGroupNumber()
	html = html + strconv.Itoa(result)
	w.Write([]byte(html))
}
