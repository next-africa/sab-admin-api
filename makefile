.PHONY: install setup config

install: setup
	go get google.golang.org/appengine
	pip install pyYaml

config:
	git config filter.gofmt.clean 'go fmt'
	git config filter.handleSecrets.clean 'scripts/handle-secrets.py remove'
	git config filter.handleSecrets.smudge 'scripts/handle-secrets.py restore'