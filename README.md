# gowatch

gowatch will provide configurable logfile analysis for your server. It will be able to parse your logfiles and create
summaries in formats ready for delivery via E-Mail or Web.

However, this is still under development and _not_ ready for use yet.

[![Build Status](https://travis-ci.org/fxnn/gowatch.svg)](https://travis-ci.org/fxnn/gowatch)

## Installation

Just `go get` the program. The following assumes that your `GOPATH` is set and your `PATH` contains your `$GOPATH/bin` directory; if not so, please read the [Code Organization](https://golang.org/doc/code.html#Organization) chapter of the go manual.

```
$ go get github.com/fxnn/gowatch
$ gowatch
2015/04/08 19:10:44 No configuration file specified
```

## Related work

* **[logwatch](http://logwatch.sourceforge.net)** is widely used by Linux server administrators round the world, and so did
  I use it for many years. However, I find it to be not flexible enough in its configuration, and as soon as I want to
  change something, I always felt it was hard to extend and hard to change. Gowatch aims to be flexible, configurable
  and extendable.
* **[logstash](http://logstash.net)** is a log processor, that became very popular in combination with the search serer
  [elasticsearch](http://www.elasticsearch.org). Those are really great tools, especially for usage in large server
  parks. However, they need several Gigabytes of RAM and that's just far too heavy for my small tiny server. Gowatch
  aims to be a small and easy-to-be-used tool with low requirements, just as logwatch always was.

## 3rd Party Libraries

[Standing on the shoulders of giants](http://en.wikipedia.org/wiki/Standing_on_the_shoulders_of_giants), this wouldn't be what it is without:

* **[gemsi/grok](http://github.com/gemsi/grok)** is a neat Grok implementation in Go. Grok itself is a simple DRY method
  for log parsing, known from [logstash](http://logstash.net/docs/latest/filters/grok), but also there as standalone C
  implementation: see for [jordansissel/grok](https://github.com/jordansissel/grok).
* **[stretchr/testify](http://github.com/stretchr/testify)** brings assertions to Go, just the way they feel right.

...among others. Thanks a lot for your work!

## License

Licensed under MIT, see for [LICENSE](LICENSE) file.
