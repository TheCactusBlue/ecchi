# Ecchi

## Getting started


```
go get -u github.com/thecactusblue/ecchi
```

## Quick Start

```go
r := ecchi.NewRouter()

type Message struct {
    Name    string
    Content string
}

r.Get("/echo", func(ctx *ecchi.Ctx) *ecchi.Ctx {
    var m Message
    ctx.ReadJSON(&m)

    return ctx.JSON(&Message{
        Name: "Echo",
        Content: m.Content,
    })
})
```

**Q.** What is this?

**A.** This is a "extension" of [errchi](https://gitlab.com/shihoya-inc/errchi/-/tree/master) by diamondburned. Since it no longer deals with just errors, I decided to name it eChi at first, short for extended chi, until I managed to come up with this pun. Major features involve auto JSON marshaling.
