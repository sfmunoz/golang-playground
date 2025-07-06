//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Fri Oct 13 05:26:35 PM UTC 2023
//

// {{{ package

package http_json

// }}}
// {{{ imports

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// }}}
// {{{ globals

const ADDR = "127.0.0.1:3333"

// }}}
// ---- funcs (private) ----
// {{{ func reqPrint()

func reqPrint(r *http.Request) {
	fmt.Printf("==== %s %s %s ====\n", r.Method, r.URL.Path, r.Proto)
	for k, v := range r.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, "|"))
	}
	if r.ContentLength < 1 {
		fmt.Println("(... no body ...)")
		return
	}
	p := make([]byte, r.ContentLength)
	n, err := r.Body.Read(p)
	if err == io.EOF {
		fmt.Println("<-- EOF -->")
	} else if err != nil {
		fmt.Println("error getting body", err)
		return
	}
	fmt.Println("BODY:", n, string(p))
}

// }}}
// {{{ func getRoot()

func getRoot(w http.ResponseWriter, r *http.Request) {
	reqPrint(r)
	w.Header().Set("X-App", "myServer")
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		js := make(map[string]string)
		js["res"] = "ko"
		b, err := json.Marshal(js)
		if err != nil {
			io.WriteString(w, "{\"RES\":\"KO\"}\n")
		} else {
			// alt: io.WriteString(w, string(b) + "\n")
			b = append(b, '\n')
			w.Write(b)
		}
		return
	}
	io.WriteString(w, "This is the root of my website!\n")
}

// }}}
// {{{ func getHello()

func getHello(w http.ResponseWriter, r *http.Request) {
	reqPrint(r)
	io.WriteString(w, "hello world\n")
}

// }}}
// {{{ func httpServer()

func httpServer() {
	prefix := "httpServer()"
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fmt.Printf("%s: listening on '%s'...\n", prefix, ADDR)
	err := http.ListenAndServe(ADDR, nil)
	fmt.Printf("%s: error %s\n", prefix, err)
}

// }}}
// {{{ func httpClient()

func httpClient() {
	prefix := "httpClient()"
	time.Sleep(time.Second * 1)
	client := &http.Client{}
	url := "http://" + ADDR
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("%s: error creating request: %s\n", prefix, err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s: error execution request: %s\n", prefix, err)
		return
	}
	fmt.Printf("%s: %s\n", prefix, resp.Status)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s: error reading body: %s\n", prefix, err)
		return
	}
	fmt.Println("BODY:", string(body))
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	fmt.Println("==== Main() ====")
	defer fmt.Println("---- Main() ----")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServer()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpClient()
		wg.Done() // XXX: temporary to enforce stop
	}()
	wg.Wait()
}

// }}}
