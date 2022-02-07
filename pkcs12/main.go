package main

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"github.com/antonioua/go-tools/v2/pkcs12/certs"
	"io/ioutil"
	"log"
	"software.sslmate.com/src/go-pkcs12"
	_ "software.sslmate.com/src/go-pkcs12"
)

func main() {
	opts := certs.MakeDefaultOptions()

	caKp, err := certs.CreateCA(opts)
	if err != nil {
		log.Fatalf("Cannot create CA certificate. %v", err)
	}

	err = ioutil.WriteFile("./files/test-ca.pem", caKp.Certificate(), 0755)
	if err != nil {
		log.Fatal("error writing Certificate file")
	}

	parsedCACrt, err := certs.ParseCertificate(caKp.Certificate())
	if err != nil {
		log.Fatalf("Failed to parse pem to der CA certificate. %v", err)
	}

	opts.DnsNames = []string{
		"mydomain.com",
	}
	opts.Org = "Test App"

	kp, err := certs.CreateCertificate(caKp, opts)
	if err != nil {
		log.Fatalf("Cannot crate certificate. %v", err)
	}

	parsedCrt, err := certs.ParseCertificate(kp.Certificate())
	if err != nil {
		log.Fatalf("Cannot parse certificate. %v", err)
	}

	err = ioutil.WriteFile("./files/test-key.pem", kp.PrivateKey(), 0755)
	if err != nil {
		log.Fatal("error writing PrivateKey file")
	}

	parsedKey, err := certs.ParsePrivateKey(kp.PrivateKey())
	if err != nil {
		log.Fatalf("Cannot parse private key. %v", err)
	}

	fmt.Println("Encoding...")
	pfxBytes, err := pkcs12.Encode(rand.Reader, parsedKey, parsedCrt, []*x509.Certificate{parsedCACrt}, pkcs12.DefaultPassword)
	if err != nil {
		log.Fatal(err)
	}

	// see if pfxBytes valid
	_, _, _, err = pkcs12.DecodeChain(pfxBytes, pkcs12.DefaultPassword)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing keystore.p12 file...")
	err = ioutil.WriteFile("./files/keystore.p12", pfxBytes, 0755)
	if err != nil {
		log.Fatalf("Cannot write PKCS12 file archive. %v", err)
	}

	// truststorep.12
}
