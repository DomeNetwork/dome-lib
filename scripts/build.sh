#!/bin/bash

# Setup the environment default to `dev`.
env=$2
if [ -z "$env" ]; then
    env="dev"
fi

case $1 in
    "docs")
        # Generate the documentation using go doc.
        # Note: godoc can be used to server the documentation.
        # go doc -all ./...
        [[ $OSTYPE == 'darwin'* ]] && \
            open -n -a "Google Chrome" --args "--new-window" "http://localhost:6060/pkg/github.com/domenetwork/dome-lib/"
        godoc -http=:6060
        ;;
    *)
        echo "Unknown command, must be one of: docs"
        echo "    docs   - build all docs"
        ;;
esac
