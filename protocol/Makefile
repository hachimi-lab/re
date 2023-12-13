# CUSTOMIZE: Change the following variables to suit your needs
CODE_DIR := generated/code
DOC_DIR := generated/doc

.PHONY: protoc
protoc:
	mkdir -p $(CODE_DIR) && mkdir -p $(DOC_DIR) && \
	protoc -I ./third -I ./define --go_out=./ --validate_out="lang=go:." --doc_out=./$(DOC_DIR) --doc_opt=html,index.html define/*.proto
