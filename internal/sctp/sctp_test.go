package sctp

import (
	"testing"

	"github.com/pkg/errors"
)

func TestParse(t *testing.T) {
	pkt := &Packet{}

	err := pkt.Unmarshal([]byte{})
	if err == nil {
		t.Errorf("Parse should fail when a packet is too small to be SCTP")
	}

	headerOnly := []byte{0x13, 0x88, 0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0x81, 0x46, 0x9d, 0xfc}
	err = pkt.Unmarshal(headerOnly)
	if err != nil {
		t.Error(errors.Wrap(err, "Parse failed for SCTP packet with no chunks"))
	} else if pkt.SourcePort != 5000 {
		t.Error(errors.Wrapf(err, "Parse passed for SCTP packet, but got incorrect source port exp: %d act: %d", 5000, pkt.SourcePort))
	} else if pkt.DestinationPort != 5000 {
		t.Error(errors.Wrapf(err, "Parse passed for SCTP packet, but got incorrect destination port exp: %d act: %d", 5000, pkt.DestinationPort))
	} else if pkt.VerificationTag != 0 {
		t.Error(errors.Wrapf(err, "Parse passed for SCTP packet, but got incorrect verification tag exp: %d act: %d", 0, pkt.VerificationTag))
	} else if pkt.Checksum != 2168888828 {
		t.Error(errors.Wrapf(err, "Parse passed for SCTP packet, but got incorrect checksum exp: %d act: %d", 2168888828, pkt.Checksum))
	}

	rawChunk := []byte{0x13, 0x88, 0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0x81, 0x46, 0x9d, 0xfc, 0x01, 0x00, 0x00, 0x56, 0x55,
		0xb9, 0x64, 0xa5, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0xe8, 0x6d, 0x10, 0x30, 0xc0, 0x00, 0x00, 0x04, 0x80,
		0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x9f, 0xeb, 0xbb, 0x5c, 0x50,
		0xc9, 0xbf, 0x75, 0x9c, 0xb1, 0x2c, 0x57, 0x4f, 0xa4, 0x5a, 0x51, 0xba, 0x60, 0x17, 0x78, 0x27, 0x94, 0x5c, 0x31, 0xe6,
		0x5d, 0x5b, 0x09, 0x47, 0xe2, 0x22, 0x06, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00}
	err = pkt.Unmarshal(rawChunk)
	if err != nil {
		t.Error(errors.Wrap(err, "Parse failed, has chunk"))
	}
}