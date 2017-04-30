# study-abroad-bot

Bot for students to get informations for studying abroad

# configure repository

Run `make config` to configure this repository with the required clean and smudge filters. Clean and smudge filters are used to clean (format,
remove secret values, etc...) before commiting.

Make sure your default GOPATH is set up and that the path workspace's bin subdirectory is included to your PATH. You can add this to your .bash_profile:
```console
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Now everytime you want to work on this project, always make sure this repository is your GOPATH. Just run

```console
$ export GOPATH=$(pwd) && export PATH=$PATH:$(pwd)/bin
```

# Dependencies
Dependencies are managed using govendor.

First you should install govendor by running:
```console
$ go get -u github.com/kardianos/govendor
```

After that, you should install all dependencies listed in the vendor.json file of each package.
To do that you need to run ```console $ govendor sync```inside the package directory.

# Running
You can run individual modules inside the gae folder. To run the admin module locally, go into the gae/admin-module and run:
```console`
$ dev_appserver.py app.yaml
``



