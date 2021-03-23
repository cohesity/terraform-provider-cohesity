# Cohesity Terraform Provider

![](docs/assets/images/terraform.png)

## Overview

This project provides a Terraform Provider for interacting with the [Cohesity DataPlatform](https://www.cohesity.com/products/data-platform). It includes Terraform modules useful for automating common tasks and orchestrating workflows in your environment.

## Table of contents :scroll:

 - [Requirements](#get-startedt)
 - [Building the Provider](#building)
 - [Documentation](#doc)
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


## <a name="contribute"></a> Contribute :handshake:

* [Refer our contribution guideline](https://github.com/cohesity/terraform-provider-cohesity/tree/master/CONTRIBUTING.md).

## <a name="suggest"></a> Suggestions and Feedback :raised_hand:

We would love to hear from you. Please send your suggestions and feedback to: [cohesity-api-sdks@cohesity.com](mailto:cohesity-api-sdks@cohesity.com)

## License

Apache 2.0
