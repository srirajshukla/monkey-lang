Compiling the wasm module:

cd to web package
```
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o main.wasm
```

Copy the wasm_exec.js file from GO installtion directory