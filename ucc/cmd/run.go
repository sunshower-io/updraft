package cmd

import (
    "gitlab.com/sunshower.io/updraft/ucc/cmd/root"
    "gitlab.com/sunshower.io/updraft/common/compiler"
    "github.com/spf13/cobra"
)

type RunCmdConfiguration struct {


}

func (f *RunCmdConfiguration) Run(
        r *root.RootCommand,
        cmp compiler.Compiler,
) *cobra.Command {
    
    
    
    cmd := &cobra.Command{
        Use: "run",
        Short: "Execute current program",
        RunE: func(cmd *cobra.Command, args []string) error {
            cmp.Compile()
            return nil
        },
    }
    
    r.AddCommand(cmd)
    return cmd
}

