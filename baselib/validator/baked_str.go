/* !!
 * File: str.go
 * File Created: Thursday, 30th September 2021 10:13:38 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 30th September 2021 10:13:38 am
 
 */

package validator

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"strings"
	"unicode/utf8"

	urn "github.com/leodido/go-urn"
)

func IsURLEncodedString(str string) bool {
	return URLEncodedRegex.MatchString(str)
}

func IsHTMLEncodedString(str string) bool {
	return HTMLEncodedRegex.MatchString(str)
}

func IsHTMLString(str string) bool {
	return HTMLRegex.MatchString(str)
}

// IsUniqueIterator is the validation function for validating if each array|slice|map value is unique
func IsUniqueIterator(obj interface{}) bool {
	v := reflect.ValueOf(struct{}{})

	objRf := reflect.ValueOf(obj)

	switch objRf.Kind() {
	case reflect.Slice, reflect.Array:
		if reflect.TypeOf(obj).Elem().Kind() == reflect.Ptr {
			m := reflect.MakeMap(reflect.MapOf(reflect.ValueOf("").Type(), v.Type()))

			for i := 0; i < objRf.Len(); i++ {
				m.SetMapIndex(reflect.ValueOf(fmt.Sprintf("%+v", objRf.Index(i))), v)
			}
			return objRf.Len() == m.Len()
		}

		m := reflect.MakeMap(reflect.MapOf(objRf.Type().Elem(), v.Type()))

		for i := 0; i < objRf.Len(); i++ {
			m.SetMapIndex(objRf.Index(i), v)
		}
		return objRf.Len() == m.Len()

	case reflect.Map:
		if reflect.TypeOf(obj).Elem().Kind() == reflect.Ptr {
			m := reflect.MakeMap(reflect.MapOf(reflect.ValueOf("").Type(), v.Type()))

			for _, k := range objRf.MapKeys() {
				m.SetMapIndex(reflect.ValueOf(fmt.Sprintf("%+v", objRf.MapIndex(k))), v)
			}

			return objRf.Len() == m.Len()
		}

		m := reflect.MakeMap(reflect.MapOf(objRf.Type().Elem(), v.Type()))

		for _, k := range objRf.MapKeys() {
			m.SetMapIndex(objRf.MapIndex(k), v)
		}

		return objRf.Len() == m.Len()
	default:
		panic(fmt.Sprintf("Bad field type %T", objRf.Interface()))
	}
}

// IsMAC is the validation function for validating if the field's value is a valid MAC address.
func IsMACString(str string) bool {
	_, err := net.ParseMAC(str)

	return err == nil
}

// IsCIDRv4 is the validation function for validating if the field's value is a valid v4 CIDR address.
func IsCIDRv4String(str string) bool {
	ip, _, err := net.ParseCIDR(str)

	return err == nil && ip.To4() != nil
}

// IsCIDRv6 is the validation function for validating if the field's value is a valid v6 CIDR address.
func IsCIDRv6String(str string) bool {
	ip, _, err := net.ParseCIDR(str)

	return err == nil && ip.To4() == nil
}

// IsCIDR is the validation function for validating if the field's value is a valid v4 or v6 CIDR address.
func IsCIDRString(str string) bool {
	_, _, err := net.ParseCIDR(str)

	return err == nil
}

// IsIPv4 is the validation function for validating if a value is a valid v4 IP address.
func IsIPv4String(str string) bool {
	ip := net.ParseIP(str)

	return ip != nil && ip.To4() != nil
}

// IsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func IsIPv6String(str string) bool {
	ip := net.ParseIP(str)

	return ip != nil && ip.To4() == nil
}

// IsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func IsIPString(str string) bool {

	ip := net.ParseIP(str)

	return ip != nil
}

// IsSSN is the validation function for validating if the field's value is a valid SSN.
func IsSSNString(str string) bool {
	if len(str) != 11 {
		return false
	}

	return SSNRegex.MatchString(str)
}

