install:
	go install .

run-sample:
	go run github.com/PoCInnovation/DagViz -c ./samples/def

test:
	go test -v ./...

fmt:
	#goimports -w .
	gofumpt -w .
