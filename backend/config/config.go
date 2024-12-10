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

// PopulateConfig Returns a config that is bootstrapped
// Defaults to tls off, port 80
// If any flag is set, will attempt to populate
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
	tls := flag.Bool("tls", false, "Enable TLS")
	port := flag.Int("port", 80, "Port to listen on")
	key := flag.String("key", "/run/secrets/key", "Path to private key")
	cert := flag.String("cert", "/run/secrets/cert", "Path to certificate")
	flag.Parse()

	c.Tls = *tls
	c.Cert = *cert
	c.Key = *key
	c.Port = *port
}

// looks for env vars
// env vars will be overwritten by config args
// looks for env vars of the format 'LTK_VARNAME'
func (c *Config) envVars() error {
	var err error
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