// IsLongitude is the validation function for validating if the field's value is a valid longitude coordinate.
func IsLongitudeString(str string) bool {
	return LongitudeRegex.MatchString(str)
}

// IsLatitude is the validation function for validating if the field's value is a valid latitude coordinate.
func IsLatitudeString(str string) bool {
	return LatitudeRegex.MatchString(str)
}

// IsDataURI is the validation function for validating if the field's value is a valid data URI.
func IsDataURIString(str string) bool {

	uri := strings.SplitN(str, ",", 2)

	if len(uri) != 2 {
		return false
	}

	if !DataURIRegex.MatchString(uri[0]) {
		return false
	}

	return Base64Regex.MatchString(uri[1])
}

// HasMultiByteCharacter is the validation function for validating if the field's value has a multi byte character.
func hasMultiByteCharacterString(str string) bool {
	if len(str) == 0 {
		return true
	}

	return MultibyteRegex.MatchString(str)
}

// IsPrintableASCII is the validation function for validating if the field's value is a valid printable ASCII character.
func IsPrintableASCIIString(str string) bool {
	return PrintableASCIIRegex.MatchString(str)
}

// IsASCII is the validation function for validating if the field's value is a valid ASCII character.
func IsASCIIString(str string) bool {
	return ASCIIRegex.MatchString(str)
}

// IsUUID5 is the validation function for validating if the field's value is a valid v5 UUID.
func IsUUID5String(str string) bool {
	return UUID5Regex.MatchString(str)
}

// IsUUID4 is the validation function for validating if the field's value is a valid v4 UUID.
func IsUUID4String(str string) bool {
	return UUID4Regex.MatchString(str)
}

// IsUUID3 is the validation function for validating if the field's value is a valid v3 UUID.
func IsUUID3String(str string) bool {
	return UUID3Regex.MatchString(str)
}

// IsUUID is the validation function for validating if the field's value is a valid UUID of any version.
func IsUUIDString(str string) bool {
	return UUIDRegex.MatchString(str)
}

// IsUUID5RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v5 UUID.
func IsUUID5RFC4122String(str string) bool {
	return UUID5RFC4122Regex.MatchString(str)
}

// IsUUID4RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v4 UUID.
func IsUUID4RFC4122String(str string) bool {
	return UUID4RFC4122Regex.MatchString(str)
}

// IsUUID3RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v3 UUID.
func IsUUID3RFC4122String(str string) bool {
	return UUID3RFC4122Regex.MatchString(str)
}

// IsUUIDRFC4122 is the validation function for validating if the field's value is a valid RFC4122 UUID of any version.
func IsUUIDRFC4122String(str string) bool {
	return UUIDRFC4122Regex.MatchString(str)
}

func IsUsernameString(str string) bool {
	return UsernameRegex.MatchString(str)
}

// IsISBN is the validation function for validating if the field's value is a valid v10 or v13 ISBN.
func IsISBNString(str string) bool {
	return IsISBN10String(str) || IsISBN13String(str)
}

// IsISBN13 is the validation function for validating if the field's value is a valid v13 ISBN.
func IsISBN13String(str string) bool {
	s := strings.Replace(strings.Replace(str, "-", "", 4), " ", "", 4)

	if !ISBN13Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	factor := []int32{1, 3}

	for i = 0; i < 12; i++ {
		checksum += factor[i%2] * int32(s[i]-'0')
	}

	return (int32(s[12]-'0'))-((10-(checksum%10))%10) == 0
}

// IsISBN10 is the validation function for validating if the field's value is a valid v10 ISBN.
func IsISBN10String(str string) bool {

	s := strings.Replace(strings.Replace(str, "-", "", 3), " ", "", 3)

	if !ISBN10Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	for i = 0; i < 9; i++ {
		checksum += (i + 1) * int32(s[i]-'0')
	}

	if s[9] == 'X' {
		checksum += 10 * 10
	} else {
		checksum += 10 * int32(s[9]-'0')
	}

	return checksum%11 == 0
}

