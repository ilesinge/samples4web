Build and package
-----------------

Install and run Docker

go install github.com/fyne-io/fyne-cross@latest
# Windows
fyne-cross windows
# Linux
fyne package -os linux -icon Icon.png
fyne build
# MacOS
Download https://developer.apple.com/download/all/?q=Command%20Line%20Tools
fyne-cross darwin-sdk-extract --xcode-path ~/Bureau/Command_Line_Tools_for_Xcode_14.3.dmg

Release
-------
Create github release

gh release upload v0.1.0 samples4web
gh release upload v0.1.0 Samples4web.tar.xz
gh release upload v0.1.0 fyne-cross/bin/windows-amd64/Samples4web.exe

Improve and automate with Github Actions:
https://github.com/wolfgangasdf/gocalcapp/blob/master/.github/workflows/go.yml