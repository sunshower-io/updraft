package pascal

import (
    "gitlab.com/sunshower.io/updraft/ucc/cmd/root"
    "gitlab.com/sunshower.io/updraft/pascal/pascal"
    "gitlab.com/sunshower.io/updraft/common/compiler"
)

type PascalConfiguration struct {



}

func (p *PascalConfiguration) NewCompiler(
        cmd *root.RootCommand,
        cfg *root.RootConfiguration,
) compiler.Compiler {
  
    
    
    return pascal.NewPascal(cfg.Reader)
    
    
    
}