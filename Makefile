install:
	@echo "Installing shell-quotes\n"
	mkdir -p "${HOME}/.config/shell-quotes"
	bin/download-quotes ${HOME}/.config/shell-quotes
	mv bin/read-quote ${HOME}/.local/bin

	@echo ""
	@echo "Installed!"
	@echo "Add the following row to your .zshrc or equivalent to run it upon shell login"
	@echo "\n\
if [[ -o login ]]; then\n\
		# Quotes from Aurelius\n\
    read-quote ${HOME}/.config/shell-quotes/aurelius.txt\n\
\n\
		# Quotes from cat-v.org\n\
    read-quote ${HOME}/.config/shell-quotes/cat-v.txt\n\
fi\n\
"

build:
	go build -i -ldflags="-s -w" -o bin/download-quotes ./download-quotes
	go build -i -ldflags="-s -w" -o bin/read-quote ./read-quote

clean:
	rm -rf bin/*
