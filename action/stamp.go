package action

import (
	"context"
	"fmt"

	"github.com/hapoon/kiku"
	"github.com/urfave/cli/v2"
)

func ActStamp(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}
	param := kiku.PostStampParam{
		LoginCompanyCode: cfg.LoginCompanyCode,
		Token:            cfg.Token,
	}
	res, err := kiku.PostStamp(context.TODO(), param)
	if err == nil {
		fmt.Println("打刻種別:", res.Type)
		fmt.Println("打刻日時:", res.StampedAt)
	}
	return
}
