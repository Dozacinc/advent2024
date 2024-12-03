
export GOPRIVATE=github.com/Dozacinc



test:
	./run_test.sh


deps:
	go mod tidy
	go mod download

