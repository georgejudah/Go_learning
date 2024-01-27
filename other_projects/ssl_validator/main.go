package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

const (
	serverAddress = "google.co.in"
	port          = "443"
)

func main() {
	err := checkSSLExpiry(serverAddress + ":" + port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Certificate is valid.")
}

func checkSSLExpiry(address string) error {
	conn, err := tls.Dial("tcp", address, nil)
	if err != nil {
		return fmt.Errorf("Failed to establish TLS connection: %v", err)
	}
	defer conn.Close()

	err = conn.VerifyHostname(serverAddress)
	if err != nil {
		return fmt.Errorf("Hostname doesn't match with certificate: %v", err)
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))

	if time.Now().After(expiry) {
		return fmt.Errorf("Certificate has expired")
	}

	return nil
}
