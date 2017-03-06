# study-abroad-bot

Bot for students to get informations for studying abroad

# How to use run it

```console
$ go get -u github.com/kardianos/govendor
```
Make sure your default GOPATH is set up and that the path workspace's bin subdirectory is included to your PATH. You can add this to your .bash_profile:

```console
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Now everytime you want to work on this project, always make sure this repository is your GOPATH. Just run

```console
$ export GOPATH=$(pwd) && export PATH=$PATH:$(pwd)/bin
```

Note that this will change your GOPATH. In order to restore the default GOPATH, just open a new Window of the terminal
if you wanna add a dependency for a module, make sure to add the go get command for it inside the module's make file, inside the install target.

To start install dependency you can do make install.


