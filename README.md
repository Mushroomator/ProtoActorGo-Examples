[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# ProtoActorGo-Examples
Small, simple programs written in Go using proto.actor actor model framework to get started with programming with Go and the actor model

## Table of Contents
- [ProtoActorGo-Examples](#protoactorgo-examples)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Examples](#examples)
    - [How-To: Install and run examples](#how-to-install-and-run-examples)
    - [List of examples](#list-of-examples)
      - [helloWorld](#helloworld)
      - [actorLifecycle](#actorlifecycle)
      - [pingPong](#pingpong)
      - [behavior](#behavior)
      - [unboundedNonDeterminism](#unboundednondeterminism)
  - [proto.actor: Documentation](#protoactor-documentation)
  - [License](#license)

## About
This repository contains several small Go programs (written using Go 1.17) which should help to get started programming within the [actor model](https://en.wikipedia.org/wiki/Actor_model) and [Go](https://golang.org/).
All examples make use of [proto.actor](https://proto.actor/), a framework for building powerful apps using the actor model.
As proto.actor as of now has not released a version 1.0 as of now (September 2021) there might be breaking changes.

I have created this repository to help me get started with the actor model and as a point of reference since the documentation provided by proto.actor is very sparse at the moment. Especially for the Go version (proto.actor is also shipped for C# and Kotlin) there is barely any documentation, so you might need to transfer some C# code to Go or just take an trial-and-error approach and use the IDE's IntelliSense.
Any documentation I have found and used can be found within the [proto.actor: Documentation](#protoactor-documentation) section.

## Examples
This section provides a short guide on how to run the example programs and includes a comprehensive list of all the examples I have written.

### How-To: Install and run examples
To run the examples you must have Go (ideally v1.17) installed and then you can download and install each of the examples using the following command
```bash
go install github.com/Mushroomator/ProtoActorGo-Examples/{example-Name}@{Version}
```
with *{example-Name}* being the Go module name of the example and *{version}* being a combination of semantic versioning and git commit hash.
To use the latest version (most recent commit on main) of an example just use *latest* for *{version}*.
For example if you want to run the latest version of the hello world example install it as follows:
```bash
go install github.com/Mushroomator/ProtoActorGo-Examples/HelloWorld@latest
```
To run the examples simply execute the created binaries within your *$GOPATH*. Make sure your *$GOPATH* is set correctly.
```bash
# run binary using $GOPATH
$GOPATH/bin/{example-name}
# if $GOPATH is on your $PATH just call the example like any other command/ binary
{example-name}
```
To run *HelloWorld* you might do it like this
```bash
# run binary using $GOPATH
$GOPATH/bin/HelloWorld
# if $GOPATH is on your $PATH just call the example like any other command/ binary
HelloWorld
```

### List of examples
#### helloWorld
Classic hello world app which will greet you with the name supplied to the "HelloActor".

See code [here](helloWorld/main.go).

#### actorLifecycle
An actor which goes through all the `lifecycles` an actor can have showing which messages are sent when and explaining what is going on in the background during each lifecycle and what are common operations to do during this lifecycle.

See code [here](actorLifecycle/main.go).

#### pingPong
Creates two actors which send "PING!" - "PONG!" messages two each other for eternity. 
Shows the importance of knowing the difference between `Tell()`, `Send()` and `Request()` methods as outlined in this [blog post by Oklahomer](https://blog.oklahome.net/2018/09/protoactor-go-messaging-protocol.html).

See code [here](pingPong/main.go).

#### behavior
Shows how one can dynamically change how an actor handles messages by using `behavior`.
The program will run in an infinite loop (stop it with `CTLR + C`)

See code [here](behavior/main.go).

#### unboundedNonDeterminism
Showcases a simple actor machine that counts up to an arbitrarily large number but is guaranteed to halt. This is something a non-deterministic Turing machine cannot due and therefore shows that the Turing model is limited in that respect in contrast to the actor model. See video of [Carl Hewitt (Creator of the Actor model) explaining this on Code BEAM SF 2020](https://www.youtube.com/watch?v=PJ4X0l2298k&list=PLDO04qnOyjC_aY9eVFEwbjYeolz3WCaUR&index=2&t=2s) starting from 42m 30s.

See code [here](unboundedNonDeterminism/main.go).

## proto.actor: Documentation
There is hardly any documentation available for Go, but there are a few helpful ressources for C# and the basic concepts:
- [proto.actor website](https://proto.actor/): This is the website for the proto.actor project
- [proto.actor Github repository for Go](https://github.com/AsynkronIT/protoactor-go): Contains all the code and some examples
- [proto.actor Bootcamp](https://proto.actor/docs/bootcamp/): Best starting point; explains basic concepts and introduces you step-by-step to the components of the library; it is written for C# only though
- [protoactro-go 101: Blog post by Oklahomer](https://blog.oklahome.net/2018/07/protoactor-go-introduction.html): First of a series of very inforamtive blog posts written by Oklahomer which introduces the basics of proto.actor when using Go and also provides links to example code and further ressources.

## License
Copyright (c) 2021 Thomas Pilz

This work is licensed under the [Apache 2.0 license](LICENSE). You may not use the content of this repository except in compliance with the license. You may obtain a copy of the license at [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0).