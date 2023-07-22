<div align="center">

&nbsp;
<h1>onelog</h1>
<p><i>A unified logging interface for Go. The library is currently still a work in progress.</i></p>

&nbsp;

[![codecov](https://codecov.io/gh/nikoksr/onelog/branch/main/graph/badge.svg?token=9KTRRRWM5A)](https://codecov.io/gh/nikoksr/onelog)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikoksr/onelog)](https://goreportcard.com/report/github.com/nikoksr/onelog)
[![Maintainability](https://api.codeclimate.com/v1/badges/8d58f3077a2b6ee2ac57/maintainability)](https://codeclimate.com/github/nikoksr/onelog/maintainability)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/nikoksr/onelog)
</div>

&nbsp;
## About <a id="about"></a>

> This is a work in progress.

_onelog_ is a general-purpose logging interface heavily inspired by the [zerolog](https://github.com/rs/zerolog) API. It is designed to provide a user-friendly API for diverse logging requirements. This package supports a wide range of data types and log levels, creating flexibility for various use cases.

_onelog_ includes adapters for several commonly used loggers, enabling easy integration and compatibility with existing logging methodologies. It reduces the friction associated with logging setup and promotes consistency in logging across different parts of a project or across different projects.

Personally, I plan on using this library in my personal projects [Notify](https://github.com/nikoksr/notify) and [doppler-go](https://github.com/nikoksr/doppler-go), to provide the users of both projects with a unified logging interface without having to force them to use a specific logging library.

## Install <a id="install"></a>

```sh
go get -u github.com/nikoksr/onelog
```

## Example usage <a id="usage"></a>

```go
package main

import (
    "go.uber.org/zap"
    "golang.org/x/exp/slog"

    "github.com/nikoksr/onelog"
    slogadapter "github.com/nikoksr/onelog/adapter/slog"
    zapadapter "github.com/nikoksr/onelog/adapter/zap"
)

type superheroTracker struct {
    logger onelog.Logger
}

func main() {
    // Let's use zap's development logger as our superhero event logger
    logger, _ := zap.NewDevelopment()

    heroes := &superheroTracker{
        logger: zapadapter.NewAdapter(logger),
    }

    // Now let's log a superhero event
    heroes.logger.Info().Msg("Superman spotted in New York!")

    // Or perhaps we'd rather use slog for logging our superhero sightings
    heroes.logger = slogadapter.NewAdapter(slog.Default())

    // And now we can log another sighting
    heroes.logger.Info().Msg("Wonder Woman seen flying over Paris!")
}
```

## Contributing <a id="contributing"></a>

Contributions of all kinds are very welcome! Feel free to check
our [open issues](https://github.com/nikoksr/onelog/issues). Please also take a look at
the [contribution guidelines](https://github.com/nikoksr/onelog/blob/main/CONTRIBUTING.md).

## Show your support <a id="support"></a>

Please give a ⭐️ if you like this project! This helps us to get more visibility and helps other people to find this
project.
