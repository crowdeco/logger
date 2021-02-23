# Crowde Logger

## Requirement

This package need environment variable `APP_DEBUG` with **boolean** value (`true` or `false`)

## Install

`go get github.com/crowdeco/logger`

## Usage

```go
logger := logger.NewLogger()


logger.Trace("Trace Message")
logger.Debug("Debug Message")
logger.Warning("Warning Message")
logger.Error("Error Message")
logger.Fatal("Fatal Message")
logger.Panic("Panic Message")
```
