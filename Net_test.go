package describe

import (
	"net"
	"testing"
)

func TestNet_IsPrivateIP(t *testing.T) {
	if !IsPrivateIP(net.ParseIP("10.0.1.2")) {
		t.Fatal("should be true")
	}
	if !IsPrivateIP(net.ParseIP("172.18.1.2")) {
		t.Fatal("should be true")
	}
	if !IsPrivateIP(net.ParseIP("192.168.1.1")) {
		t.Fatal("should be true")
	}
	if IsPrivateIP(net.ParseIP("11.23.44.4")) {
		t.Fatal("should be false")
	}
	if IsPrivateIP(net.ParseIP("172.100.2.3")) {
		t.Fatal("should be false")
	}
	if IsPrivateIP(net.ParseIP("127.0.0.1")) {
		t.Fatal("should be false")
	}
}

func TestNet_IsPublicIP(t *testing.T) {
	if !IsPublicIP(net.ParseIP("123.3.3.5")) {
		t.Fatal("should be true")
	}
	if IsPublicIP(net.ParseIP("127.0.0.1")) {
		t.Fatal("should be false")
	}
	if IsPublicIP(net.ParseIP("10.0.0.1")) {
		t.Fatal("should be false")
	}
}
