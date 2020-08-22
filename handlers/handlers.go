package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Matias-Barrios/echoMe/models"
)

func PlainTextEcho(w http.ResponseWriter, r *http.Request) {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	writers = append(writers, w)
	dest := io.MultiWriter(writers...)
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(dest, err.Error())
		return
	}
	fmt.Fprintf(dest, "Method -> \n")
	fmt.Fprintf(dest, "\t%s\n\n", r.Method)
	fmt.Fprintf(dest, "Protocol -> \n")
	fmt.Fprintf(dest, "\t%s\n\n", r.Proto)
	fmt.Fprintf(dest, "Headers -> \n")
	for k, v := range r.Header {
		fmt.Fprintf(dest, "\t%s: %s\n", k, v)
	}
	fmt.Fprintf(dest, "\nContent length (in Bytes) -> \n")
	fmt.Fprintf(dest, "\t%d\n\n", r.ContentLength)
	fmt.Fprintf(dest, "\nRemote address -> \n")
	fmt.Fprintf(dest, "\t%s\n", r.RemoteAddr)
	fmt.Fprintf(dest, "\nPayload\n")
	fmt.Fprintf(dest, "####################\n")
	fmt.Fprintf(dest, "%s\n", payload)
	fmt.Fprintf(dest, "####################\n")
}

func JsonEcho(w http.ResponseWriter, r *http.Request) {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	writers = append(writers, w)
	dest := io.MultiWriter(writers...)
	if r.URL.Path != "/json" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(dest, err.Error())
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
		fmt.Fprintf(dest, err.Error())
		return
	}
	fmt.Fprintf(dest, string(jsonbody)+"\n")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Path not found. Only valid paths are '/' and '/json'\n")
	}
}
