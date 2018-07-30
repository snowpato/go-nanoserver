# NanoServer v0.1
Tested with **Golang v1.10**

# About
This script uses the **net/http** library to run a nano web server in **localhost:8081**. The server will consider the app folder as the root path.

## But why?

Well, there are many nano servers built with many addons (XAMPP, WAMPP, etc...) and i wanted something lighter to run only static HTML sites for developing.

## Ok, give me the commands

### The short way:
>`cd server/src && go run server.go`

### The long way:
 1. Compile the source code with the build script. This will create executables for Windows, Mac & Linux for 32 & 64 bits (you can modify the build script to accommodate your needs).
>`cd server && ./build`
2. Depending on your architecture, launch the executable. Remember to enable the firewall, if applies.

## How can I change the port?

Modify the `server.go` script. There is a section in which is declared the port. Simple, huh? Obviously, you need to recompile the script if you are using an executable.

# Licenses
You can use the code as you want. Just keep a mention. The example page uses jQuery + Bootstrap.

# Mentions
This app was created based in this 2 pages:
> [Creating A Simple Web Server With Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/) & [How To Build Go Executables for Multiple Platforms](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)