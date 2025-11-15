package legendary

import (
	"epic-session/pkg/utils"
	"encoding/json"
	"fmt"
	"os/exec"
)

var venvPath string

func SetupVenv(path string) {
	venvPath = path
	if _, err := utils.RunCommand("python3", "-m", "venv", venvPath); err != nil && !os.IsExist(err) {
		utils.Logger.Error("Błąd tworzenia venv:", err)
	}
	InstallLegendary()
}

func InstallLegendary() {
	cmdStr := fmt.Sprintf("source %s/bin/activate && pip install --upgrade pip && pip install legendary-gl", venvPath)
	if _, err := utils.RunBash(cmdStr); err != nil {
		utils.Logger.Error("Błąd instalacji legendary:", err)
	}
}

func RunLegendary(args string) (string, error) {
	cmdStr := fmt.Sprintf("source %s/bin/activate && legendary %s", venvPath, args)
	return utils.RunBash(cmdStr)
}

func Login(authCode string) (string, error) {
	args := "auth"
	if authCode != "" {
		args = fmt.Sprintf("auth --code %s", authCode)
	}
	return RunLegendary(args)
}

func ListGames() ([]utils.Game, error) {
	output, err := RunLegendary("list-installed --json")
	if err != nil {
		return nil, err
	}
	var games []utils.Game
	json.Unmarshal([]byte(output), &games)
	return games, nil
}

func InstallGame(appName string) (string, error) {
	return RunLegendary(fmt.Sprintf("install %s --yes", appName))
}

func UninstallGame(appName string) (string, error) {
	return RunLegendary(fmt.Sprintf("uninstall %s --yes", appName))
}
