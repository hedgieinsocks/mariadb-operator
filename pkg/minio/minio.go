package minio

import (
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioOpts struct {
	TLS        bool
	CACertPath string
	Region     string
}

type MinioOpt func(m *MinioOpts)

func WithTLS(tls bool) MinioOpt {
	return func(m *MinioOpts) {
		m.TLS = tls
	}
}

func WithCACertPath(caCertPath string) MinioOpt {
	return func(m *MinioOpts) {
		m.CACertPath = caCertPath
	}
}

func WithRegion(region string) MinioOpt {
	return func(m *MinioOpts) {
		m.Region = region
	}
}

func NewMinioClient(endpoint string, mOpts ...MinioOpt) (*minio.Client, error) {
	opts := MinioOpts{}
	for _, setOpt := range mOpts {
		setOpt(&opts)
	}

	minioOpts, err := getMinioOptions(opts)
	if err != nil {
		return nil, fmt.Errorf("error creating Minio client options: %v", err)
	}
	client, err := minio.New(endpoint, minioOpts)
	if err != nil {
		return nil, fmt.Errorf("error creating Minio client: %v", err)
	}
	return client, nil
}

func getMinioOptions(opts MinioOpts) (*minio.Options, error) {
	transport, err := getTransport(&opts)
	if err != nil {
		return nil, fmt.Errorf("error getting transport: %v", err)
	}

	// Use a chained credentials provider to support multiple sources:
	// 1. Environment variables (set by custom resource)
	// 2. IAM role (for EC2 Meta Data, EKS service accounts when environment variables are not set)
	chainedCreds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvAWS{},
			&credentials.IAM{},
		},
	)

	minioOpts := &minio.Options{
		Creds:     chainedCreds,
		Region:    opts.Region,
		Secure:    opts.TLS,
		Transport: transport,
	}
	return minioOpts, nil
}

func getTransport(opts *MinioOpts) (*http.Transport, error) {
	transport, err := minio.DefaultTransport(opts.TLS)
	if err != nil {
		return nil, err
	}
	if !opts.TLS {
		return transport, nil
	}

	if transport.TLSClientConfig.RootCAs == nil {
		pool, err := x509.SystemCertPool()
		if err != nil {
			transport.TLSClientConfig.RootCAs = x509.NewCertPool()
		} else {
			transport.TLSClientConfig.RootCAs = pool
		}
	}

	if opts.CACertPath != "" {
		caBytes, err := os.ReadFile(opts.CACertPath)
		if err != nil {
			return nil, fmt.Errorf("error reading CA cert: %v", err)
		}
		if ok := transport.TLSClientConfig.RootCAs.AppendCertsFromPEM(caBytes); !ok {
			return nil, errors.New("unable to add CA cert to pool")
		}
	}

	return transport, nil
}
