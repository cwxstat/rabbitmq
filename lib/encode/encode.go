package encode

import (
	"encoding/base64"
	"os"

)

func ReadEncode(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	encoded := Encode(data)
	return encoded, err

}

func WriteDecode(file string, msg string) error {
	data, err := Decode(msg)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
	return err
}

func Encode(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	return encoded
}

func Decode(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return decoded, err
}

