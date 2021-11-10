cd web
yarn
yarn build

cd ..
echo "Start building backend"
go mod download
mkdir dist
set CGO_ENABLED=0
set GOOS=windows
go build --tags "release" -ldflags="-w -s" -o dist/urlshortener.exe ./main.go
set CGO_ENABLED=
set GOOS=
echo "Finish building backend"

REM Start service by following command
REM cd dist
REM set CONFIGFILE=../config/config.json&set WEBPATH=../web/dist&./urlshortener&set CONFIGFILE=&set WEBPATH=
