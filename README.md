Terraform Provider for Cohesity
=========================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by [Cohesity](https://www.cohesity.com/)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.20+
-	[Go](https://golang.org/doc/install) 1.12.6+ (to build the provider plugin)
-   [Cohesity](https://www.cohesity.com/) DataPlatform 6.4+

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-cohesity`

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-cohesity $GOPATH/src/github.com/terraform-providers/terraform-provider-cohesity
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-cohesity
$ make build
```

The provider binary can be found in `$GOPATH/bin` directory.

Using the provider
----------------------
The Cohesity provider documentation can be found on [provider's website](https://www.terraform.io/docs/providers/cohesity/index.html)

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12.6+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-cohesity
...
```
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```
In order to run a specific acceptance test, change the variables in acceptance_test_variables.go to suit your environment and run `make testacc`

*Note:* Set the environment variables for configuring the Cohesity provider and sensitive resource arguments not set in the acceptance_test_variables.go file

```sh
$ make testacc TEST=./cohesity TESTARGS='-run=TestAccVirtualEditionCluster_basic'
```