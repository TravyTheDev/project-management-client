package main

import (
	"database/sql"
	"embed"
	"log"
	"project-management-client/cookie"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	_ "modernc.org/sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cookieStore := cookie.NewCookieStore(db)
	// Create an instance of the app structure
	app := NewApp(cookieStore)
	projectsHandler := NewProjectsHandler()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "project-management-client",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			projectsHandler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
