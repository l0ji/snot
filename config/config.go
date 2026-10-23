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
	c.ListenAddr = os.Getenv("LISTEN_ADDR")
	if c.ListenAddr == "" {
		c.ListenAddr = "127.0.0.1:8080"
	}
	c.BaseURL = os.Getenv("BASE_URL")
	if c.BaseURL == "" {
		c.BaseURL = "http://" + c.ListenAddr
	}
	c.RenderedPagesFolder = os.Getenv("RENDERED_PAGES_FOLDER")
	if c.RenderedPagesFolder == "" {
		c.RenderedPagesFolder = "./cache"
	}
	c.CustomCss = os.Getenv("CUSTOM_CSS")
	return c
}
