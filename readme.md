# Playing with Go and Raylib

## ToDo:
+ [ ] Finish working build of snake game
+ [ ] Crossplatform:
    - [X] Windows
    - [ ] Android

## Dependencies
+ `go`
+ (optional) `mingw` to build engine for Windows
+ (for linux) [Raylib dependecies](https://github.com/raysan5/raylib/wiki/Working-on-GNU-Linux#dependencies)

## Building from source (development)
The easiest way is to compile a program for project building and run it after
editing the game's source code:

```console
# Linux
$ go build -o mbs ./scripts/my_build_system.go
# Windows
$ go build -o mbs.exe .\scripts\my_build_system.go
```

## MBS:building
### For Linux
```console
$ ./mbs
```

### For Windows
You will need `raylib.dll` and by default `mbs` compiles raylib from source
code, but you can use the already built dynamic library(compiled with `mvsc`) by
simply enabling `-prebuilt-lib` key. Both methods in the end place file
`raylib.dll` into `build` directory. Also, if `raylib.dll` exists in `build`
directory, `mbs` will skip this step by going straight to rebuilding the
project.

By the way, if you want to build raylib from source code, you should download
[archive](https://github.com/raysan5/raylib/archive/refs/tags/5.5.zip),
unpack `src` folder into `third_party` and copy `MakeFile` from `third_party` to
`third_party/src`. This is custom makefile which allows you to build engine as
 dynamic library.


```console
# Linux(crosscompile w/ building engine from source code)
$ ./mbs -target windows

# Linux (crosscompile w/ pre-built engine)
$ ./mbs -target windows -prebuilt-lib

# Windows (w/ building from src)
$ mbs.exe
# Windows (w/ pre-built)
$ mbs.exe -prebuilt-lib
```

## MBS:other projects
`cmd` directory also stores other projects for testing my shitcode, you can
build them(why not).

Example:
```console
$ ./mbs -list
Available projects:
cgo_check
checkDir
remake
rltest

$ ./mbs -name cgo_check
```
