install:
	@echo "Installing bash-quotes\n"
	mkdir -p "${HOME}/.config/bash-quotes"
	bin/download-quotes ${HOME}/.config/bash-quotes
	mv bin/read-quote ${HOME}/.local/bin

	@echo ""
	@echo "Installed!"
	@echo "Add the following row to your .zshrc or equivalent to run it upon shell login"
	@echo "\n\
if [[ -o login ]]; then\n\
    read-quote $HOME/.config/bash-quotes/quotes.txt\n\
fi\n\
"

build:
	go build -i -ldflags="-s -w" -o bin/download-quotes ./download-quotes
	go build -i -ldflags="-s -w" -o bin/read-quote ./read-quote

clean:
	rm -rf bin/*
