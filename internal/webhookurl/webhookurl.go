// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package webhookurl

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// IsBase64URL indicates whether a given string is a webhook URL composed of a
// single base64 string or multiple base64 strings ("segments") separated by
// commas.
func IsBase64URL(s string) bool {
	_, err := DecodeBase64(s)

	return err == nil
}

// DecodeBase64 supports decoding a given string as a single base64 string or
// multiple base64 encoded strings ("segments") separated by commas. Each
// segment may be generated using any of these supported base64 encoding
// formats as defined in RFC 4648:
//
//   - standard base64 encoding
//   - standard raw, unpadded base64 encoding
//   - alternate base64 encoding
//   - unpadded alternate base64 encoding
//
// Each segment will be safely combined and decoded. An error is returned if
// the given string is not composed of one or more segments encoded using one
// of these formats.
func DecodeBase64(input string) ([]byte, error) {
	input = strings.TrimSpace(input)

	if strings.Contains(input, ",") {
		csvBase64Strings := strings.Split(input, ",")

		combined, err := JoinBase64Segments(csvBase64Strings...)
		if err == nil {
			return decodeBase64(combined)
		}

		input = strings.ReplaceAll(input, ",", "")
	}

	return decodeBase64(input)
}

// EncodeToBase64String encodes the given input using unpadded alternate
// base64 encoding (base64url).
func EncodeToBase64String(input []byte) string {
	return base64.RawURLEncoding.EncodeToString(input)
}

// decodeBase64 decodes base64 strings originally generated using any of the
// supported base64 encoding formats as defined in RFC 4648:
//
// - standard base64 encoding
// - standard raw, unpadded base64 encoding
// - alternate base64 encoding
// - unpadded alternate base64 encoding
func decodeBase64(input string) ([]byte, error) {
	input = strings.TrimSpace(input)

	if data, err := base64.StdEncoding.DecodeString(input); err == nil {
		return data, nil
	}

	if data, err := base64.RawStdEncoding.DecodeString(input); err == nil {
		return data, nil
	}

	if data, err := base64.URLEncoding.DecodeString(input); err == nil {
		return data, nil
	}

	return base64.RawURLEncoding.DecodeString(input)
}

// JoinBase64Segments safely joins one or more base64 encoded strings together
// as a combined string using unpadded alternate base64 encoding (base64url).
// See also the DecodeBase64 function.
func JoinBase64Segments(encodedStrs ...string) (string, error) {
	var totalBytes []byte

	// Decode each string individually and append to the byte slice.
	for _, str := range encodedStrs {
		decoded, err := DecodeBase64(str)
		if err != nil {
			return "", fmt.Errorf("failed to decode segment: %w", err)
		}
		totalBytes = append(totalBytes, decoded...)
	}

	// Re-encode the fully combined bytes.
	return base64.RawURLEncoding.EncodeToString(totalBytes), nil
}
