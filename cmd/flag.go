package cmd

import (
	"strings"

	"github.com/urfave/cli"
)

const (
	DEFAULT_LOG_LEVEL           = 1
	DEFAULT_LOG_FILE_PATH       = "./Log/"
	DEFAULT_BLOCK_CHAIN_RPC_URL = "http://localhost:8545"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: DEFAULT_LOG_LEVEL,
	}
	LogDirFlag = cli.StringFlag{
		Name:  "log-dir",
		Usage: "log output to the file",
		Value: DEFAULT_LOG_FILE_PATH,
	}
	RpcUrlFlag = cli.StringFlag{
		Name:  "chain-rpc-url",
		Usage: "Set block chain rpc url",
		Value: DEFAULT_BLOCK_CHAIN_RPC_URL,
	}
	DisableLogFileFlag = cli.BoolFlag{
		Name:  "disable-log-file",
		Usage: "Discard log output to file",
	}
	PortFlag = cli.Uint64Flag{
		Name:  "server-port",
		Usage: "server port",
		Value: 8088,
	}
	ConfigFileFlag = cli.StringFlag{
		Name:  "config-file",
		Usage: "config file",
		Value: "./config.json",
	}
	OperationFlag = cli.StringFlag{
		Name:  "operation",
		Usage: "new-wallet | register-did",
	}
)

// GetFlagName deal with short flag, and return the flag name whether flag name have short name
func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
