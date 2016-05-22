# Kafka scheduler - Command line interface

## Introduction

Kafkactl is a command line tool written in GO for controlling kafka scheduler
that runs on top of Apache Mesos. A simple REST client for the scheduler
remote API.

## Features

* Manage brokers (add, list, update, remove, start, stop)
* Manage topics (list, add, update, rebalance)
* CLI configuration file

## Installation

`go get github.com/eddyzags/kafka-cli`

## Usage

```
$ kafka --help
A simple REST client for the Kafka scheduler remote API.

Usage:
        kafka [command]

Available Commands:
        broker          Manage your brokers
        topic           Manage your topics

Flags:
        -v, --version   Print version information
        -a, --api       Target a specific remote API
```
