package api

import (
	"epic-session/pkg/config"
	"epic-session/pkg/gamescope"
	"epic-session/pkg/legendary"
	"epic-session/pkg/proton"
	"epic-session/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/status", status)
	r.GET("/login", login)
	r.GET("/list-games", listGames)
	r.POST("/install-game", installGame)
	r.POST("/uninstall-game", uninstallGame)
	r.POST("/launch-game", launchGame)
	r.GET("/list-protons", listProtons)
	r.POST("/install-proton", installProton)
	r.GET("/update-proton-list", updateProtonList)

	return r
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "running", "version": "1.0.0"})
}

func login(c *gin.Context) {
	authCode := c.Query("code")
	output, err := legendary.Login(authCode)
	if err != nil {
		utils.Logger.Error("Błąd logowania:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, output)
}

func listGames(c *gin.Context) {
	games, err := legendary.ListGames()
	if err != nil {
		utils.Logger.Error("Błąd listy gier:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

func installGame(c *gin.Context) {
	appName := c.PostForm("app_name")
	output, err := legendary.InstallGame(appName)
	if err != nil {
		utils.Logger.Error("Błąd instalacji gry:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, output)
}

func uninstallGame(c *gin.Context) {
	appName := c.PostForm("app_name")
	output, err := legendary.UninstallGame(appName)
	if err != nil {
		utils.Logger.Error("Błąd deinstalacji gry:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, output)
}

func launchGame(c *gin.Context) {
	appName := c.PostForm("app_name")
	protonPath := c.PostForm("proton_path")
	err := gamescope.LaunchInSession(appName, protonPath)
	if err != nil {
		utils.Logger.Error("Błąd uruchamiania gry:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Launched")
}

func listProtons(c *gin.Context) {
	protons, err := proton.ListProtons()
	if err != nil {
		utils.Logger.Error("Błąd listy protonów:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, protons)
}

func installProton(c *gin.Context) {
	version := c.PostForm("version")
	err := proton.InstallProton(version)
	if err != nil {
		utils.Logger.Error("Błąd instalacji protona:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Installed")
}

func updateProtonList(c *gin.Context) {
	// Symuluj update, np. pobierz najnowsze wersje z GitHub API (ale bez net, więc mock)
	c.String(http.StatusOK, "Updated list")
}
