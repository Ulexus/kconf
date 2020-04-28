<p align="center">
  <h1 align="center">kconf</h1>
  <p align="center">An opinionated command line tool for managing multiple kubeconfigs.</p>
  <p align="center">
    <a href="https://github.com/particledecay/kconf/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/particledecay/kconf"></a>
    <a href="https://github.com/particledecay/kconf/actions?query=workflow%3Atests"><img alt="GitHub Workflow Status" src="https://github.com/particledecay/kconf/workflows/tests/badge.svg"></a>
    <a href="https://codeclimate.com/github/particledecay/kconf/maintainability"><img src="https://api.codeclimate.com/v1/badges/dd1904e8f1f515bad0b5/maintainability" /></a>
    <a href="https://codeclimate.com/github/particledecay/kconf/test_coverage"><img src="https://api.codeclimate.com/v1/badges/dd1904e8f1f515bad0b5/test_coverage" /></a>
  </p>
</p>



## Description

kconf works by storing all kubeconfig information in a single file (`$HOME/.kube/config`). This file is looked at by default when using `kubectl`.

## Usage

##### Add in a new kubeconfig file:

```sh
kconf add /path/to/kubeconfig.conf
```

or

```sh
kconf add /path/to/kubeconfig.conf --context-name=myContext
```

##### Remove an existing kubeconfig:

```sh
kconf rm myContext
```

##### List all saved contexts in the kubeconfig:

```sh
kconf ls
```

##### View and print a single context's kubeconfig (you can pipe or export to a file):

```sh
kconf view myContext
```

##### Switch to an existing context:

```sh
kconf use myContext
```

##### Set a preferred namespace

```sh
kconf use myContext -n kube-system
```

## Why?

I was previously managing my kubeconfigs using the `$KUBECONFIG` environment variable. However, in order to automate this process, you have to do something like this in your rc files:

```bash
KUBECONFIG=$(find $HOME/.kube -type f -name '*.conf' 2> /dev/null | sed ':a;N;$!ba;s/\n/:/g')
```

... that gets you a `$KUBECONFIG` variable with all your kubeconfigs separated by colons. The problem is that if you're frequently working with new/modified kubeconfigs, you'd have to trigger this command each time something changed.

With the `kconf` command, there's no need for `$KUBECONFIG` since `kubectl` already looks at `$HOME/.kube/config` by default. Additionally, as soon as you have a new kubeconfig, you can `add` it pretty easily and quickly.

## Known Issues

Check out the [Issues](https://github.com/particledecay/kconf/issues) section or specifically [issues created by me](https://github.com/particledecay/kconf/issues?q=is:issue+is:open+sort:updated-desc+author:particledecay)
