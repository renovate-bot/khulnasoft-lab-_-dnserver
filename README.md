[![DNServer](https://dnserver.khulnasoft.com/images/DNServer_Colour_Horizontal.png)](https://dnserver.khulnasoft.com)

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/khulnasoft-lab/dnserver)
![CodeQL](https://github.com/khulnasoft-lab/dnserver/actions/workflows/codeql-analysis.yml/badge.svg)
![Go Tests](https://github.com/khulnasoft-lab/dnserver/actions/workflows/go.test.yml/badge.svg)
[![CircleCI](https://circleci.com/gh/khulnasoft-lab/dnserver.svg?style=shield)](https://circleci.com/gh/khulnasoft-lab/dnserver)
[![Code Coverage](https://img.shields.io/codecov/c/github/khulnasoft-lab/dnserver/master.svg)](https://codecov.io/github/khulnasoft-lab/dnserver?branch=master)
[![Docker Pulls](https://img.shields.io/docker/pulls/khulnasoft-lab/dnserver.svg)](https://hub.docker.com/r/khulnasoft-lab/dnserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft-lab/dnserver)](https://goreportcard.com/report/khulnasoft-lab/dnserver)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1250/badge)](https://bestpractices.coreinfrastructure.org/projects/1250)

DNServer is a DNS server/forwarder, written in Go, that chains [plugins](https://dnserver.khulnasoft.com/plugins).
Each plugin performs a (DNS) function.

DNServer is a [Cloud Native Computing Foundation](https://cncf.io) graduated project.

DNServer is a fast and flexible DNS server. The key word here is *flexible*: with DNServer you
are able to do what you want with your DNS data by utilizing plugins. If some functionality is not
provided out of the box you can add it by [writing a plugin](https://dnserver.khulnasoft.com/explugins).

DNServer can listen for DNS requests coming in over:
* UDP/TCP (go'old DNS).
* TLS - DoT ([RFC 7858](https://tools.ietf.org/html/rfc7858)).
* DNS over HTTP/2 - DoH ([RFC 8484](https://tools.ietf.org/html/rfc8484)).
* DNS over QUIC - DoQ ([RFC 9250](https://tools.ietf.org/html/rfc9250)). 
* [gRPC](https://grpc.io) (not a standard).

Currently DNServer is able to:

* Serve zone data from a file; both DNSSEC (NSEC only) and DNS are supported (*file* and *auto*).
* Retrieve zone data from primaries, i.e., act as a secondary server (AXFR only) (*secondary*).
* Sign zone data on-the-fly (*dnssec*).
* Load balancing of responses (*loadbalance*).
* Allow for zone transfers, i.e., act as a primary server (*file* + *transfer*).
* Automatically load zone files from disk (*auto*).
* Caching of DNS responses (*cache*).
* Use etcd as a backend (replacing [SkyDNS](https://github.com/skynetservices/skydns)) (*etcd*).
* Use k8s (kubernetes) as a backend (*kubernetes*).
* Serve as a proxy to forward queries to some other (recursive) nameserver (*forward*).
* Provide metrics (by using Prometheus) (*prometheus*).
* Provide query (*log*) and error (*errors*) logging.
* Integrate with cloud providers (*route53*).
* Support the CH class: `version.bind` and friends (*chaos*).
* Support the RFC 5001 DNS name server identifier (NSID) option (*nsid*).
* Profiling support (*pprof*).
* Rewrite queries (qtype, qclass and qname) (*rewrite* and *template*).
* Block ANY queries (*any*).
* Provide DNS64 IPv6 Translation (*dns64*).

And more. Each of the plugins is documented. See [dnserver.khulnasoft.com/plugins](https://dnserver.khulnasoft.com/plugins)
for all in-tree plugins, and [dnserver.khulnasoft.com/explugins](https://dnserver.khulnasoft.com/explugins) for all
out-of-tree plugins.

## Compilation from Source

To compile DNServer, we assume you have a working Go setup. See various tutorials if you don’t have
that already configured.

First, make sure your golang version is 1.20 or higher as `go mod` support and other api is needed.
See [here](https://github.com/golang/go/wiki/Modules) for `go mod` details.
Then, check out the project and run `make` to compile the binary:

~~~
$ git clone https://github.com/khulnasoft-lab/dnserver
$ cd dnserver
$ make
~~~

This should yield a `dnserver` binary.

## Compilation with Docker

DNServer requires Go to compile. However, if you already have docker installed and prefer not to
setup a Go environment, you could build DNServer easily:

```
$ docker run --rm -i -t -v $PWD:/v -w /v golang:1.21 make
```

The above command alone will have `dnserver` binary generated.

## Examples

When starting DNServer without any configuration, it loads the
[*whoami*](https://dnserver.khulnasoft.com/plugins/whoami) and [*log*](https://dnserver.khulnasoft.com/plugins/log) plugins
and starts listening on port 53 (override with `-dns.port`), it should show the following:

~~~ txt
.:53
DNServer-1.6.6
linux/amd64, go1.16.10, aa8c32
~~~

The following could be used to query the DNServer server that is running now:

~~~ txt
dig @127.0.0.1 -p 53 www.example.com
~~~

Any query sent to port 53 should return some information; your sending address, port and protocol
used. The query should also be logged to standard output.

The configuration of DNServer is done through a file named `Corefile`. When DNServer starts, it will
look for the `Corefile` from the current working directory. A `Corefile` for DNServer server that listens
on port `53` and enables `whoami` plugin is:

~~~ corefile
.:53 {
    whoami
}
~~~

Sometimes port number 53 is occupied by system processes. In that case you can start the DNServer server
while modifying the `Corefile` as given below so that the DNServer server starts on port 1053.

~~~ corefile
.:1053 {
    whoami
}
~~~

If you have a `Corefile` without a port number specified it will, by default, use port 53, but you can
override the port with the `-dns.port` flag: `dnserver -dns.port 1053`, runs the server on port 1053.

You may import other text files into the `Corefile` using the _import_ directive.  You can use globs to match multiple
files with a single _import_ directive.

~~~ txt
.:53 {
    import example1.txt
}
import example2.txt
~~~

You can use environment variables in the `Corefile` with `{$VARIABLE}`.  Note that each environment variable is inserted
into the `Corefile` as a single token. For example, an environment variable with a space in it will be treated as a single
token, not as two separate tokens.

~~~ txt
.:53 {
    {$ENV_VAR}
}
~~~

A Corefile for a DNServer server that forward any queries to an upstream DNS (e.g., `8.8.8.8`) is as follows:

~~~ corefile
.:53 {
    forward . 8.8.8.8:53
    log
}
~~~

Start DNServer and then query on that port (53). The query should be forwarded to 8.8.8.8 and the
response will be returned. Each query should also show up in the log which is printed on standard
output.

To serve the (NSEC) DNSSEC-signed `example.org` on port 1053, with errors and logging sent to standard
output. Allow zone transfers to everybody, but specifically mention 1 IP address so that DNServer can
send notifies to it.

~~~ txt
example.org:1053 {
    file /var/lib/dnserver/example.org.signed
    transfer {
        to * 2001:500:8f::53
    }
    errors
    log
}
~~~

Serve `example.org` on port 1053, but forward everything that does *not* match `example.org` to a
recursive nameserver *and* rewrite ANY queries to HINFO.

~~~ txt
example.org:1053 {
    file /var/lib/dnserver/example.org.signed
    transfer {
        to * 2001:500:8f::53
    }
    errors
    log
}

. {
    any
    forward . 8.8.8.8:53
    errors
    log
}
~~~

IP addresses are also allowed. They are automatically converted to reverse zones:

~~~ corefile
10.0.0.0/24 {
    whoami
}
~~~
Means you are authoritative for `0.0.10.in-addr.arpa.`.

This also works for IPv6 addresses. If for some reason you want to serve a zone named `10.0.0.0/24`
add the closing dot: `10.0.0.0/24.` as this also stops the conversion.

This even works for CIDR (See RFC 1518 and 1519) addressing, i.e. `10.0.0.0/25`, DNServer will then
check if the `in-addr` request falls in the correct range.

Listening on TLS (DoT) and for gRPC? Use:

~~~ corefile
tls://example.org grpc://example.org {
    whoami
}
~~~

Similarly, for QUIC (DoQ):

~~~ corefile
quic://example.org {
    whoami
    tls mycert mykey
}
~~~

And for DNS over HTTP/2 (DoH) use:

~~~ corefile
https://example.org {
    whoami
    tls mycert mykey
}
~~~
in this setup, the DNServer will be responsible for TLS termination

you can also start DNS server serving DoH without TLS termination (plain HTTP), but beware that in such scenario there has to be some kind
of TLS termination proxy before DNServer instance, which forwards DNS requests otherwise clients will not be able to communicate via DoH with the server
~~~ corefile
https://example.org {
    whoami
}
~~~

Specifying ports works in the same way:

~~~ txt
grpc://example.org:1443 https://example.org:1444 {
    # ...
}
~~~

When no transport protocol is specified the default `dns://` is assumed.

## Community

We're most active on Github (and Slack):

- Github: <https://github.com/khulnasoft-lab/dnserver>
- Slack: #dnserver on <https://slack.cncf.io>

More resources can be found:

- Website: <https://dnserver.khulnasoft.com>
- Blog: <https://dnserver.khulnasoft.com/blog/>
- Twitter: [@khulnasoft](https://twitter.com/khulnasoft)

## Contribution guidelines

If you want to contribute to DNServer, be sure to review the [contribution
guidelines](./.github/CONTRIBUTING.md).

## Deprecation Policy

When there is a backwards incompatible change in DNServer the following process is followed:

*  Release x.y.z: Announce that in the next release we will make backward incompatible changes.
*  Release x.y+1.0: Increase the minor version and set the patch version to 0. Make the changes,
   but allow the old configuration to be parsed. I.e. DNServer will start from an unchanged
   Corefile.
*  Release x.y+1.1: Increase the patch version to 1. Remove the lenient parsing, so DNServer will
   not start if those features are still used.

E.g. 1.3.1 announce a change. 1.4.0 a new release with the change but backward compatible config.
And finally 1.4.1 that removes the config workarounds.

### Reporting security vulnerabilities

If you find a security vulnerability or any security related issues, please DO NOT file a public
issue, instead send your report privately to `security@dnserver.khulnasoft.com`. Security reports are greatly
appreciated and we will publicly thank you for it.

Please consult [security vulnerability disclosures and security fix and release process
document](https://github.com/khulnasoft-lab/dnserver/blob/master/.github/SECURITY.md)
