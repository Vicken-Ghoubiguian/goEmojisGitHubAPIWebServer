#Put the golang image as image's base
FROM golang:latest

#
LABEL maintainer="ericghoubiguian@live.fr"

#Copy all the files and directories in the newly created directory goEmojisGitHubAPIWebServer
COPY . /goEmojisGitHubAPIWebServer

#Change work directory for the goEmojisGitHubAPIWebServer project one
WORKDIR /goEmojisGitHubAPIWebServer

#Edit the environment variable value GOPATH for the go_world_server directory
ENV GOPATH /goEmojisGitHubAPIWebServer

#Expose the docker container listening port
EXPOSE 80

#Container instruction as entrypoint: 'go run goEmojisGitHubAPIWebServer.go'
ENTRYPOINT ["go", "run", "goEmojisGitHubAPIWebServer.go"]