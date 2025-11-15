package gamescope

import (
	"epic-session/pkg/legendary"
	"epic-session/pkg/utils"
	"fmt"
	"os"
	"os/exec"
)

const gamescopeCmd = "gamescope -f -- " // Pełna sesja, dostosuj do config

func LaunchInSession(appName, protonPath string) error {
	prefix := fmt.Sprintf("%s/pfx_%s", os.Getenv("HOME"), appName)
	os.MkdirAll(prefix, 0755)

	env := os.Environ()
	env = append(env, fmt.Sprintf("STEAM_COMPAT_DATA_PATH=%s", prefix))
	env = append(env, fmt.Sprintf("STEAM_COMPAT_CLIENT_INSTALL_PATH=%s", prefix))

	cmdStr := fmt.Sprintf("source %s/bin/activate && legendary launch %s --no-wine --wrapper '\"%s/proton\" run'", legendary.venvPath, appName, protonPath)
	fullCmd := gamescopeCmd + cmdStr

	cmd := exec.Command("bash", "-c", fullCmd)
	cmd.Env = env
	return cmd.Run()
}
