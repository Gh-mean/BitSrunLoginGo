package controllers

import (
	"github.com/Mmx233/BitSrunLoginGo/global"
	BitSrun "github.com/Mmx233/BitSrunLoginGo/v1"
	"github.com/Mmx233/BitSrunLoginGo/v1/transfer"
)

func Login(output bool, skipCheck bool) error {
	return BitSrun.Login(&srunTransfer.Login{
		Demo:     global.Config.Settings.DemoMode,
		OutPut:   output,
		CheckNet: !skipCheck,
		Timeout:  global.Config.Settings.Timeout,
		LoginInfo: srunTransfer.LoginInfo{
			Form: &global.Config.Form,
			Meta: &global.Config.Meta,
		},
	})
}
