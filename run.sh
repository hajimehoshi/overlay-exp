mkdir -p /tmp/malloc_test

runtime_dir=$(go list -f {{.Dir}} runtime)
echo "{\"Replace\":{
    \"$runtime_dir/mem_darwin.go\":\"$(pwd)/overlay/mem_darwin.go\",
    \"$runtime_dir/sys_darwin.go\":\"$(pwd)/overlay/sys_darwin.go\",
    \"$runtime_dir/sys_darwin_amd64.s\":\"$(pwd)/overlay/sys_darwin_amd64.s\",
    \"$runtime_dir/malloc_test.go\":\"$(pwd)/overlay/malloc_test.go\"
}}" > /tmp/malloc_test/overlay.json
go run -overlay /tmp/malloc_test/overlay.json .

echo "test runtime"
go test -overlay /tmp/malloc_test/overlay.json -short -v runtime
