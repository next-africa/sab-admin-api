.PHONY: install setup config

install: setup
	go get google.golang.org/appengine
	go get github.com/next-africa/mux
	go get github.com/next-africa/context
	go get github.com/next-africa/testify
	go get github.com/next-africa/graphql-go
	go get github.com/next-africa/graphql-go-handler
	go get github.com/next-africa/graphql-go-relay
	pip install pyYaml

config:
	ln -s -f ../../hooks/pre-commit .git/hooks/pre-commit
	git config filter.handleSecrets.clean 'scripts/handle-secrets.py remove'
	git config filter.handleSecrets.smudge 'scripts/handle-secrets.py restore'