
# clone-gitlab-repo-recursively
Allows to clone a tree of gitlab repositories. 

## Init go projects
In order to init a golang project go to ***ssh*** or ***https*** folder and enter : <br>
- `go mod init main.go` 
- `go mod tidy`

## Usage

### Clone using SSH protocol

    go run main.go ID_GROUP  TOKEN_GITLAB PATH_DIRECTORY  PATH_TO_SSH_PRIVATE_KEY SHS_PASSWORD <br>

or after a build: `env GOOS=darwin GOARCH=amd64 go build main.go` 

    ./main ID_GROUP  TOKEN_GITLAB PATH_DIRECTORY  PATH_TO_SSH_PRIVATE_KEY SHS_PASSWORD

### Clone using HTTPS protocol

    go run main.go ID_GROUP  TOKEN_GITLAB PATH_DIRECTORY  GITLAB_USER

or after a build : `env GOOS=darwin GOARCH=amd64 go build main.go` 
`./main ID_GROUP  TOKEN_GITLAB PATH_DIRECTORY  GITLAB_USER`


## Windows
For windows, only the HTTPS version is working.
To build a windows executable : `env GOOS=windows GOARCH=amd64 go build main.go` 


