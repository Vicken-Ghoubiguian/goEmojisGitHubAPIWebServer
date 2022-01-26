# goEmojisGitHubAPIWebServer 🧮 ⚓ 🥈 🚡 🥇 🇦🇸 🥑 ⚗️ 🉑 🥉 ⚕️ 🧑‍🚀 🇳🇷 🐙 🎠 🛠️ 🗾 📽️ 🎞️ 🇮🇴 🎰 ☦️ 🇰🇵 🇲🇵 🩹 🚡 👶 ⚔️ 🥖 🇦🇲 ⏰ 🐱 😽 💕 🇦🇱 🇦🇽 💪 🥄 🇧🇸 ♒ 🚑 🐤 👽 🇦🇸 🏺 💢 👼 🏀 🚲 👙 🎱 🔋 🇨🇺 🎹 🎵 🐜 📆 🌇 🍎 🇦🇶 💔 👾 🧬 🪕 🇯🇪 🦠 🤖 ☄️ 🌃 ♾️ ℹ️ 🤿

Web application written in Go to collect and display all the emojis 🧮 ⚓ 🥈 🚡 🥇 🇦🇸 🥑 ⚗️ 🉑 🥉 ⚕️ 🧑‍🚀 🇳🇷 🐙 🎠 🛠️ 🗾 📽️ 🎞️ 🇮🇴 🎰 ☦️ 🇰🇵 🇲🇵 🩹 🚡 👶 ⚔️ 🥖 🇦🇲 ⏰ 🐱 😽 💕 🇦🇱 🇦🇽 💪 🥄 🇧🇸 ♒ 🚑 🐤 👽 🇦🇸 🏺 💢 👼 🏀 🚲 👙 🎱 🔋 🇨🇺 🎹 🎵 🐜 📆 🌇 🍎 🇦🇶 💔 👾 🧬 🪕 🇯🇪 🦠 🤖 ☄️ 🌃 ♾️ ℹ️ 🤿 used in GitHub via its API...

## Contents

