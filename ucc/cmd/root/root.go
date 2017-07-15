package root

import (
    "github.com/spf13/cobra"
    "github.com/sunshower-io/updraft/common/io"
    "os"
)

type RootCommand struct {
    *cobra.Command
    
}

type RootConfiguration struct {
    File                string
    
    ParserMode          string
    
    BackendMode         string
    
    
    Reader              *io.MultiReader
    
}

func (c *RootConfiguration) ResolveFiles() error {
    if _, er := os.Stat(c.File); er != nil {
        return er
    }
    
    file, er := os.Open(c.File)
    
    if er != nil {
        return er
    }
    
    c.Reader.Append(file)
    
    return nil
    
}

