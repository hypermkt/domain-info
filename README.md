# domain-info

A simple command line tool to lookup dns information. It lists NS, A, MX, TXT records of fqdn.

## Usage

```sh
$ domain-info FQDN
```

### Example

```sh
$ domain-info yahoo.co.jp
fqdn: yahoo.co.jp
Name Server: ns12.yahoo.co.jp.
Name Server: ns11.yahoo.co.jp.
Name Server: ns01.yahoo.co.jp.
Name Server: ns02.yahoo.co.jp.
A: 182.22.59.229
A: 183.79.135.206
MX: mx2.mail.yahoo.co.jp.
MX: mx3.mail.yahoo.co.jp.
MX: mx5.mail.yahoo.co.jp.
MX: mx1.mail.yahoo.co.jp.
TXT: v=spf1 include:spf.yahoo.co.jp ~all
```

## Installation

```sh
$ go get github.com/hypermkt/domain-info/cmd/domain-info
```

## License
[MIT](./LICENSE)

## Author
[hypermkt](https://github.com/hypermkt)
