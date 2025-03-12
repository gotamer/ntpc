# NTP Client
============

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/gotamer/ntpc?tab=doc)

This is an implementation of a simple NTP client in Go. It uses the `encoding/binary` package to encode and decode NTP packets sent to and received from a remote NTP server over UDP.

It also does:

- log via linux syslog
- local date update

I use it to periodically update my Linux laptop system time.
On my system it runs once an hour via a cron job.


This is a fork of Vladimir Vivien's go-ntp-client. (Thank you Vladimir)  
The original code is explained by Vladimir Vivien in this writeup titled [Let's make an NTP client in Go](https://medium.com/learning-the-go-programming-language/lets-make-an-ntp-client-in-go-287c4b9a969f)
