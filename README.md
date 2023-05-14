# Samples4web

![Samples4web logo](https://raw.githubusercontent.com/ilesinge/samples4web/main/Logo.png)

Samples4web is a simple tool for music livecoders to expose their audio samples on the web (for [Estuary](https://github.com/dktr0/estuary) and [Strudel](https://strudel.tidalcycles.org/) platforms).

It does it by either:
- Generating a JSON file listing your samples, to be hosted somewhere on the Web (e.g. GitHub Pages)
- Exposing your samples directly through a temporary local webserver

## Use

- Download and install the tool for your platform (Linux/MacOS/Windows) from the [release page](https://github.com/ilesinge/samples4web/releases/) 
  - Caveat: on Windows, the MSI is not signed yet, so it may trigger when installing.

- Launch the app

![Samples4web screenshot](https://raw.githubusercontent.com/ilesinge/samples4web/main/Screenshot.png)

- If you want to host your samples elsewhere (such as GitHub Pages) and you target Strudel, fill the "Base URL" field with the target host URL
- Pick the folder containing the samples you want to expose
- Click on the "Save JSON" button
  - You are now ready to host your files somewhere, including the strudel.json and estuary.json files generated in the chosen folder
- If you want to play with your samples locally, click the "Start local webserver" button
- The samples and the JSON files are now exposed on http://localhost:3000/

## Hack

- Install Go: https://go.dev/doc/install
- Install Fyne prerequisites: https://developer.fyne.io/started/
- Install Fyne: `go install fyne.io/fyne/v2/cmd/fyne@latest`
- Run the app: `go run main.go`
- Build the app:
  - `fyne build` to generate the barebones binary
  - `fyne package -os <linux|windows|macos> -icon Icon.png`
- Package (DMG/MSI) and release process: Look at the [GitHub Actions file](https://github.com/ilesinge/samples4web/blob/main/.github/workflows/latest_build.yml)

## Roadmap

- Allow to change the webserver port
- Add a help section
- Improve design
- Sign the MSI: https://github.com/marketplace/actions/code-sign-a-file-with-pfx-certificate
- Notarize the DMG: https://federicoterzi.com/blog/automatic-code-signing-and-notarization-for-macos-apps-using-github-actions/