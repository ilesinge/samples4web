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