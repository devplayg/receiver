package engine

import (
	"os"
	"path/filepath"
)

type Option interface {
	Debug() bool
	Name() string
	Version() string
	Description() string
}

type Server interface {
	Start(engine *Engine) error
}

type Engine struct {
	name        string
	description string
	version     string
	debug       bool

	Option     Option
	server     Server
	WorkingDir string

	//	//Config      map[string]string
	//	debug      bool
	//	//cpuCount    int
	//	//processName string
	//	workingDir  string
	//	//TempDir     string
	//	//logOutput   int // 0: STDOUT, 1: File
	//	//LogPrefix string
	//	//DB          *sql.DB
	//	//TimeZone    *time.Location
}

//func NewEngine(option obj.Option) *Engine {
//}
//
//
func NewEngine(option Option) *Engine {
	e := Engine{
		Option: option,
		//name:        option.Name(),
		//description: option.Version(),
		//version:     option.Version(),
		//debug:       option.Debug(),
	}
	//e.LogPrefix = "[" + e.processName + "] "
	workingDir, _ := filepath.Abs(os.Args[0])
	e.WorkingDir = filepath.Dir(workingDir)
	return &e
}

func (e *Engine) Start(server Server) error {
	err := server.Start(e)
	if err != nil {
		panic(err)
	}

	return nil
}