// IsEthereumAddress is the validation function for validating if the field's value is a valid ethereum address based currently only on the format
func IsEthereumAddressString(str string) bool {
	address := str

	if !EthAddressRegex.MatchString(address) {
		return false
	}

	if EthaddressRegexUpper.MatchString(address) || EthAddressRegexLower.MatchString(address) {
		return true
	}

	// checksum validation is blocked by https://github.com/golang/crypto/pull/28

	return true
}

// IsBitcoinAddress is the validation function for validating if the field's value is a valid btc address
func IsBitcoinAddressString(str string) bool {
	address := str

	if !BtcAddressRegex.MatchString(address) {
		return false
	}

	alphabet := []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

	decode := [25]byte{}

	for _, n := range []byte(address) {
		d := bytes.IndexByte(alphabet, n)

		for i := 24; i >= 0; i-- {
			d += 58 * int(decode[i])
			decode[i] = byte(d % 256)
			d /= 256
		}
	}

	h := sha256.New()
	_, _ = h.Write(decode[:21])
	d := h.Sum([]byte{})
	h = sha256.New()
	_, _ = h.Write(d)

	validchecksum := [4]byte{}
	computedchecksum := [4]byte{}

	copy(computedchecksum[:], h.Sum(d[:0]))
	copy(validchecksum[:], decode[21:])

	return validchecksum == computedchecksum
}

// IsBitcoinBech32Address is the validation function for validating if the field's value is a valid bech32 btc address
func IsBitcoinBech32AddressString(str string) bool {
	address := str

	if !BtcLowerAddressRegexBech32.MatchString(address) && !BtcUpperAddressRegexBech32.MatchString(address) {
		return false
	}

	am := len(address) % 8

	if am == 0 || am == 3 || am == 5 {
		return false
	}

	address = strings.ToLower(address)

	alphabet := "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

	hr := []int{3, 3, 0, 2, 3} // the human readable part will always be bc
	addr := address[3:]
	dp := make([]int, 0, len(addr))

	for _, c := range addr {
		dp = append(dp, strings.IndexRune(alphabet, c))
	}

	ver := dp[0]

	if ver < 0 || ver > 16 {
		return false
	}

	if ver == 0 {
		if len(address) != 42 && len(address) != 62 {
			return false
		}
	}

	values := append(hr, dp...)

	GEN := []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

	p := 1

	for _, v := range values {
		b := p >> 25
		p = (p&0x1ffffff)<<5 ^ v

		for i := 0; i < 5; i++ {
			if (b>>uint(i))&1 == 1 {
				p ^= GEN[i]
			}
		}
	}

	if p != 1 {
		return false
	}

	b := uint(0)
	acc := 0
	mv := (1 << 5) - 1
	var sw []int

	for _, v := range dp[1 : len(dp)-6] {
		acc = (acc << 5) | v
		b += 5
		for b >= 8 {
			b -= 8
			sw = append(sw, (acc>>b)&mv)
		}
	}

	if len(sw) < 2 || len(sw) > 40 {
		return false
	}

	return true
}

// ExcludesRune is the validation function for validating that the field's value does not contain the rune specified within the param.
func excludesRuneString(str string) bool {
	return !containsRuneString(str)
}

// ExcludesAll is the validation function for validating that the field's value does not contain any of the characters specified within the param.
func excludesAllString(str string) bool {
	return !containsAnyString(str)
}

// Excludes is the validation function for validating that the field's value does not contain the text specified within the param.
func excludesString(str string) bool {
	return !containsString(str)
}

// ContainsRune is the validation function for validating that the field's value contains the rune specified within the param.
func containsRuneString(str string) bool {
	r, _ := utf8.DecodeRuneInString(str)

	return strings.ContainsRune(str, r)
}

// ContainsAny is the validation function for validating that the field's value contains any of the characters specified within the param.
func containsAnyString(str string) bool {
	return strings.ContainsAny(str, str)
}

