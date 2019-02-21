package main

import (
	"fmt"
    "log"
    "os"
    "crypto/tls"
	"net/http"
    "github.com/hashicorp/go-rootcerts"
    "github.com/hashicorp/go-cleanhttp"
)

// This is just for testing go modules
func httpClient() (*http.Client, error) {
	tlsConfig := &tls.Config{}
	err := rootcerts.ConfigureTLS(tlsConfig, &rootcerts.Config{
		CAFile: os.Getenv("MYAPP_CAFILE"),
		CAPath: os.Getenv("MYAPP_CAPATH"),
	})
	if err != nil {
		return nil, err
	}
	c := cleanhttp.DefaultClient()
	t := cleanhttp.DefaultTransport()
	t.TLSClientConfig = tlsConfig
	c.Transport = t
	return c, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "go-sample: %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}