// samples4web
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// BaseSample for bank description
type BaseSample struct {
	Bank string
	Path string
}

// EstuarySample for json
type EstuarySample struct {
	Bank   string `json:"bank"`
	Number int    `json:"n"`
	Type   string `json:"type"`
	URL    string `json:"url"`
}

func isSample(file fs.FileInfo) bool {
	return !file.IsDir() && (strings.HasSuffix(file.Name(), ".wav") || strings.HasSuffix(file.Name(), ".mp3"))
}

func listSamplesFilesInFolder(folderPath string) ([]BaseSample, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	samples := []BaseSample{}

	for _, file := range files {
		if isSample(file) {
			samples = append(samples, BaseSample{
				Bank: filepath.Base(folderPath),
				Path: filepath.Join(folderPath, file.Name()),
			})
		}
		if file.IsDir() {
			recursiveFiles, err := listSamplesFilesInFolder(filepath.Join(folderPath, file.Name()))
			if err != nil {
				return nil, err
			}
			samples = append(samples, recursiveFiles...)
		}
	}
	return samples, nil
}

func generateStrudelJSON(folderPath string, samples []BaseSample, baseURL string) error {
	assets := map[string]any{
		"_base": baseURL,
	}
	for _, sample := range samples {
		samplename := sample.Bank
		relativePath := strings.TrimPrefix(filepath.ToSlash(sample.Path), folderPath+"/")
		if assets[samplename] == nil {
			assets[samplename] = []string{}
		}
		assets[samplename] = append(assets[samplename].([]string), relativePath)
	}
	data, _ := json.MarshalIndent(assets, " ", "\t")
	err := os.WriteFile(filepath.Join(folderPath, "strudel.json"), data, 0644)
	return err
}

func generateEstuaryJSON(folderPath string, samples []BaseSample) error {
	assets := []EstuarySample{}
	sampleNumber := map[string]int{}
	for _, sample := range samples {
		relativePath := strings.TrimPrefix(filepath.ToSlash(sample.Path), folderPath+"/")
		samplename := sample.Bank
		estuarysample := EstuarySample{
			Bank:   samplename,
			Number: sampleNumber[samplename],
			Type:   "audio",
			URL:    relativePath,
		}
		sampleNumber[samplename]++
		assets = append(assets, estuarysample)
	}
	data, _ := json.MarshalIndent(assets, " ", "\t")
	err := os.WriteFile(filepath.Join(folderPath, "estuary.json"), data, 0644)
	return err
}

func customHeaders(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
		w.Header().Set("pragma", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		handler.ServeHTTP(w, r)
	}
}

func main() {

	app := app.New()
	window := app.NewWindow("samples4web")

	hello := widget.NewLabel("Hello Livecoder!")

	urlentry := widget.NewEntry()
	urlentry.SetText("http://localhost:3000/")

	folder := widget.NewEntry()

	output := widget.NewLabel("")
	outputEstuary := widget.NewLabel("")
	outputEstuaryEditable := widget.NewEntry()
	outputStrudel := widget.NewLabel("")
	outputStrudelEditable := widget.NewEntry()

	containerEstuary := container.New(layout.NewFormLayout(), outputEstuary)
	containerStrudel := container.New(layout.NewFormLayout(), outputStrudel)

	server := &http.Server{Addr: ":3000"}

	window.SetContent(container.New(layout.NewPaddedLayout(), container.NewVBox(
		hello,
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Base URL"), urlentry,
			widget.NewButton("Pick Folder", func() {
				dialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
					if uri != nil {
						folder.SetText(uri.Path())
					}
				}, window)
				dialog.SetDismissText("Exit")
				dialog.Show()
			}),
			folder,
		),
		container.New(layout.NewCenterLayout(), widget.NewButton("Save JSON", func() {
			log.Println("URL was: ", urlentry.Text)
			log.Println("Folder was: ", folder.Text)
			if folder.Text != "" {
				samples, err := listSamplesFilesInFolder(folder.Text)
				if err != nil {
					fmt.Println(err)
				} else {
					generateStrudelJSON(folder.Text, samples, urlentry.Text)
					generateEstuaryJSON(folder.Text, samples)
					output.SetText("JSON saved")
				}
			} else {
				output.SetText("No folder selected")
			}
		})),
		container.New(layout.NewCenterLayout(), widget.NewButton("Start local webserver", func() {
			go func() {
				ctx, _ := context.WithTimeout(context.Background(), time.Second)
				server.Shutdown(ctx)
				server = &http.Server{Addr: ":3000"}
				mux := http.NewServeMux()
				fs := http.FileServer(http.Dir(folder.Text))
				mux.Handle("/", customHeaders(fs))
				server.Handler = mux
				log.Print("Listening on http://localhost:3000/")
				output.SetText("Listening on http://localhost:3000/")
				outputEstuary.SetText("On Estuary? Type:")
				outputEstuaryEditable.SetText("!reslist \"http://localhost:3000/estuary.json\"")
				containerEstuary.RemoveAll()
				containerEstuary.Add(outputEstuary)
				containerEstuary.Add(outputEstuaryEditable)
				containerEstuary.Refresh()
				outputStrudel.SetText("On Strudel? Type:")
				outputStrudelEditable.SetText("samples('http://localhost:3000/strudel.json')")
				containerStrudel.RemoveAll()
				containerStrudel.Add(outputStrudel)
				containerStrudel.Add(outputStrudelEditable)
				containerStrudel.Refresh()
				err := server.ListenAndServe()
				if err != nil {
					log.Print(err)
				}
			}()
		})),
		layout.NewSpacer(),
		containerEstuary,

		containerStrudel,
		container.NewHBox(layout.NewSpacer(), output),
	)))

	window.Resize(fyne.NewSize(800, 600))

	window.ShowAndRun()
}
