all: dev

dev: ../revel/bin/revel
	$^ run nx.com

../revel/bin/revel:
	go install github.com/revel/revel
	go install github.com/revel/cmd/revel