// Contains is the validation function for validating that the field's value contains the text specified within the param.
func containsString(str string) bool {
	return strings.Contains(str, str)
}

// StartsWith is the validation function for validating that the field's value starts with the text specified within the param.
func startsWithString(str string) bool {
	return strings.HasPrefix(str, str)
}

// EndsWith is the validation function for validating that the field's value ends with the text specified within the param.
func endsWithString(str string) bool {
	return strings.HasSuffix(str, str)
}

// IsBase64 is the validation function for validating if the current field's value is a valid base 64.
func IsBase64String(str string) bool {
	return Base64Regex.MatchString(str)
}

// IsBase64URL is the validation function for validating if the current field's value is a valid base64 URL safe string.
func IsBase64URLString(str string) bool {
	return Base64URLRegex.MatchString(str)
}

// IsURI is the validation function for validating if the current field's value is a valid URI.
func IsURIString(str string) bool {
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(str, "#"); i > -1 {
		str = str[:i]
	}

	if len(str) == 0 {
		return false
	}

	_, err := url.ParseRequestURI(str)

	return err == nil
}

// IsURL is the validation function for validating if the current field's value is a valid URL.
func IsURLString(str string) bool {
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(str, "#"); i > -1 {
		str = str[:i]
	}

	if len(str) == 0 {
		return false
	}

	url, err := url.ParseRequestURI(str)

	if err != nil || url.Scheme == "" {
		return false
	}

	return err == nil
}

// isUrnRFC2141 is the validation function for validating if the current field's value is a valid URN as per RFC 2141.
func IsUrnRFC2141String(str string) bool {
	_, match := urn.Parse([]byte(str))

	return match
}

// IsFile is the validation function for validating if the current field's value is a valid file path.
func IsFileString(str string) bool {
	fileInfo, err := os.Stat(str)
	if err != nil {
		return false
	}

	return !fileInfo.IsDir()
}

// IsE164 is the validation function for validating if the current field's value is a valid e.164 formatted phone number.
func IsE164String(str string) bool {
	return E164Regex.MatchString(str)
}

// IsEmail is the validation function for validating if the current field's value is a valid email address.
func IsEmailString(str string) bool {
	return EmailRegex.MatchString(str)
}

// IsHSLA is the validation function for validating if the current field's value is a valid HSLA color.
func IsHSLAString(str string) bool {
	return HSLaRegex.MatchString(str)
}

// IsHSL is the validation function for validating if the current field's value is a valid HSL color.
func IsHSLString(str string) bool {
	return HSLRegex.MatchString(str)
}

// IsRGBA is the validation function for validating if the current field's value is a valid RGBA color.
func IsRGBAString(str string) bool {
	return RgbaRegex.MatchString(str)
}

// IsRGB is the validation function for validating if the current field's value is a valid RGB color.
func IsRGBString(str string) bool {
	return RgbRegex.MatchString(str)
}

// IsHEXColor is the validation function for validating if the current field's value is a valid HEX color.
func IsHEXColorString(str string) bool {
	return HexcolorRegex.MatchString(str)
}

// IsHexadecimal is the validation function for validating if the current field's value is a valid hexadecimal.
func IsHexadecimalString(str string) bool {
	return HexadecimalRegex.MatchString(str)
}

// IsNumber is the validation function for validating if the current field's value is a valid number.
func IsNumberString(str string) bool {
	return NumberRegex.MatchString(str)
}

// IsNumeric is the validation function for validating if the current field's value is a valid numeric value.
func IsNumericString(str string) bool {
	return NumericRegex.MatchString(str)
}

// IsAlphanum is the validation function for validating if the current field's value is a valid alphanumeric value.
func IsAlphanumString(str string) bool {
	return AlphaNumericRegex.MatchString(str)
}

// IsAlpha is the validation function for validating if the current field's value is a valid alpha value.
func IsAlphaString(str string) bool {
	return AlphaRegex.MatchString(str)
}

