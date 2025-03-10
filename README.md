---
title: NTP Client
description: Implementation of a trivial NTP client in Go
tags: [ntp, RFC5905, client]
---

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/gotamer/ntpc?tab=doc)


# A trivial NTP Client
This repository is an implementation of a (really) trivial NTP client in Go. It uses the `encoding/binary` package to encode and decode NTP packets sent to and received from a remote NTP server over UDP. You can learn more about NTP here, read the specs RFC5905, and find a (seemingly) way better Go NTP client, with many features implemented, [here](https://github.com/beevik/ntp).


This is a fork of Vladimir Vivien's go-ntp-client. (Thank you Vladimir)

#### I added following items to make it actually useful:

- syslog
- local date update.

I use it to periodically update my Linux laptop system time.
On my system it runs once an hour via a cron job.


The original code is explained by Vladimir Vivien in this writeup titled [Let's make an NTP client in Go](https://medium.com/learning-the-go-programming-language/lets-make-an-ntp-client-in-go-287c4b9a969f).
