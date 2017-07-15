package cmd

import (
    "fmt"
    "os"
    
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/sunshower-io/updraft/ucc/cmd/pascal"
    "github.com/sunshower-io/updraft/ucc/cmd/backend"
    "github.com/sunshower-io/updraft/ucc/cmd/root"
    homedir "github.com/mitchellh/go-homedir"
    "github.com/sunshower-io/convection/core"
    "github.com/sunshower-io/updraft/common/io"
)

type RootCmdConfiguration struct {
}

func NewModeConfiguration() *ModeConfiguration {
    
    modeConfiguration := new(ModeConfiguration)
    modeConfiguration.AddParserMode(
        PASCAL,
        &pascal.PascalConfiguration{},
    )
    
    modeConfiguration.AddBackendMode(
        INTERPRETER,
        &backend.InterpreterConfiguration{},
    )
    
    return modeConfiguration
}

func (c *RootCmdConfiguration) RootConfiguration() *root.RootConfiguration {
    return &root.RootConfiguration{
        Reader: io.NewMultiReader().(*io.MultiReader),
    }
}


type EntryPoint struct {
    cmd     *root.RootCommand
    rcfg    *root.RootConfiguration
    rcmdcfg *RootCmdConfiguration
}

func (e *EntryPoint) Run() {
    e.rcmdcfg.runMode(e.cmd, e.rcfg)
    
}

func (c *RootCmdConfiguration) EntryPoint(
        cmd *root.RootCommand,
        r *root.RootConfiguration,
) *EntryPoint {
    return &EntryPoint{
        cmd:  cmd,
        rcfg: r,
    }
}

func (c *RootCmdConfiguration) NewRootCommand(
        r *root.RootConfiguration,
) *root.RootCommand {
    
    cmd := &root.RootCommand{}
    
    ccmd := &cobra.Command{
        Use:   "ucc",
        Short: "updraft compiler collection",
        Long: `
            Updraft is a modern compiler
            framework and collection written in
            pure Go for portability, speed,
            and embeddability
        `,
        RunE: func(u *cobra.Command, args []string) error {
            println("Coolbeans", r.File)
            return nil
        },
    }
    cmd.Command = ccmd
    
    cmd.PersistentFlags().StringVarP(
        &r.BackendMode,
        "execute",
        "e",
        "interpreter",
        "select the execution model",
    )
    
    
    cmd.PersistentFlags().StringVarP(
        &r.ParserMode,
        "language",
        "l",
        "pascal",
        "select the target language",
    )
    
    cmd.PersistentFlags().StringVarP(
        &r.File,
        "file",
        "f",
        "",
        "file(s) to compile",
    )
    
    cmd.PersistentFlags().StringVar(
        &cfgFile,
        "config",
        "",
        "config file (default is $HOME/.ucc.yaml)",
    )
    
    return cmd
}

func (c *RootCmdConfiguration) runMode(
        cmd *root.RootCommand,
        r *root.RootConfiguration,
) error {
    
    
    
    modeConfig := NewModeConfiguration()
    
    var (
        backendMode interface{}
        parserMode  interface{}
        er          error
    )
    
    if parserMode, er = modeConfig.ResolveParserMode(
        ParserMode(r.ParserMode),
    ); er != nil {
        return er
    }
    if backendMode, er = modeConfig.ResolveBackendMode(
        BackendMode(r.BackendMode),
    ); er != nil {
        return er
    }
    
    ctx := core.NewApplicationContext()
    
    ctx.RegisterSingleton(r)
    ctx.RegisterSingleton(cmd)
    ctx.Scan(
        c,
        parserMode,
        backendMode,
        &TraceConfiguration{},
    )
    
    return cmd.Execute()
    
}

var cfgFile string

func init() {
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := homedir.Dir()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        
        viper.AddConfigPath(home)
        viper.SetConfigName(".ucc")
    }
    
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