// IsAlphanumUnicode is the validation function for validating if the current field's value is a valid alphanumeric unicode value.
func IsAlphanumUnicodeString(str string) bool {
	return AlphaUnicodeNumericRegex.MatchString(str)
}

// IsAlphaUnicode is the validation function for validating if the current field's value is a valid alpha unicode value.
func IsAlphaUnicodeString(str string) bool {
	return AlphaUnicodeRegex.MatchString(str)
}

// IsTCP4AddrResolvable is the validation function for validating if the field's value is a resolvable tcp4 address.
func IsTCP4AddrResolvableString(str string) bool {
	if !IsIP4AddrString(str) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp4", str)
	return err == nil
}

// IsTCP6AddrResolvable is the validation function for validating if the field's value is a resolvable tcp6 address.
func IsTCP6AddrResolvableString(str string) bool {
	if !IsIP6AddrString(str) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp6", str)

	return err == nil
}

// IsTCPAddrResolvable is the validation function for validating if the field's value is a resolvable tcp address.
func IsTCPAddrResolvableString(str string) bool {
	if !IsIP4AddrString(str) && !IsIP6AddrString(str) {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp", str)

	return err == nil
}

// IsUDP4AddrResolvable is the validation function for validating if the field's value is a resolvable udp4 address.
func IsUDP4AddrResolvableString(str string) bool {
	if !IsIP4AddrString(str) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp4", str)

	return err == nil
}

// IsUDP6AddrResolvable is the validation function for validating if the field's value is a resolvable udp6 address.
func IsUDP6AddrResolvableString(str string) bool {
	if !IsIP6AddrString(str) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp6", str)

	return err == nil
}

// IsUDPAddrResolvable is the validation function for validating if the field's value is a resolvable udp address.
func IsUDPAddrResolvableString(str string) bool {
	if !IsIP4AddrString(str) && !IsIP6AddrString(str) {
		return false
	}

	_, err := net.ResolveUDPAddr("udp", str)

	return err == nil
}

// IsIP4AddrResolvable is the validation function for validating if the field's value is a resolvable ip4 address.
func IsIP4AddrResolvableString(str string) bool {

	if !IsIPv4String(str) {
		return false
	}

	_, err := net.ResolveIPAddr("ip4", str)

	return err == nil
}

// IsIP6AddrResolvable is the validation function for validating if the field's value is a resolvable ip6 address.
func IsIP6AddrResolvableString(str string) bool {

	if !IsIPv6String(str) {
		return false
	}

	_, err := net.ResolveIPAddr("ip6", str)

	return err == nil
}

// IsIPAddrResolvable is the validation function for validating if the field's value is a resolvable ip address.
func IsIPAddrResolvableString(str string) bool {

	if !IsIPString(str) {
		return false
	}

	_, err := net.ResolveIPAddr("ip", str)

	return err == nil
}

// IsUnixAddrResolvable is the validation function for validating if the field's value is a resolvable unix address.
func IsUnixAddrResolvableString(str string) bool {

	_, err := net.ResolveUnixAddr("unix", str)

	return err == nil
}

func IsIP4AddrString(str string) bool {

	val := str

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		val = val[0:idx]
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() != nil
}

func IsIP6AddrString(str string) bool {

	val := str

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		if idx != 0 && val[idx-1:idx] == "]" {
			val = val[1 : idx-1]
		}
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() == nil
}

func IsHostnameRFC952String(str string) bool {
	return HostnameRegexRFC952.MatchString(str)
}

func IsHostnameRFC1123String(str string) bool {
	return HostnameRegexRFC1123.MatchString(str)
}

func IsFQDNString(str string) bool {
	val := str

	if val == "" {
		return false
	}

	if val[len(val)-1] == '.' {
		val = val[0 : len(val)-1]
	}

	return strings.ContainsAny(val, ".") &&
		HostnameRegexRFC952.MatchString(val)
}

// IsDir is the validation function for validating if the current field's value is a valid directory.
func IsDirString(str string) bool {
	fileInfo, err := os.Stat(str)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
