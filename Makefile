build:
	go build -o ftp-proxy github.com/daniilty/hackaton-keysystems-ftp-proxy/cmd/ftp-proxy
prepare:
	cp config.example.json config.json
