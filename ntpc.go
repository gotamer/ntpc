package main

import (
	"encoding/binary"
	"flag"
	"net"
	"os"
	"os/exec"
	"os/user"
	"time"
)

const ntpEpochOffset = 2208988800

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
type packet struct {
	Settings       uint8  // leap yr indicator, ver number, and mode
	Stratum        uint8  // stratum of local clock
	Poll           int8   // poll exponent
	Precision      int8   // precision exponent
	RootDelay      uint32 // root delay
	RootDispersion uint32 // root dispersion
	ReferenceID    uint32 // reference id
	RefTimeSec     uint32 // reference timestamp sec
	RefTimeFrac    uint32 // reference timestamp fractional
	OrigTimeSec    uint32 // origin time secs
	OrigTimeFrac   uint32 // origin time fractional
	RxTimeSec      uint32 // receive time secs
	RxTimeFrac     uint32 // receive time frac
	TxTimeSec      uint32 // transmit time secs
	TxTimeFrac     uint32 // transmit time frac
}

var debug bool

// This program implements a trivial NTP client over UDP.
func main() {
	var err error
	var host string
	var save bool
	var conn net.Conn

	flag.StringVar(&host, "e", "pool.ntp.org", "NTP host")
	flag.BoolVar(&save, "s", false, "Update system date & time")
	flag.BoolVar(&debug, "d", false, "Show detailed results")
	flag.Parse()

	host = host + ":123"

	logger()

	if save && !isRoot() {
		Error.Println("System clock update can only be done by root")
		save = false
	}

	// Setup a UDP connection
	conn, err = net.Dial("udp", host)
	if err != nil {
		// Fallback host
		if host != "pool.ntp.org" {
			conn, err = net.Dial("udp", "pool.ntp.org")
			if err != nil {
				Error.Fatalf("Failed to connect: %v", err)
				os.Exit(1)
			}
		} else {
			Error.Fatalf("Failed to connect: %v", err)
			os.Exit(1)
		}
	}
	defer conn.Close()
	Debug.Printf("Connected to %v", host)

	if err = conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		Error.Fatalf("Failed to set deadline: %v", err)
	}
	Debug.Print("Got responce")

	// configure request settings by specifying the first byte as
	// 00 011 011 (or 0x1B)
	// |  |   +-- client mode (3)
	// |  + ----- version (3)
	// + -------- leap year indicator, 0 no warning
	req := &packet{Settings: 0x1B}

	// send time request
	if err = binary.Write(conn, binary.BigEndian, req); err != nil {
		Error.Fatalf("failed to send request: %v", err)
	}

	// block to receive server response
	rsp := &packet{}
	dateLoc := time.Now()
	if err := binary.Read(conn, binary.BigEndian, rsp); err != nil {
		Error.Fatalf("Failed to read server response, may not be an NTP server: %v", err)
		os.Exit(1)
	}

	// On POSIX-compliant OS, time is expressed
	// using the Unix time epoch (or secs since year 1970).
	// NTP seconds are counted since 1900 and therefore must
	// be corrected with an epoch offset to convert NTP seconds
	// to Unix time by removing 70 yrs of seconds (1970-1900)
	// or 2208988800 seconds.
	secs := float64(rsp.TxTimeSec) - ntpEpochOffset
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32 // convert fractional to nanos

	dateNTP := time.Unix(int64(secs), nanos)
	dateNTPF := dateNTP.Format(time.RFC3339Nano)

	var updated = "Not Updated"
	if save && dateLoc.Sub(dateNTP).Abs() > time.Millisecond*100 {
		out, err := exec.Command("/bin/date", "-s", dateNTPF).Output()
		if err != nil {
			Error.Fatalf("Date out: %v, cmd: %v", out, err)
		}
		updated = "Updated"
	}

	Debug.Printf("Time Diff %s: %v", updated, dateLoc.Sub(dateNTP))
	Debug.Println("Time Local: ", dateLoc.Format(time.RFC3339Nano))
	Debug.Println("Time NTP  : ", dateNTPF)
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		Error.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
}
