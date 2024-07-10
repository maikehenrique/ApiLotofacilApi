package main

import (
	"apilotofacil/configs"
	"apilotofacil/dao"
	"apilotofacil/resources"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron/v3"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	var nameMethod = "|main"

	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0777)
	}

	nameFile := configs.DiretorioLog + string(filepath.Separator) + configs.ArquivoLog
	lf := &lumberjack.Logger{
		Filename:   nameFile,
		MaxSize:    10,   // Tamanho máximo em MB do arquivo
		MaxBackups: 200,  // Número que será mantido (após são deletados)
		MaxAge:     7,    // Número de dias a ser mantido (após são deletados)
		Compress:   true, // disabled by default
		LocalTime:  true, // Horário local na máquina
	}
	log.SetOutput(io.MultiWriter(lf, os.Stdout))
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := configs.LoadFromJSON(configs.FileName)
	if err != nil {
		log.Printf(nameMethod + "|Error|" + err.Error())
		panic(err)
	}
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{Output: lf}),
		middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:         configs.XSSProtection,
			ContentTypeNosniff:    configs.ContentTypeNosniff,
			XFrameOptions:         configs.XFrameOptions,
			HSTSMaxAge:            configs.HSTSMaxAge,
			ContentSecurityPolicy: configs.ContentSecurityPolicy,
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
		}),
		middleware.Recover(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
	)

	e.HideBanner = configs.Conf.Server.HideBanner

	e.GET("/api/dashboard/:ultimos_concurso", resources.ListDashboard)
	e.POST("/api/lotofacil/monte_seu_jogo", resources.BuildYourGame)
	e.GET("/api/lotofacil", resources.ListAllLottery)
	e.GET("/api/lotofacil/:concurso", resources.ListLottery)
	e.GET("/api/lotofacil/latest", resources.SearchLastContest)
	e.GET("/api/utilidades/server_active", resources.ValidateServerAtive)

	e.Static(configs.RootPath, configs.FrontEndRoot)

	c := cron.New()
	id, _ := c.AddFunc("@every 24h", func() {
		if err := dao.UpdateDB(); err != nil {
			log.Printf(nameMethod + "|Error|" + err.Error())
		}
	})
	c.Entry(id).Job.Run()
	c.Start()

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configs.Conf.Server.Port)))
}
