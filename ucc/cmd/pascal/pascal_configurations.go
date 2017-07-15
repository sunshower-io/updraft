package pascal

import (
    "github.com/sunshower-io/updraft/ucc/cmd/root"
    "github.com/sunshower-io/updraft/pascal/pascal"
    "github.com/sunshower-io/updraft/common/compiler"
)

type PascalConfiguration struct {



}

func (p *PascalConfiguration) NewCompiler(
        cmd *root.RootCommand,
        cfg *root.RootConfiguration,
) compiler.Compiler {
  
    
    
    return pascal.NewPascal(cfg.Reader)
    
    
    
}