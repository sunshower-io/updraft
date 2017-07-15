package cmd

import (
    "fmt"
    "bytes"
)



/**
Which syntax are we targeting?
 */
type ParserMode string

const (
    PASCAL ParserMode = "pascal"
    GROOVY ParserMode = "groovy"
)


/**
What's the runtime mode?
 */
type BackendMode string

const (
    COMPILER    BackendMode = "compiler"
    INTERPRETER BackendMode = "interpreter"
)



type ModeConfiguration struct {
    parserModes         map[ParserMode]interface{}
    backendModes        map[BackendMode]interface{}
}

func(m *ModeConfiguration) ResolveBackendMode(
        mode BackendMode,
) (interface{}, error) {
    if pmode, exists := m.backendModes[mode]; exists {
        return pmode, nil
    }
    return nil, m.noSuchBackendMode(string(mode))
}



func(m *ModeConfiguration) ResolveParserMode(
        mode ParserMode,
) (interface{}, error) {

    if pmode, exists := m.parserModes[mode]; exists {
        return pmode, nil
    }
    return nil, m.noSuchParserMode(string(mode))
}




func (m *ModeConfiguration) AddParserMode(
        mode ParserMode,
        cfg interface{},
) {
    if m.parserModes == nil {
        m.parserModes = make(map[ParserMode]interface{})
    }

    m.parserModes[mode] = cfg
}


func(m *ModeConfiguration) AddBackendMode(
        mode BackendMode,
        cfg interface{},
) {
    if m.backendModes == nil {
        m.backendModes = make(map[BackendMode]interface{})
    }
    
    m.backendModes[mode] = cfg
}



func (m *ModeConfiguration) noSuchBackendMode(requestedMode string) error {
    
    var buf bytes.Buffer
    
    for k := range m.backendModes{
        buf.WriteString(fmt.Sprintf("\t%s\n", k))
    }
    
    return fmt.Errorf(
        "Requested backend mode " +
                "'%s' is not available.  " +
                "Known back-ends are \n %s",
        requestedMode,
        buf.String(),
    )
}

func (m *ModeConfiguration) noSuchParserMode(requestedMode string) error {
    
    var buf bytes.Buffer
    
    for k := range m.parserModes {
        buf.WriteString(fmt.Sprintf("\t%s\n", k))
    }
    
    
    return fmt.Errorf(
        "Requested language mode " +
        "'%s' is not available.  " +
        "Known front-ends are \n %s",
        requestedMode,
        buf.String(),
    )
    
}

