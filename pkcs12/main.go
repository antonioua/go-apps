package main

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"github.com/antonioua/go-tools/v2/pkcs12/certs"
	"io/ioutil"
	"log"
	"software.sslmate.com/src/go-pkcs12"
)

func main() {
	opts := certs.MakeDefaultOptions()

	caKp, err := certs.CreateCA(opts)
	if err != nil {
		log.Fatalf("Cannot create CA certificate. %v", err)
	}

	err = ioutil.WriteFile("./files/ca.pem", caKp.Certificate(), 0755)
	if err != nil {
		log.Fatal("error writing CA certificate file")
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
		log.Fatalf("Cannot create certificate. %v", err)
	}

	parsedCrt, err := certs.ParseCertificate(kp.Certificate())
	if err != nil {
		log.Fatalf("Cannot parse certificate. %v", err)
	}

	err = ioutil.WriteFile("./files/key.pem", kp.PrivateKey(), 0755)
	if err != nil {
		log.Fatal("Error writing PrivateKey file")
	}

	err = ioutil.WriteFile("./files/crt.pem", kp.Certificate(), 0755)
	if err != nil {
		log.Fatal("Error writing Certificate file")
	}

	parsedKey, err := certs.ParsePrivateKey(kp.PrivateKey())
	if err != nil {
		log.Fatalf("Cannot parse private key. %v", err)
	}

	fmt.Println("Encoding PFX Keystore...")
	keystorePFXBytes, err := pkcs12.Encode(rand.Reader, parsedKey, parsedCrt, []*x509.Certificate{parsedCACrt}, pkcs12.DefaultPassword)
	if err != nil {
		log.Fatalf("Cannot encode PFX Keystore. %v", err)
	}

	// see if pfxBytes valid
	_, _, _, err = pkcs12.DecodeChain(keystorePFXBytes, pkcs12.DefaultPassword)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing keystore.p12 file...")
	err = ioutil.WriteFile("./files/keystore.p12", keystorePFXBytes, 0755)
	if err != nil {
		log.Fatalf("Cannot write keystore.p12 file. %v", err)
	}

	fmt.Println("Encoding PFX Truststore...")
	truststorePFXBytes, err := pkcs12.EncodeTrustStore(rand.Reader, []*x509.Certificate{parsedCACrt}, pkcs12.DefaultPassword)
	if err != nil {
		log.Fatalf("Cannot encode PFX Truststore. %v", err)
	}

	fmt.Println("Writing truststore.p12 file...")
	err = ioutil.WriteFile("./files/truststore.p12", truststorePFXBytes, 0755)
	if err != nil {
		log.Fatalf("Cannot write truststore.p12 file. %v", err)
	}

}
