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
yahoo.co.jp.	223	IN	NS	ns11.yahoo.co.jp.
yahoo.co.jp.	223	IN	NS	ns02.yahoo.co.jp.
yahoo.co.jp.	223	IN	NS	ns01.yahoo.co.jp.
yahoo.co.jp.	223	IN	NS	ns12.yahoo.co.jp.
yahoo.co.jp.	33	IN	A	182.22.59.229
yahoo.co.jp.	33	IN	A	183.79.135.206
yahoo.co.jp.	804	IN	MX	10 mx5.mail.yahoo.co.jp.
yahoo.co.jp.	804	IN	MX	10 mx1.mail.yahoo.co.jp.
yahoo.co.jp.	804	IN	MX	10 mx2.mail.yahoo.co.jp.
yahoo.co.jp.	804	IN	MX	10 mx3.mail.yahoo.co.jp.
yahoo.co.jp.	816	IN	TXT	"v=spf1 include:spf.yahoo.co.jp ~all"
```

## Installation

```sh
$ go get github.com/hypermkt/domain-info/cmd/domain-info
```

## License
[MIT](./LICENSE)

## Author
[hypermkt](https://github.com/hypermkt)
