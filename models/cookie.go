package models

type Cookie struct {
	Id         string `cookie:"id"`
	PrivateKey string `cookie:"private_key"`
	PublicKey  string `cookie:"public_key"`
}
