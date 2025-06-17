$env:GOOS="linux"
$env:CGO_ENABLED="0"
$env:GOARCH="arm64"
$env:GOTOOLCHAIN="auto"

# -tags lambda.norpc
#   reduces binary size by excluding the remote procedure call component
# -ldflags="-s -w"
#   reduces binary size by removing debug libraries
go build -tags lambda.norpc -ldflags="-s -w" -o ./dist/bootstrap main.go

cd ./dist

build-lambda-zip -output bootstrap.zip bootstrap