1. [Introduction](#introduction)
2. [Presentation](#presentation)
3. [Project's structure](#project_s_structure)
4. [How was this project developed ?](#how_was_this_project_developed)
5. [How does this project work ?](#how_does_this_project_work)
6. [How to use it ?](#how_to_use_it)
    * [By source code](#by_source_code)
    * [By Docker](#by_docker)
        * [By Dockerfile](#by_dockerfile)
        * [By Docker Hub](#by_docker_hub)
7. [How to deploy it ?](#how_to_deploy_it)
    * [From source code](#from_source_code)
    * [From Docker](#from_docker)
        * [From Dockerfile](#from_dockerfile)
        * [From Docker Hub](#from_docker_hub)
    * [With Balena](#with_balena)
8. [Where to use it ?](#where_to_use_it)
9. [Useful links](#useful_links)
10. [Conclusion](#conclusion)

<a name="introduction"></a>
## Introduction

This is a web application which collect and display all the emojis used in GitHub via its API. This web application was developed when my contract at IMERIR ended successfully. My interest in emojis then grew and so I started to develop projects on emojis. This is the first of a long list putting emojis front and center.

<a name="presentation"></a>
## Presentation

Coming soon...

<a name="project_s_structure"></a>
## Project's structure

What is this project and this web application made of ? Here are all of the elements... Therefore:

* The Go file `goEmojisGitHubAPIWebServer.go`, which is the main file, the executable file, of this web application,
* The `assets` folder
* The `tmpl` folder
* The `Dockerfile` file, 
* The `README.md` file, which is a markdown file, which presents the current project, the web application and serves as a manual for the latter,
* The `LICENSE` file is the software license which this web application is under, so the 'MIT' one in this present case,
* The `.gitattributes` file
* The `.gitignore` file

<a name="how_was_this_project_developed"></a>
## How was this project developed ?

Coming soon...

<a name="how_does_this_project_work"></a>
## How does this project work ?

Coming soon...

<a name="how_to_use_it"></a>
## How to use it ?

There are 2 ways to use the ```goEmojisGitHubAPIWebServer``` web application: through source code and through Docker. Here they are:

<a name="by_source_code"></a>
### By source code...

It will be first discuss the installation, deployment and use of this web application by its source code.

First of all, you must clone the project ```goEmojisGitHubAPIWebServer``` from GitHub using this command:

```bash
git clone https://github.com/Vicken-Ghoubiguian/goEmojisGitHubAPIWebServer
```

In a second time, you must go to the directory of the cloned project using the following command:

```bash
cd goEmojisGitHubAPIWebServer
```

In a third time, 

```bash
export GOPATH=${PWD}
```

Finaly, it's time to run this web app by executing this command bellow:

```bash
go run goEmojisGitHubAPIWebServer.go
```
Now the web app is working completely fine. To see this, please click [here](http://localhost).

<a name="by_docker"></a>
### By Docker...

<a name="by_dockerfile"></a>
#### By Dockerfile...

```bash
git clone https://github.com/Vicken-Ghoubiguian/goEmojisGitHubAPIWebServer
```

```bash
cd goEmojisGitHubAPIWebServer
```

```bash
docker image build -t goemojisgithubapiwebserver:latest .
```

__clarification:__ the character `.` specified at the end of the command refers to the current directory.

```bash
docker image ls
```

```bash
docker image history goemojisgithubapiwebserver:latest
```

```bash
docker container run -d --name goemojisgithubapiwebserver -p 80:80 goemojisgithubapiwebserver:latest
```

```bash
<container_s_IP_address>:80
```

```bash
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' goemojisgithubapiwebserver
```

<a name="by_docker_hub"></a>
#### By Docker Hub...

And you can also deploy this web application using Docker and without the associated Dockerfile. The web application `goEmojisGitHubAPIWebServer 🧮 ⚓ 🥈 🚡 🥇 🇦🇸 🥑 ⚗️ 🉑 🥉 ⚕️ 🧑‍🚀 🇳🇷 🐙 🎠 🛠️ 🗾 📽️ 🎞️ 🇮🇴 🎰 ☦️ 🇰🇵 🇲🇵 🩹 🚡 👶 ⚔️ 🥖 🇦🇲 ⏰ 🐱 😽 💕 🇦🇱 🇦🇽 💪 🥄 🇧🇸 ♒ 🚑 🐤 👽 🇦🇸 🏺 💢 👼 🏀 🚲 👙 🎱 🔋 🇨🇺 🎹 🎵 🐜 📆 🌇 🍎 🇦🇶 💔 👾 🧬 🪕 🇯🇪 🦠 🤖 ☄️ 🌃 ♾️ ℹ️ 🤿` has an official image on Docker Hub. So please run the following command:

```bash
docker pull wicken/goemojisgithubapiwebserver:latest
```
Now it's time to create the Docker container and get it working as discussed above.

<a name="how_to_deploy_it"></a>
## How to deploy it ?

<a name="from_source_code"></a>
### From source code

Like see in [this section above](#how_to_use_it),

```bash
git clone https://github.com/Vicken-Ghoubiguian/goEmojisGitHubAPIWebServer
```

```bash
cd goEmojisGitHubAPIWebServer
```

```bash
export GOPATH=${PWD}
```

```bash
go run goEmojisGitHubAPIWebServer.go
```

Now ! It's your turn to play...

<a name="from_docker"></a>
### From Docker

It will be now discuss the installation, deployment and use of this web application by Docker.

<a name="from_dockerfile"></a>
#### From Dockerfile

First of all, you must clone the project ```goEmojisGitHubAPIWebServer``` from GitHub using this command:

```bash
git clone https://github.com/Vicken-Ghoubiguian/goEmojisGitHubAPIWebServer
```
Now, it is time for you to position yourself in the `goEmojisGitHubAPIWebServer` repository, using this bellow command:

```bash
cd goEmojisGitHubAPIWebServer
```
Then, you'll build the Docker image following the bellow command:

```bash
docker image build -t goemojisgithubapiwebserver:latest .
```

__clarification:__ the character `.` specified at the end of the command refers to the current directory.

If you want to list all Docker images on your system, you must follow the next command:

```bash
docker image ls
```

```bash
docker image history goemojisgithubapiwebserver:latest
```

```bash
docker container run -d --name goemojisgithubapiwebserver -p 80:80 goemojisgithubapiwebserver:latest
```
To access to the web application from your browser, you must type the Docker container's IP adress following the colon (':') character and the port number (here '80'). It is show in the following command:

```bash
<container_s_IP_address>:80
```


```bash
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' goemojisgithubapiwebserver
```

<a name="from_docker_hub"></a>
#### From Docker Hub

```bash
docker pull wicken/goemojisgithubapiwebserver:latest
```

<a name="with_balena"></a>
### With Balena

<a name="where_to_use_it"></a>
## Where to use it ?

You can use this web application in many cases. You can use it, for example, in an educational case: to teach the Go language and web APIs to schoolchildren and students through an exciting and rewarding project. This web application could also be used for computer enthusiasts, to also learn the Go language and development with APIs. Of course, this web application can be used for other types of audiences and other ideas can be found... It's now your turn to play...

<a name="useful_links"></a>
## Useful links

* [Emojis - GitHub Docs](https://docs.github.com/en/rest/reference/emojis),
* [Emojis from GitHub API - link](https://api.github.com/emojis),
* [Experimenting with JSON (structured and unstructured data) and Go - Henry S. Coelho](https://hcoelho.com/blog/62/Experimenting_with_JSON_structured_and_unstructured_data_and_Go),
* [jQuery UI - official website](https://jqueryui.com),
* [jQuery UI's ThemeRoller - official website](https://jqueryui.com/themeroller/)

<a name="conclusion"></a>
## Conclusion

Coming soon...
