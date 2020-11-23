package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"github.com/bitrise-tools/go-steputils/tools"
)

const (
	baseURL = "https://api.qrserver.com/v1/create-qr-code/"
)

// Config ...
type Config struct {
	PublicInstallPageURL string `env:"public_install_page_url,required"`
	QRCodeSize           string `env:"qr_code_size,required"`
}

func main() {
	var cfg Config
	log.Errorf("cathy fucked this up", "4")
	if err := stepconf.Parse(&cfg); err != nil {
		log.Errorf("Error: %s\n", err)
		os.Exit(1)
	}

	stepconf.Print(cfg)
	fmt.Println()

	QRCodeURL, err := generateQRCode(cfg.PublicInstallPageURL, cfg.QRCodeSize)
	if err != nil {
		failf("Failed to generate QR Code for %s, error: %s", cfg.PublicInstallPageURL, err)
	}

	log.Printf("$BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL=(%s)", QRCodeURL)

	if err := tools.ExportEnvironmentWithEnvman("BITRISE_PUBLIC_INSTALL_PAGE_QR_CODE_IMAGE_URL", QRCodeURL); err != nil {
		failf("Failed to generate output")
	}
}

func generateQRCode(installPageURL string, qrCodeSize string) (string, error) {
	requestURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	quearryValues := requestURL.Query()
	quearryValues.Add("size", qrCodeSize)
	quearryValues.Add("data", installPageURL)
	requestURL.RawQuery = quearryValues.Encode()

	return requestURL.String(), nil
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}
