package srunModels

import "github.com/Mmx233/BitSrunLoginGo/v1/transfer"

type LoginInfo struct {
	UrlLoginPage       string
	UrlGetChallengeApi string
	UrlLoginApi        string
	UrlCheckApi        string

	Ip              string
	Token           string
	EncryptedInfo   string
	Md5             string
	EncryptedMd5    string
	EncryptedChkstr string
	LoginResult     string

	Form *srunTransfer.LoginForm
	Meta *srunTransfer.LoginMeta
}
