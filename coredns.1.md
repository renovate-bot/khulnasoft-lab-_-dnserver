## DNServer

*dnserver* - pluggable DNS nameserver optimized for service discovery and flexibility.

## Synopsis

*dnserver* **[-conf FILE]** **[-dns.port PORT}** **[OPTION]**...

## Description

DNServer is a DNS server that chains plugins. Each plugin handles a DNS feature, like rewriting
queries, kubernetes service discovery or just exporting metrics. There are many other plugins,
each described on <https://dnsserver.khulnasoft.com/plugins> and their respective manual pages. Plugins not
bundled by default in DNServer are listed on <https://dnsserver.khulnasoft.com/explugins>.

When started without options DNServer will look for a file named `Corefile` in the current
directory, if found, it will parse its contents and start up accordingly. If no `Corefile` is found
it will start with the *whoami* plugin (dnserver-whoami(7)) and start listening on port 53 (unless
overridden with `-dns.port`).

Available options:

**-conf** **FILE**
: specify Corefile to load, if not given DNServer will look for a `Corefile` in the current
  directory.

**-dns.port** **PORT** or **-p** **PORT**
: override default port (53) to listen on.

**-pidfile** **FILE**
: write PID to **FILE**.

**-plugins**
: list all plugins and quit.

**-quiet**
: don't print any version and port information on startup.

**-version**
: show version and quit.

## Authors

DNServer Authors.

## Copyright

Apache License 2.0

## See Also

Corefile(5) @@PLUGINS@@.
