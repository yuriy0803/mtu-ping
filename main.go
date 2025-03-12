package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// findBestMTU determines the maximum MTU using a binary search
func findBestMTU(host string) int {
	low, high := 1200, 1500
	var mtu int

	for low <= high {
		mid := (low + high) / 2
		cmd := exec.Command("ping", "-f", "-l", strconv.Itoa(mid), "-n", "1", host) // Windows syntax

		output, err := cmd.CombinedOutput()
		if err == nil && !strings.Contains(string(output), "Packet needs to be fragmented") {
			mtu = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return mtu
}

// getBestPing measures the best ping for a target
func getBestPing(host string) float64 {
	cmd := exec.Command("ping", "-n", "4", host) // Windows: send 4 packets

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while pinging:", err)
		return -1
	}

	// Regular expression to extract "Minimum = xxms, Maximum = xxms, Average = xxms"
	re := regexp.MustCompile(`Average = (\d+)ms`)
	match := re.FindStringSubmatch(string(output))

	if len(match) > 1 {
		ping, _ := strconv.Atoi(match[1])
		return float64(ping)
	}

	fmt.Println("Error measuring ping")
	return -1
}

func main() {
	host := "8.8.8.8"

	mtu := findBestMTU(host)
	fmt.Printf("Best MTU for %s: %d\n", host, mtu)

	ping := getBestPing(host)
	if ping != -1 {
		fmt.Printf("Best Ping for %s: %.2f ms\n", host, ping)
	}
}
