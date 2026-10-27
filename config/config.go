// Package config provides methods, functions, and structs for the server configuration
package config

import (
	"errors"
	"net/netip"
	"net/url"
	"os"
)

// Config handles config data as a single struct
type Config struct {
	ListenAddr          string
	BaseURL             string
	InstanceName        string
	RenderedPagesFolder string
	CustomCss           string
}

// IsValid validates whether a generated Config is valid or not
func (c *Config) IsValid() (err error) {
	_, err = netip.ParseAddrPort(c.ListenAddr)
	if err != nil {
		return err
	}

	_, err = url.Parse(c.BaseURL)
	if err != nil {
		return err
	}

	info, err := os.Stat(c.RenderedPagesFolder)
	if err != nil {
		if !os.IsExist(err) {
			err := os.Mkdir(c.RenderedPagesFolder, 0750)
			if err != nil {
				return err
			}
			info, err = os.Stat(c.RenderedPagesFolder)
			if err != nil && !os.IsExist(err) {
				return err
			}
		} else {
			return err
		}
	}
	if !info.IsDir() {
		return errors.New("rendered pages folder is not a directory")
	}

	info, err = os.Stat(c.CustomCss)
	if err != nil {
		if os.IsExist(err) {
			if !info.IsDir() {
				return err
			}
		}
	}

	return nil
}

// New creates a Config object based on some hard-coded defaults, otherwise
// it should use your configured environment variables
func New() (c Config) {
	var isSet bool

	c.ListenAddr, isSet = os.LookupEnv("LISTEN_ADDR")
	if !isSet {
		c.ListenAddr = "127.0.0.1:8080"
	}
	c.BaseURL, isSet = os.LookupEnv("BASE_URL")
	if !isSet {
		c.BaseURL = "http://" + c.ListenAddr
	}
	c.InstanceName, isSet = os.LookupEnv("INSTANCE_NAME")
	if !isSet {
		c.InstanceName = "snot"
	}
	c.RenderedPagesFolder, isSet = os.LookupEnv("RENDERED_PAGES_FOLDER")
	if !isSet {
		c.RenderedPagesFolder = "./cache"
	}
	c.CustomCss, isSet = os.LookupEnv("CUSTOM_CSS")
	if !isSet {
		c.CustomCss = ""
	}

	return c
}
