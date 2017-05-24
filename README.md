# Teresa

Teresa is an extremely simple platform as a service that runs on top of [Kubernetes](https://github.com/kubernetes/kubernetes).  
This is the CLI to be used in conjunction with the [API](https://github.com/luizalabs/teresa-api).

## Installation

Run `make` and copy the `teresa` binary to a directory in your PATH.

## Usage

Steps to setup a new cluster and deploy a new application, assuming you already have the API running:

    $ teresa config set-cluster mycluster --server https://myapi.com
    $ teresa config use-cluster mycluster
    $ teresa login --user myuser@mydomain.com

Create a new team (optional, requires admin privileges):

    $ teresa team create myteam
    $ teresa team add-user --team myteam --user myuser@mydomain.com

Create web app:

    $ teresa app create myapp --team myteam

Create worker app:

    $ teresa app create myapp --team myteam --process-type worker
    
Create the [Procfile](https://devcenter.heroku.com/articles/procfile#procfile-naming-and-location) (optional)

    $ cat ./Procfile
    web: app.sh

Finally deploy the application

    $ teresa deploy /path/to/myapp --app myapp --description "release 1.0"

For further questions, read the help

    $ teresa -h

    $ teresa [command] --help
