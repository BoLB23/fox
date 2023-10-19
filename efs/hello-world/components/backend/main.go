package main

import "github.com/xigxog/kubefox/libs/core/kit"

func main() {
	k := kit.New()
	k.Default(sayHello)
	k.Start()
}

func sayHello(k kit.Kontext) error {
	who := k.EnvDef("who", "World")
	k.Log().Infof("The who is '%s'!", who)

	return k.Resp().SendStr(who)
}
