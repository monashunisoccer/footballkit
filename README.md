# footballkit

Go module that will render images of football kits given an English description.

NOTE:

This module is mainly designed to support local (small) club teams and is used by
our web page for showing possible strip clashes. It is not particularly accurate
if attempting to render 'famous' club strips, though I have given a few examples
here of our nearest approximation.

### Monash University Soccer Club (pre 2018)

`RenderImage("body stripes skyblue white shorts navy socks navy")`

![kit](example-output/monashunisoccer.png)

### Manchester United

`RenderImage("body red shorts white socks black")`

![kit](example-output/manutd.png)

### Aston Villa

`RenderImage("body claret shorts white socks light blue")`

![kit](example-output/astonvilla.png)

### Celtic

`RenderImage("hoops green white shorts white socks green")`

![kit](example-output/celtic.png)

### Peru

`RenderImage("leftsash white red shorts white socks red")`

![kit](example-output/peru.png)

### Croatia

`RenderImage("checks red white shorts white socks blue")`

![kit](example-output/croatia.png)


## Build

Install the go-bindata utility

```shell
go install github.com/go-bindata/go-bindata/...
```

then if *any* of the images (in `data/`) have been modified, regenerate the static image data file.

```shell
~/go/bin/go-bindata -pkg footballkit data/
```
