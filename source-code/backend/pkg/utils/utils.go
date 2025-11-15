package utils

import (
	"os/exec"

	"github.com/go-logr/logr"
	"go.uber.org/zap"
)

var Logger logr.Logger

func InitLogger(level string) {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel) // Dostosuj level
	zapLog, _ := zapCfg.Build()
	Logger = zapr.NewLogger(zapLog)
}

func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	return string(output), err
}

func RunBash(cmdStr string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.Output()
	return string(output), err
}

type Game struct {
	AppName string `json:"app_name"`
	Title   string `json:"title"`
}

type Proton struct {
	Name string
	Path string
}
