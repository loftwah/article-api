---
layout: test
title: Article API - Dean Lofts
---

![Nine Publishing](https://repository-images.githubusercontent.com/504058214/88657ffb-1213-4d40-b0b8-be00dbed134f)

## Nine Publishing - Backend / DevOps Articles API Test

[Test](https://ffxblue.github.io/interview-tests/test/article-api/)

This is my submission for the technical test for Nine Publishing.

I'm going to build this in `Golang`, which is actually a bad idea given it isn't a language I'm familiar with. I will also write the instructions to go witih it.

## Usage

[Golang](https://go.dev/doc/install) | [Golang official Docker image](https://hub.docker.com/_/golang) | [Managing Go Installations](https://go.dev/doc/manage-install)

`Clone` my repository and change to the directory it has cloned to. I'm going to assume you're set up and able to run `Git` commands.

Download and install `Golang`

```bash
curl -O -L -C - https://go.dev/dl/go1.18.3.linux-amd64.tar.gz # Check the latest version here: https://golang.org/dl/
sudo tar -C /usr/local -xvzf go1.18.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
source $HOME/.profile
go version
```

Download and install `Docker` (I'm using `Ubuntu` so any instructions will be for that)

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
sudo service docker start
sudo systemctl enable docker
sudo usermod -aG docker <username>
```

And `Docker-Compose`

```bash
sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose version
```

You will need the `Golang` docker image to get started

```bash
docker pull golang
```

Now it is time to clone the repository and build the image

```bash
git clone git@github.com:loftwah/article-api.git
cd article-api
```

This part can be skipped because it is how I built it, you are cloning a version where this has already been completed

```bash
export GOFLAGS=-mod=vendor # might not be needed
export GO111MODULE=on # might not be needed
go mod init github.com/loftwah/article-api
```

(Optional) Use `Docker` and build the `Dockerfile`

```bash
docker build -t loftwah/article-api:latest .
```

(Optional) Run the `Docker` image. I use a different port for each test so I can run multiple tests at once.

```bash
docker run -p 8001:8000 loftwah/article-api
```

## Solution

[mux](https://github.com/gorilla/mux)

Initialize the project

```bash
go mod init github.com/loftwah/article-api
```

Install mux router

```bash
go get -u github.com/gorilla/mux
```

Build source code

```bash
go build
```

Run the `API`

```bash
./article-api
```

### Meet requirement 1

The first endpoint, POST `/articles` should handle the receipt of some article data in json format, and store it within the service.

Create a new article

```bash
curl -H "Content-Type: application/json" -X POST -d '{"id":"1","title":"Article Three","date":"2021-09-17","body":"This is the body of article three","tags":["tag1","tag2","tag3"]}' http://127.0.0.1:8000/articles
```

You can check that your article has been added with the following command

```bash
curl http://127.0.0.1:8000/articles
```

### Meet requirement 2

The second endpoint GET `/articles/{id}` should return the JSON representation of the article.

Use `cURL` to show that your article was added. My `ID` is `19727887` which may be different to yours.

```bash
curl http://127.0.0.1:8000/articles/19727887
```

### Meet requirement 3

The final endpoint, GET `/tags/{tagName}/{date}` will return the list of articles that have that tag name on the given date and some summary data about that tag for that day.

The GET `/tags/{tagName}/{date}` endpoint should produce the following JSON. Note that the actual url would look like /tags/health/20160922.

```json
{
  "tag": "health",
  "count": 17,
  "articles": ["1", "7"],
  "related_tags": ["science", "fitness"]
}
```

The `related_tags` field contains a list of tags that are on the articles that the current tag is on for the same day. It should not contain duplicates.

The `count` field shows the number of tags for the tag for that day.

The `articles` field contains a list of ids for the last 10 articles entered for that day.

If you don't have data, run the following and change it up if you want to add more.

```bash
curl -H "Content-Type: application/json" -X POST -d '{"id":"1","title":"Article Four","date":"2021-09-17","body":"This is the body of article four","tags":["tag1","tag2","tag3"]}' http://127.0.0.1:8000/articles
```

Test the `API` endpoint

```bash
curl http://127.0.0.1:8000/tags/tag1/2021-09-17
```

## Deliverables

1. Source code for the solution described above

This has been delivered within this GitHub [repository](https://github.com/loftwah/article-api)

2. Setup/installation instructions

Documentation has been written to show how this works, and how to use it

3. A quick (1-2 page) description of your solution, outlining anything of interest about the code you have produced. This could be anything from why you chose the language and or libraries, why you structured the project the way that you did, why you chose a particular error handling strategy, how you approached testing etc

I have put together a small and quick `API` to demonstrate my ability to code and understanding of how an `API` works. I chose the language because it matches the stack that I could be working with if I end up being hired for a job, and I chose the `mux` library because I found that it would have been too much effort to attempt to build this in `GraphQL` when I wasn't familiar with it so much already and it looked like a good way to build the `API` that was described in the test document and I knew I'd be able to build what I needed with it even though I'd never used it before.

I haven't actually accounted for error handling or testing as it wasn't mentioned as part of the requirements. If I had to add it in, I'd probably use a library like `go-test-deep` to test the error handling and testing, although I don't know certainly that this would work.

The code consists of some common libraries I found while doing some tutorials before building this, and some code I wrote myself. It exists in the `main.go` file and starts by declaring the `Article` structure, which is made up of an `id`, `title`, `date`, `body` and `tags` field. I actually zoned out and lost track of what I was doing and built an entire `CRUD` style `API` and then worked out that I'd met the requirements later on.

I had to revist my `getArticleByTagAndDate` function because I made some assumptions at first that didn't actually match the requirements I needed to meet. With the small amount of data I'd been working with it was hard to tell if I was getting back what I should be, so I wrote a script (I didn't include it here, just some hacky bash code) to test the function and see if it was working as I expected.

I hardcoded the data into the function to make it easy to work with and I also included some routes that weren't part of the requirements. The `server` will run on port `8000`. I have also included a `Dockerfile` to make it easy to build the image and run it.

4. A list of assumptions that you've made while putting this together. We've only given you a very loose spec, so you'll probably need to fill in some blanks while you are working. If you note down the assumptions, for us, then we will be able review the code within the context of those assumptions

I assumed that this is an `API` to demonstrate that I understand how a basic `CRUD` application works, that the user has a reasonable understanding of an operating system that is able to run this application, and that the user has a basic understanding of the `Go` language.

5. [Optional] Tell us what you thought of the test and how long it took you to complete

It took me about an hour to complete this, but I had to stop and do other things before I was able to come back and complete it.

6. [Optional] Tell us what else you would have added to the code if you had more time

I would have built it as a `GraphQL API` like I said I would, but I already chose to use a programming language that I've never really used before and thought that was enough of an extra challenge, given it was a simple application. I also would attach it to a database rather than just storing things in memory.

**Note:** We prefer that you send us a link to a (public) repository. If you send an attachment via a zip file with your source code, please be aware that your email may get blocked. You will receive a confirmation email for your submission.

[Link to repository](https://github.com/loftwah/article-api)
