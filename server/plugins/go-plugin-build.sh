for f in ./*.go; do
    scriptName="${f%.*}"
    go build -buildmode=plugin -o "$scriptName.so" "$scriptName.go"
done