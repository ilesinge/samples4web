# Samples4web

![Samples4web logo](https://raw.githubusercontent.com/ilesinge/samples4web/master/Logo.png)

Samples4web is a simple tool for music livecoders to expose their audio samples on the web (for [Estuary](https://github.com/dktr0/estuary) and [Strudel](https://strudel.tidalcycles.org/) platforms).

It does it by either:
- Generating a JSON file listing your samples, to be hosted somewhere on the Web (e.g. GitHub Pages)
- Exposing your samples directly through a temporary local webserver

## Use

Download and install the tool for your platform (Linux/MacOS/Windows) from the release page: https://github.com/ilesinge/samples4web/releases/

(Usage docs to complete)

## Hack

### Build and package

Install Fyne prerequisites: https://developer.fyne.io/started/

    go install fyne.io/fyne/v2/cmd/fyne@latest

#### Windows (from Linux)

Install and run Docker

    go install github.com/fyne-io/fyne-cross@latest
    fyne-cross windows

#### Windows (from Windows)

    fyne package -os windows -icon Icon.png
	

#### Linux (from Linux)

    fyne package -os linux -icon Icon.png
    fyne build

#### MacOS (from MacOS)

    npm install --global create-dmg

    fyne package -os darwin -icon Icon.png
    create-dmg Samples4web.app

### Release

Create github release

    gh release upload v0.1.0 samples4web
    gh release upload v0.1.0 Samples4web.tar.xz
    gh release upload v0.1.0 fyne-cross/bin/windows-amd64/Samples4web.exe

## Roadmap

- Add a screenshot to the README
- Allow to change the webserver port
- Add a help section
- Improve design
- Sign the MSI: https://github.com/marketplace/actions/code-sign-a-file-with-pfx-certificate
- Notarize the DMG: https://federicoterzi.com/blog/automatic-code-signing-and-notarization-for-macos-apps-using-github-actions/