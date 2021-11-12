# goEmojisGitHubAPIWebServer 🧮 ⚓ 🥈 🚡 🥇 🇦🇸 🥑 ⚗️ 🉑 🥉 ⚕️ 🧑‍🚀 🇳🇷 🐙 🎠 🛠️ 🗾 📽️ 🎞️ 🇮🇴 🎰 ☦️ 🇰🇵 🇲🇵 🩹 🚡 👶 ⚔️ 🥖 🇦🇲 ⏰ 🐱 😽 💕 🇦🇱 🇦🇽 💪 🥄 🇧🇸 ♒ 🚑 🐤 👽 🇦🇸 🏺 💢 👼 🏀 🚲 👙 🎱 🔋 🇨🇺 🎹 🎵 🐜 📆 🌇 🍎 🇦🇶 💔 👾 🧬 🪕 🇯🇪 🦠 🤖 ☄️ 🌃 ♾️ ℹ️ 🤿

Web application written in Go to collect and display all the emojis 🧮 ⚓ 🥈 🚡 🥇 🇦🇸 🥑 ⚗️ 🉑 🥉 ⚕️ 🧑‍🚀 🇳🇷 🐙 🎠 🛠️ 🗾 📽️ 🎞️ 🇮🇴 🎰 ☦️ 🇰🇵 🇲🇵 🩹 🚡 👶 ⚔️ 🥖 🇦🇲 ⏰ 🐱 😽 💕 🇦🇱 🇦🇽 💪 🥄 🇧🇸 ♒ 🚑 🐤 👽 🇦🇸 🏺 💢 👼 🏀 🚲 👙 🎱 🔋 🇨🇺 🎹 🎵 🐜 📆 🌇 🍎 🇦🇶 💔 👾 🧬 🪕 🇯🇪 🦠 🤖 ☄️ 🌃 ♾️ ℹ️ 🤿 used in GitHub via its API...

## Contents

1. [Introduction](#introduction)
2. [Presentation](#presentation)
3. [Project's structure](#project_s_structure)
4. [How was this project developed ?](#how_was_this_project_developed)
5. [How does this project work ?](#how_does_this_project_work)
6. [How to use it ?](#how_to_use_it)
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

Coming soon...

<a name="presentation"></a>
## Presentation

Coming soon...

<a name="project_s_structure"></a>
## Project's structure

Coming soon...

<a name="how_was_this_project_developed"></a>
## How was this project developed ?

Coming soon...

<a name="how_does_this_project_work"></a>
## How does this project work ?

Coming soon...

<a name="how_to_use_it"></a>
## How to use it ?

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

<a name="from_dockerfile"></a>
#### From Dockerfile



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

<a name="from_docker_hub"></a>
#### From Docker Hub

```bash
docker pull wicken/goemojisgithubapiwebserver:latest
```

<a name="with_balena"></a>
### With Balena

<a name="where_to_use_it"></a>
## Where to use it ?

Coming soon...

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
