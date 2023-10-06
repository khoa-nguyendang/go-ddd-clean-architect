package randomizer

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	chInt   = make(chan int, 10000)
	chInt64 = make(chan int64, 10000)
)

func init() {
	go genInt()
	go genInt64()
}

// genInt generates integer into channel to prevent concurrent access of rand function.
// See https://groups.google.com/forum/#!topic/golang-nuts/oyTWypHlHog for details.
func genInt() {
	// source is not safe for concurrent use by multiple goroutines, so we create
	// it for each goroutine.
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		chInt <- s.Int()
	}
}

func genInt64() {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		chInt64 <- s.Int63()
	}
}

// String generates a random string with given length from letters and digits
func String(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return StringByCharSet(n, s)
}

// StringAntiAmbiguous generates a random string with given length from letters and digits
// without ['o','O','0', 'i', 'I', 'l' , 'L', '1']
func StringAntiAmbiguous(n int) string {
	s := "abcdefghjkmnpqrstuvwxzABCDEFGHJKMNPQRSTUVWXZ23456789"
	return StringByCharSet(n, s)
}

// LowerASCII generates a random string with given length from lowercase letters
func LowerASCII(n int) string {
	s := "abcdefghijklmnopqrstuvwxyz"
	return StringByCharSet(n, s)
}

// Hex generates a random string with given length from hex numbers
func Hex(n int) string {
	s := "0123456789abcdef"
	return StringByCharSet(n, s)
}

// Digit generates a random string with given length (n) from digits
func Digit(n int) string {
	s := "0123456789"
	return StringByCharSet(n, s)
}

// SnowFlakeBase58 generates a random string from base58 base on snowflake
func SnowFlakeBase58() string {
	r := <-chInt64
	node, _ := snowflake.NewNode(r % 1023)
	return node.Generate().Base58()
}

// StringByCharSet generates a random string with given length from given char set
func StringByCharSet(n int, charSet string) string {
	// Ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	b := make([]byte, n)
	for i := range b {
		r := <-chInt64
		b[i] = charSet[r%int64(len(charSet))]
	}
	return string(b)
}
