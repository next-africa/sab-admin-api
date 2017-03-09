#!/bin/sh


# ###############################################################
# http://tip.golang.org/misc/git/pre-commit

# git gofmt pre-commit hook
#
# ################################################################

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "$gofiles" ] && exit 0

unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] && exit 0

# Some files are not gofmt'd. Format them.
echo >&2 "Will format your unformatted go files:"

for fn in $unformatted; do
    gofmt -w $fn || exit 1
    git add  $fn || exit 1
done

exit 0