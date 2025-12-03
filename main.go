package main

import (
	"embed"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

)

func do(name string) func() {
	return func() { log.Println("action:", name) }
}

//go:embed resources/locales/*.json
var TranslationsFS embed.FS

func main() {
	App := app.New()

	log.Println(lang.SystemLocale())
	lang.AddTranslationsFS(TranslationsFS, "resources/locales")

	Window := App.NewWindow(lang.L("Echoid"))

	MainMenu := fyne.NewMainMenu(
		fyne.NewMenu(
			lang.L("File"),
			&fyne.MenuItem{
				Label: lang.L("New ..."),
				ChildMenu: fyne.NewMenu(
					lang.L("New"),
					&fyne.MenuItem{
						Label:  lang.L(".epp project"),
						Action: do("New .epp project"),
					},
				),
			},
			&fyne.MenuItem{
				Label: lang.L("Open ..."),
				ChildMenu: fyne.NewMenu(
					lang.L("Open"),
					&fyne.MenuItem{
						Label: lang.L("Recent ..."),
					},
					&fyne.MenuItem{
						Label:  lang.L("Path ..."),
						Action: do("Open path ..."),
					},
				),
			},
			&fyne.MenuItem{
				Label:  lang.L("Render ..."),
				Action: do("Render ..."),
			},
		),
		fyne.NewMenu(
			lang.L("View"),
		),
		fyne.NewMenu(
			lang.L("Info"),
			&fyne.MenuItem{
				Label:  lang.L("About Echoid"),
				Action: do("About"),
			},
		),
	)

	MainTabs := container.NewDocTabs()

	TestTab := container.NewTabItemWithIcon(
		"EPP file",
		theme.FileIcon(),
		widget.NewLabel("This is a tab"),
	)

	Window.Resize(
		fyne.NewSize(600, 400),
	)
	MainTabs.SetTabLocation(
		container.TabLocationTop,
	)
	Window.SetMainMenu(MainMenu)
	MainTabs.Append(TestTab)
	Window.SetContent(MainTabs)

	Window.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	log.Println("exited")
}
