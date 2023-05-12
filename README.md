Build and package
-----------------

Install and run Docker

go install fyne.io/fyne/v2/cmd/fyne@latest

# Windows (from Linux)

go install github.com/fyne-io/fyne-cross@latest
fyne-cross windows

# Windows (from Windows)



# Linux (from Linux)

fyne package -os linux -icon Icon.png
fyne build

# MacOS (from MacOS)

npm install --global create-dmg

fyne package -os darwin -icon Icon.png
create-dmg Samples4web.app

Release
-------
Create github release

gh release upload v0.1.0 samples4web
gh release upload v0.1.0 Samples4web.tar.xz
gh release upload v0.1.0 fyne-cross/bin/windows-amd64/Samples4web.exe

Improve and automate with Github Actions:
https://github.com/wolfgangasdf/gocalcapp/blob/master/.github/workflows/go.yml