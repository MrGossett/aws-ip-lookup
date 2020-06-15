# aws-ip-lookup

## Overview

This tool downloads and parses the list of [AWS IP address ranges](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html), and then searches through those ranges to find one that contains a given IPv4 or IPv6 address.

## Install

```
go get github.com/MrGossett/aws-ip-lookup
```

## Example

```
$ dig www.amazon.com A +short
tp.47cf2c8c9-frontier.amazon.com.
d3ag4hukkh62yn.cloudfront.net.
13.249.104.206

$ aws-ip-lookup -target 13.249.104.206
Target:  13.249.104.206
Prefix:  13.249.0.0/16
Region:  GLOBAL
NBG:     GLOBAL
Service: AMAZON

```

## License

MIT.
