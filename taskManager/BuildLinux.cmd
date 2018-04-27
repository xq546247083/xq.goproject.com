set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0 
go build -o taskManager.linux

if %ERRORLEVEL% gtr 0 (
    pause
)