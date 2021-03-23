# Cohesity Terraform Provider

![](docs/assets/images/terraform.png)

## Overview

This project provides a Terraform Provider for interacting with the [Cohesity DataPlatform](https://www.cohesity.com/products/data-platform). It includes Terraform modules useful for automating common tasks and orchestrating workflows in your environment.

## Table of contents :scroll:

 - [Requirements](#get-startedt)
 - [Building the Provider](#building)
 - [Documentation](#doc)
 - [Developing the Provider](#developing-the-provider)
 - [How can you contribute](#contribute)
 - [Suggestions and Feedback](#suggest)
 

## <a name="get-started"></a> Let's get started :hammer_and_pick:

The pre-requisites for using the Cohesity Terraform Provider are as below:


-	[Terraform](https://www.terraform.io/downloads.html) 0.12.20+
-	[Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
-   [Cohesity](https://www.cohesity.com/) DataPlatform 6.4+

## <a name="building"></a> Building the Provider :gear:

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


## <a name="doc"></a> Documentation :book:

* [Documentation for Cohesity Terraform Provider](https://github.com/cohesity/terraform-provider-cohesity/tree/master/docs).

## <a name ="developing-the-provider"></a> Developing the Provider :hammer_and_pick:

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12 is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

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

## <a name="contribute"></a> Contribute :handshake:

* [Refer our contribution guideline](https://github.com/cohesity/terraform-provider-cohesity/tree/master/CONTRIBUTING.md).

## <a name="suggest"></a> Suggestions and Feedback :raised_hand:

We would love to hear from you. Please send your suggestions and feedback to: [cohesity-api-sdks@cohesity.com](mailto:cohesity-api-sdks@cohesity.com)

## License

Apache 2.0
