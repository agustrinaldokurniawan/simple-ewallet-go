package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

var (
	first_name string
	last_name  string
	currency   string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomNumberForPhone(n int) string {
	var sb strings.Builder

	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomFirstName() string {
	value := RandomString(8)
	first_name = value
	return value
}

func RandomLastName() string {
	value := RandomString(8)
	last_name = value
	return value
}

func RandomEmail() string {
	return first_name + "." + last_name + "@email.com"
}

func RandomMoney(c string) int64 {
	var value int64

	if c != "" {
		currency = c
	}

	switch currency {
	case "USD":
		value = RandomInt(0, 1000)
	case "IDR":
		value = RandomInt(0, 20) * 50000
	default:
		value = 0
	}

	return value
}

func RandomCurrency() string {
	currencies := []string{"USD", "IDR"}
	n := len(currencies)
	value := currencies[rand.Intn(n)]
	currency = value
	return value
}

func RandomPhoneNumber() string {
	switch currency {

	case "IDR":
		return "+62" + "8" + RandomNumberForPhone(10)

	case "USD":
		return "+1" + RandomNumberForPhone(10)

	default:
		return ""
	}
}
