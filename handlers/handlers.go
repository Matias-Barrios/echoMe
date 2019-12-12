package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Matias-Barrios/echoMe/models"
	"io/ioutil"
	"net/http"
)

func PlainTextEcho(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Method -> \n")
	fmt.Fprintf(w, "\t%s\n\n", r.Method)
	fmt.Fprintf(w, "Protocol -> \n")
	fmt.Fprintf(w, "\t%s\n\n", r.Proto)
	fmt.Fprintf(w, "Headers -> \n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t%s: %s\n", k, v)
	}
	fmt.Fprintf(w, "\nContent length (in Bytes) -> \n")
	fmt.Fprintf(w, "\t%d\n\n", r.ContentLength)
	fmt.Fprintf(w, "\nRemote address -> \n")
	fmt.Fprintf(w, "\t%s\n", r.RemoteAddr)
	fmt.Fprintf(w, "\nPayload\n")
	fmt.Fprintf(w, "####################\n")
	fmt.Fprintf(w, "%s\n", payload)
	fmt.Fprintf(w, "####################\n")
}

func JsonEcho(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/json" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	var resp models.Response = models.Response{
		Method:        r.Method,
		Payload:       string(payload),
		RemoteAddress: r.RemoteAddr,
		Protocol:      r.Proto,
		ContentLength: r.ContentLength,
	}
	resp.Headers = make(map[string][]string)
	for k, v := range r.Header {
		resp.Headers[k] = v
	}
	jsonbody, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(jsonbody)+"\n")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Path not found. Only valid paths are '/' and '/json'\n")
	}
}
