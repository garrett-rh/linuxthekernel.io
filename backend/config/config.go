// parses down the config values
// environment variables are overwritten by cmd args
package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Tls  bool
	Port int
	Cert string
	Key  string
}

// New returns a config with reasonable defaults
func New() *Config {
	return &Config{
		Tls:  false,
		Port: 8080,
		Cert: "",
		Key:  "",
	}
}

// PopulateConfig If any flag is set, will attempt to populate
// If tls is set to false, cert and key flags will be ignored
func (c *Config) PopulateConfig() *Config {
	err := c.envVars()
	if err != nil {
		log.Fatalln("error parsing environment variables")
	}
	c.cmdArgs()

	return c
}

func (c *Config) cmdArgs() {
	// if set, key and cert should be set as well
	// otherwise they will default to /run/secrets/key and /run/secrets/cert respectfully
	tls := flag.Bool("tls", false, "Enable TLS")

	// arbitrary port to listen on
	port := flag.Int("port", 80, "Port to listen on")
	// these below flags will be ignored if tls is not set
	key := flag.String("key", "/run/secrets/key", "Path to private key")
	cert := flag.String("cert", "/run/secrets/cert", "Path to certificate")
	flag.Parse()

	if isFlagPassed("tls") {
		c.Tls = *tls
	}
	if isFlagPassed("port") {
		c.Port = *port
	}
	if isFlagPassed("key") {
		c.Key = *key
	}
	if isFlagPassed("cert") {
		c.Cert = *cert
	}
}

// verifies that we were given a specific flag
// used to verify that
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// looks for env vars
// env vars will be overwritten by config args
// looks for env vars of the format 'LTK_VARNAME'
func (c *Config) envVars() error {
	var err error
	// only accept a true value, naively assume all other values are false
	if os.Getenv("LTK_TLS") == "true" {
		c.Tls = true
	}

	if os.Getenv("LTK_PORT") != "" {
		c.Port, err = strconv.Atoi(os.Getenv("LTK_PORT"))
		if err != nil {
			return err
		}
	}

	if os.Getenv("LTK_KEY") != "" {
		c.Key = os.Getenv("LTK_KEY")
	}

	if os.Getenv("LTK_CERT") != "" {
		c.Cert = os.Getenv("LTK_CERT")
	}
	return nil
}
