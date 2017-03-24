.PHONY: install setup config

install: setup
	go get google.golang.org/appengine
	go get github.com/gorilla/mux
	go get github.com/gorilla/context
	go get github.com/stretchr/testify
	pip install pyYaml

config:
	ln -s -f ../../hooks/pre-commit .git/hooks/pre-commit
	git config filter.handleSecrets.clean 'scripts/handle-secrets.py remove'
	git config filter.handleSecrets.smudge 'scripts/handle-secrets.py restore'