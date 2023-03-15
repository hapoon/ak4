package action

import (
	"context"
	"fmt"

	"github.com/hapoon/kiku"
	"github.com/urfave/cli/v2"
)

func ActToken(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}

	param := kiku.PostTokenReissueParam{
		LoginCompanyCode: cfg.LoginCompanyCode,
		Token:            cfg.Token,
	}
	res, err := kiku.PostTokenReissue(context.TODO(), param)
	if err != nil {
		return
	}

	cfg.Token = res.Token
	if err = cfg.Write("", res.Token); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("トークンを更新しました")
	fmt.Println("有効期限:", res.ExpiredAt)
	return
}
