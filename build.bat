go build vendor\github.com\akavel\rsrc\rsrc.go
rsrc -ico icon.ico
go build -ldflags="-s -w -H=windowsgui" -o bin/automate-tasks.exe
ie4uinit.exe -show
del rsrc_windows_amd64.syso
del rsrc.exe