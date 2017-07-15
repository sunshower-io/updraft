package main

import (
	"gitlab.com/sunshower.io/convection/core"
	"gitlab.com/sunshower.io/updraft/ucc/cmd"
)


func init() {
	core.SetLogger(&core.NoOpLogger{})
}


func main() {
	ctx := core.NewApplicationContext()
	
	ctx.Scan(&cmd.RootCmdConfiguration{})
	rootc := ctx.GetByName("EntryPoint").(*cmd.EntryPoint)
	rootc.Run()
}
