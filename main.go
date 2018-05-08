package main

import (
	"encoding/pem"
	"time"
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/elliptic"
	"crypto"
	"os"
	"math/big"
)

var (
	ReqCA = &x509.Certificate{
		IsCA: true,
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country: []string{"US",},
			Organization: []string{"GOCA",},
			CommonName: "Test CA with GO",
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(24 * 365 * time.Hour),
		MaxPathLen: 0,
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCRLSign | x509.KeyUsageCertSign,
	}

	Req = &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Country: []string{"US",},
			Organization: []string{"GOCA",},
			CommonName: "blih.toto.com",
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(24 * time.Hour),
	}
)

func encodePrivateKey(key crypto.PrivateKey) ([]byte, error) {
	raw, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{
		Type: "PRIVATE KEY",
		Bytes: raw,
	}

	payload := pem.EncodeToMemory(block)

	return payload, nil
}

func encodeCertificate(cert *x509.Certificate) ([]byte) {
	block := &pem.Block{
		Type: "CERTIFICATE",
		Bytes: cert.Raw,
	}

	return pem.EncodeToMemory(block)
}

func main() {
	caKey, caCert, err := createCert(ReqCA, ReqCA, nil)
	if err != nil {
		panic(err)
	}

	key, cert, err := createCert(Req, caCert, caKey)
	if err != nil {
		panic(err)
	}

	payload, err := encodePrivateKey(key)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(payload)

	payload = encodeCertificate(cert)

	os.Stdout.Write(payload)

	payload, err = encodePrivateKey(key)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(payload)

	payload = encodeCertificate(cert)

	os.Stdout.Write(payload)
}

func createCert(template *x509.Certificate, req *x509.Certificate, priv crypto.PrivateKey) (crypto.PrivateKey, *x509.Certificate, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	if priv == nil {
		priv = key
	}

	raw, err := x509.CreateCertificate(rand.Reader, template, req, key.Public(), priv)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(raw)
	if err != nil {
		return nil, nil, err
	}

	return key, cert, nil
}
