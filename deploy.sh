
cd ./src
GOOS=js GOARCH=wasm go build -o parallel_shooter.wasm github.com/ntakahashi016/parallel_shooter

cd ../
mv ./src/parallel_shooter.wasm ./public/
