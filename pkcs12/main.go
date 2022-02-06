package main

import (
	"crypto/x509"
	"io/ioutil"
	"log"
	"software.sslmate.com/src/go-pkcs12"
	_ "software.sslmate.com/src/go-pkcs12"
	"strings"
)

func main() {
	key, err := ioutil.ReadFile("./certs/localhost_key.pem")
	if err != nil {
		log.Fatalf("Cannot read key. Error: %v", err)
	}
	//fmt.Printf("keyFile contents: %s", key)

	cert, err := ioutil.ReadFile("./certs/localhost_cert.pem")
	if err != nil {
		log.Fatalf("Cannot read cert. Error: %v", err)
	}
	parsedCert, err := x509.ParseDERCRL(cert)
	if err != nil {
		log.Fatalf("Cannot parse cert. Error: %v", err)
	}

	ca, err := ioutil.ReadFile("./certs/ca_cert.pem")
	if err != nil {
		log.Fatalf("Cannot read ca. Error: %v", err)
	}
	parsedCA, err := x509.ParseCertificate(ca)
	if err != nil {
		log.Fatalf("Cannot parse ca. Error: %v", err)
	}

	parsedCAs := make([]*x509.Certificate, 1)
	parsedCAs = append(parsedCAs, parsedCA)

	encodedBytes, err := pkcs12.Encode(strings.NewReader("blablabla"), key, parsedCert, parsedCAs, "password")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./certs/archive.p12", encodedBytes, 0755)
	if err != nil {
		return
	}
}
