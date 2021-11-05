mkdir -p /tmp/malloc_test

runtime_dir=$(go list -f {{.Dir}} runtime)
echo "{\"Replace\":{
    \"$runtime_dir/mem_darwin.go\":\"$(pwd)/overlay/mem_darwin.go\"
}}" > /tmp/malloc_test/overlay.json
go run -overlay /tmp/malloc_test/overlay.json .

echo "test runtime"
go test -overlay /tmp/malloc_test/overlay.json -short -v runtime
