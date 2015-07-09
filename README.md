# go-choises

[![screencast](https://raw.githubusercontent.com/geeks-dev/go-choices/images/screencast.gif)](https://raw.githubusercontent.com/geeks-dev/go-choices/images/screencast.gif "screencast")


## Description

It is a package that requires a simple selection.

## KeyBind

| Key            | Description           |
|:--------------:|:---------------------:|
| Enter          | Decision              |
| Esc            | Exit                  |
| Key Arrow Down | Move to choices below |
| Key Arrow Up   | Move to choices above |
| Key Arrow Right| Page Next             |
| Key Arrow Left | Page Prev             |


## Installation

To install, use `go get`:

```bash
$ go get github.com/geeks-dev/go-choices
```

## Usage

```go
gamemode := []string{
	"HUMAN",
	"DEVIL HUNTER",
	"SON OF SPARDA",
	"HEAVEN OR HELL",
	"LEGENDARY DARK KNIGHT",
	"DANTE MUST DIE",
	"HELL AND HELL",
	"BLOODY PALACE",
}
mode, answered := choices.Ask(
	gamemode,
	"Devil May Cry",
)

if answered {
	fmt.Println(gamemode[mode])
}
```

## Author

[geeks-dev](https://github.com/geeks-dev)
