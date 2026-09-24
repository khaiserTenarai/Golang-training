# poweshell script
param( $task)

if ($task -eq "run"){
    go run .
}
elseif($task -eq "build"){
    go build .
}
elseif($task -eq "fmt"){
    go fmt ./...
}
elseif($task -eq "test"){
    go test ./...
}
elseif($task -eq "vet"){
    go vet ./...
}

