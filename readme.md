# Playing with Go and Raylib

## Quick Start
First, compile a builder app:
```console
# On Unix-like
go build -o ./scripts/b ./cmd/building

# On Windows
go build -o .\scripts\b.exe .\cmd\building
```

Then, build main application:
```console
# Unix-like
./scripts/b -v # '-v' to print commands

# Windows-like
.\scripts\b.exe -v # '-v' to print commands
```

Binary will stores in `build`
***

### Crosscompile from Linux to Windows
Install `mingw` for your distro:  
+ `mingw-w64` — Arch
+ `mingw64-gcc` — Fedora
+ `cross-i686-w64-mingw32` — Void
+ `gcc-mingw-w64` — Ubuntu

Then run builder app with `-t windows`:
```console
./scripts/b -t windows
```

### Crosscompile from Linux to Android
Install `android-ndk 26`, `android-sdk-build-tools 34` and `openjdk 17 jre`, then run builder app with `-t android`:
```console
./scripts/b -t android
```
If everything is successfully built apk can be found in the `android/build/outputs`

***

## ToDo:
### Crosscompile from linux to macOS
**Command:**
```console
# Old
CGO_ENABLED=1 \
  CC=x86_64-apple-darwin21.1-clang \
  GOOS=darwin \
  GOARCH=amd64 \
  go build -x \
  -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=10.15'" \
  -o ./build/

# New
CGO_ENABLED=1 \
  CC=aarch64-apple-darwin21.1-clang \
  GOOS=darwin \
  GOARCH=arm64 \
  go build -x \
  -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=12.0.0'" \
  -o ./build/
```
