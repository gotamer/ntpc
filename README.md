NTP Client
==========

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/gotamer/ntpc?tab=doc)

A simple Network Time Protocol (NTP) Client
-------------------------------------------

This NTP Client is specially designd for computers without a hardware clock, 
such as the Raspberry Pi, and for Laptops.

Installation
------------

- Download the latest release for your OS from [here](https://github.com/gotamer/ntpc/releases)

**Mapping Raspberry Pi Models to Architectures**
| Raspberry Pi Model | Architecture |
|-------------------|----------------|
| Raspberry Pi 1 (A, B, A+, B+), Zero, Zero W | `arm6` |
| Raspberry Pi 2 (v1.1) | `arm7` |
| Raspberry Pi 2 (v1.2), 3, 3+, CM3 | `arm7` |
| Raspberry Pi 4, 400, CM4 (32-bit OS) | `arm7` |
| Raspberry Pi 4, 400, CM4 (64-bit OS) | `arm64` |
| Raspberry Pi 5 (64-bit OS) | `arm64` |

- Rename the executable to ntpc

Usage
----------
Run manually, at system start as a service, and/or as a cron job
```
Usage of ntpc:
  -d	Show detailed debug results
  -e string
    	NTP host (default "pool.ntp.org")
  -s	Update system date & time

Example:
# ntpc -s -e en.pool.ntp.org
OR
# ntpc -s -e 10.10.10.1
```

NTP packet format
-----------------
This info is for coders!
```
// NTP packet format (v3 with optional v4 fields removed)
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |LI | VN  |Mode |    Stratum     |     Poll      |  Precision   |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         Root Delay                            |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         Root Dispersion                       |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                          Reference ID                         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                                                               |
// +                     Reference Timestamp (64)                  +
// |                                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                                                               |
// +                      Origin Timestamp (64)                    +
// |                                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                                                               |
// +                      Receive Timestamp (64)                   +
// |                                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                                                               |
// +                      Transmit Timestamp (64)                  +
// |                                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```


This is a fork of Vladimir Vivien's go-ntp-client. (Thank you Vladimir)  
The original code is explained by Vladimir Vivien in this writeup titled [Let's make an NTP client in Go](https://medium.com/learning-the-go-programming-language/lets-make-an-ntp-client-in-go-287c4b9a969f)
