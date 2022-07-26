go generate
go build -ldflags="-s -w -H=windowsgui"  -o bin/automate-tasks.exe
ie4uinit.exe -show
del rsrc_windows_amd64.syso