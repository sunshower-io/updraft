package cmd

import (
    "github.com/spf13/cobra"
    "github.com/sunshower-io/updraft/common/compiler"
    "github.com/sunshower-io/updraft/ucc/cmd/root"
    "github.com/sunshower-io/updraft/front/parser"
    "github.com/sunshower-io/updraft/pascal/common"
)

type TraceConfiguration struct {
    summarize bool
    
    tokens   bool
    compiler compiler.Compiler
}

type TraceCommand struct {
    *cobra.Command
}

func (t *TraceConfiguration) TraceCommand(
        cmd *root.RootCommand,
        c compiler.Compiler,
        cfg *root.RootConfiguration,
) {
    
    var (
        summarize bool
        tokens    bool
        symbols   bool
    )
   
    
    tc := &cobra.Command{
        Use:   "trace",
        Short: "trace various aspects of configuration and execution",
        Long:  "Print compiler statistics such as performance, AST, IR, parsing progress, etc.",
        RunE: func(ccmd *cobra.Command, args []string) error {
            
           
            err := cfg.ResolveFiles()
            
            if err != nil {
                return err
            }
            
            
            if summarize {
                c.AddListener(
                    compiler.PARSING,
                    &common.ParserMessageListener{
                        DumpSymbols : symbols,
                    },
                )
            }
            
            if tokens {
                c.AddListener(
                    compiler.PARSING,
                    &parser.ParserMessageListener{},
                )
            }
            c.Compile()
            return nil
        },
    }
    
    
    tc.Flags().BoolVarP(
        &symbols,
        "symbols",
        "y",
        false,
        "Dump symbol tables",
    )
    
    
    
    tc.Flags().BoolVarP(
        &summarize,
        "summarize",
        "s",
        false,
        "print execution statistics",
    )
    
    tc.Flags().BoolVarP(
        &tokens,
        "tokens",
        "t",
        false,
        "display program tokenization information",
    )
   
    cmd.AddCommand(tc)
}
