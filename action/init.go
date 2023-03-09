package action

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
)

type Config struct {
	fp               string `toml:"-"`
	LoginCompanyCode string `toml:"login_company_code"` // AKASHI企業ID
	Token            string `toml:"token"`              // トークン
}

func (c *Config) Load(profile string) (err error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return
	}
	dir = filepath.Join(dir, "ak4")

	var fp string
	if profile == "" {
		fp = filepath.Join(dir, "config.toml")
	} else {
		fp = filepath.Join(dir, fmt.Sprintf("config-%s.toml", profile))
	}
	os.MkdirAll(filepath.Dir(fp), 0700)
	c.fp = fp

	b, err := os.ReadFile(fp)
	if err != nil {
		return
	}

	err = toml.Unmarshal(b, c)

	return
}

func (c Config) Write(loginCompanyCode, token string) (err error) {
	if loginCompanyCode == "" && token == "" {
		return errors.New("login_company_code or token must be set")
	}
	if loginCompanyCode != "" {
		c.LoginCompanyCode = loginCompanyCode
	}
	if token != "" {
		c.Token = token
	}

	b, err := toml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(c.fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func ActInit(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		fmt.Println(err)
		fmt.Println("create new profile")
	}

	if err = cfg.Write(ctx.String("login_company_code"), ctx.String("token")); err != nil {
		return
	}

	return
}
