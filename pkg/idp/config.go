package web3idp

import (
	"fmt"
	"io"
	"os"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/log"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/mysql"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/web3oidc/hydra"

	"github.com/kelseyhightower/envconfig"
	"github.com/pborman/getopt/v2"
	"gopkg.in/yaml.v2"
)

// MQTTConfig : structure to hold local to service configuration
type LocalConfig struct {
	URL      string `yaml:"url" envconfig:"WEB3_IDP_BIND_ADDRESS"`
	LogLevel uint   `yaml:"loglevel" envconfig:"WEB3_IDP_LOGLEVEL"`
}

func (x *LocalConfig) Init() {
	x.LogLevel = 1
	x.URL = "0.0.0.0:4000"
}

// Config : structure to hold configuration
type Config struct {
	MySQL    mysql.Config `yaml:"mysql"`
	Hydra    hydra.Config `yaml:"hydra"`
	Settings LocalConfig  `yaml:"settings"`
}

func (x *Config) Init() {
	x.MySQL.Init()
}

func (cfg *Config) defConfig() {
	cfg.Init()
}

func (cfg *Config) readOpts() {
	helpFlag := false
	getopt.Flag(&helpFlag, 'h', "display help")
	getopt.Flag(&cfg.Settings.LogLevel, 'l', "be verbose")

	getopt.Parse()
	if helpFlag {
		getopt.Usage()
		os.Exit(0)
	}
}

func (cfg *Config) processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func (cfg *Config) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (cfg *Config) readFile(path string) {
	if !cfg.fileExists(path) {
		return
	}
	f, err := os.Open(path)
	if err != nil {
		cfg.processError(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		if err != io.EOF {
			cfg.processError(err)
		}
	}
}

func (cfg *Config) readEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		cfg.processError(err)
	}
}

func (cfg *Config) PrettyPrint() {
	d, _ := yaml.Marshal(cfg)
	log.Logf(1, "--- Config ---\n%s\n\n", string(d))
}

// GetConfig : get config file
func GetConfig(path string, enableFlags bool) Config {
	var cfg Config

	cfg.defConfig()

	cfg.readFile(path)
	cfg.readEnv()

	if enableFlags {
		cfg.readOpts()
	}
	log.SetLogLevel(cfg.Settings.LogLevel)

	cfg.PrettyPrint()
	return cfg
}
