OUTPUT_DIR=_output

myoci: init
	go build -o ${OUTPUT_DIR}/myoci cmd/main.go 

init:
	mkdir -p ${OUTPUT_DIR}


