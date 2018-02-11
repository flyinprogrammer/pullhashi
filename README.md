# pullhashi

This is a CLI tool which will download and install the latest versions of all the [HashiCorp](https://www.hashicorp.com/)
tools (except Vagrant) to your host. By default it will install the files in your `${HOME}/bin` directory but you can
also use `-bindir` to override this.

Most shells will put `${HOME}/bin` on your `$PATH`. If it isn't there, add:

```
export PATH=${HOME}/bin:${PATH}
```

To the bottom of your `${HOME}/.profile`, and then run:

```
source ${HOME}/.profile
```

And then in theory all the tools will be accessible from your CLI.

# Usage

```
$ pullhashi -h
Usage of pullhashi:
  -arch string
    	the arch to filter packages on (default "amd64")
  -bindir string
    	download the binaries to a specific folder (default "/Users/flyinprogrammer/bin")
  -os string
    	the os to filter packages on (default "darwin")
```

Note: All the default values are dynamic, and based on your current system and user.

## How to install the latest everything?

```
$ pullhashi
2018-02-10T21:36:41-06:00 |DEBU| shasums were signed by hashicorp product=consul
2018-02-10T21:36:41-06:00 |DEBU| file sha matched shasums product=consul
2018-02-10T21:36:41-06:00 |INFO| created: /Users/flyinprogrammer/bin/consul product=consul
...
```

# Why?

Because I got tired of needing to install `gpg` & `curl` on hosts and having to add the Hashicorp GPG key, and then run
and maintain a Bash script which had to know about the os and variance in `sha` cli tooling.

And I got really tired of manually going to the websites to figure out if I was still running the latest
versions of tooling locally.

# What about Vagrant?

Vagrant is all Ruby, and their packaging and releasing of it make the [JSON](https://releases.hashicorp.com/index.json)
all different and weird. If someone wants to submit a PR I would gladly review it.

# TODO

- `docker-base` zip contains a `bin` directory, we should flatten that or drop support.
- can probably drop support for `docker-basetool` and `otto`.
