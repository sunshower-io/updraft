package cmd

import (
    "github.com/spf13/cobra"
    "github.com/sunshower-io/updraft/common/compiler"
    "github.com/sunshower-io/updraft/ucc/cmd/root"
    "github.com/sunshower-io/updraft/front/parser"
    "github.com/sunshower-io/updraft/pascal/common"
    common2 "github.com/sunshower-io/updraft/common"
    "github.com/sunshower-io/updraft/common/ir"
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
        printIr   bool
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
                    common2.PARSING,
                    &common.ParserMessageListener{
                        DumpSymbols : symbols,
                    },
                )
            }
            
            if tokens {
                c.AddListener(
                    common2.PARSING,
                    &parser.ParserMessageListener{},
                )
            }
            
            
            
            result := c.Compile()
            
            if printIr {
                str := new(ir.JsonExecutionModelPrinter).Print(result.GetExecutionModel())
                println(str)
            }
            return nil
        },
    }
    
    tc.Flags().BoolVarP(
        &printIr,
        "intermediate",
        "p",
        false,
        "Print IR form",
    )
    
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
