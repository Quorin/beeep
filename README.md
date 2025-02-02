## beeep
[![TravisCI Build Status](https://travis-ci.org/gen2brain/beeep.svg?branch=master)](https://travis-ci.org/gen2brain/beeep) 
[![AppVeyor Build Status](https://ci.appveyor.com/api/projects/status/4u7avrhsdxua2c9b?svg=true)](https://ci.appveyor.com/project/gen2brain/beeep)
[![GoDoc](https://godoc.org/github.com/gen2brain/beeep?status.svg)](https://godoc.org/github.com/gen2brain/beeep) 
[![Go Report Card](https://goreportcard.com/badge/github.com/gen2brain/beeep?branch=master)](https://goreportcard.com/report/github.com/gen2brain/beeep) 
<!--[![Go Cover](http://gocover.io/_badge/github.com/gen2brain/beeep)](http://gocover.io/github.com/gen2brain/beeep)-->

`beeep` provides a cross-platform library for sending desktop notifications, alerts and beeps.

### Installation

    go get -u github.com/Quorin/beeep

### Examples

```go
err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
if err != nil {
    panic(err)
}
```

```go
err := beeep.Notify("App name", "Title", "Message body", "assets/information.png")
if err != nil {
    panic(err)
}
```

```go
err := beeep.Alert("App name", "Title", "Message body", "assets/warning.png")
if err != nil {
    panic(err)
}
```


## macOS

For icons to show up when using Alert() or Notify(), you will need to bundle your application
with a app icon.

## More

For cross-platform dialogs and input boxes see [dlgs](https://github.com/gen2brain/dlgs).
