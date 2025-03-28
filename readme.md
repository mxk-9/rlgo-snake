# Playing with Go and Raylib

## ToDo:
+ [X] Finish working build of snake game
+ [ ] Crosscompile to windows:
```
CGO_ENABLED=1 \
  CC=x86_64-w64-mingw32-gcc \
  GOOS=windows \
  GOARCH=amd64 \
  go build \
  -x -ldflags "-s -w" \
  -o ./build/

```
+ [ ] Crosscompile to macOS:
```
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

+ [ ] To [android](https://github.com/gen2brain/raylib-go/tree/master/examples/others/android/example)
     
+ [ ] Crossplatform:
    - [X] Windows
    - [ ] Android
+ [X] Add 7z support(for `w64devkit`)
