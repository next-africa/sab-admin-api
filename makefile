.PHONY: install setup config

install: setup
	go get -u github.com/kardianos/govendor
	pip install pyYaml

config:
	ln -s -f ../../hooks/pre-commit .git/hooks/pre-commit
	git config filter.handleSecrets.clean 'scripts/handle-secrets.py remove'
	git config filter.handleSecrets.smudge 'scripts/handle-secrets.py restore'