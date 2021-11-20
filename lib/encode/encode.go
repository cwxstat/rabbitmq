package encode

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"

	"github.com/cwxstat/rabbitmq/lib/compress"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
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

// TODO: (mmc)  Should this be here?  And checkIfexist may need to be cleaned up
type HS struct {
	count   int64
	File    string
	data    string
	DestDir string
	FS      fs.FileMode
}

func (h *HS) checkIfexist() error {

	if h.DestDir != "" {
		_, err := os.Stat(h.DestDir)
		if err != nil {
			if os.IsNotExist(err) {
				if h.FS == 0 {
					h.FS = 0755
				}
				if err := os.MkdirAll(h.DestDir, os.FileMode(h.FS)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (h *HS) Handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		h.count += 1
		h.data = string(d.Body)
		log.Printf("here... file:(%s)\n", h.File)
		err := h.checkIfexist()
		if err != nil {
			log.Printf("Handle file create error: %s\n", err)
		}
		err = WriteDecode(h.File, h.data)
		if err != nil {
			log.Printf(
				"ERROR: got count(%d) %dB delivery: [%v] %q: %q\n\n%v\n",
				h.count,
				len(d.Body),
				d.DeliveryTag,
				"body..snip",
				d.AppId,
				err,
			)
		} else {
			log.Printf(
				"got count(%d) %dB delivery: [%v] %q: %q",
				h.count,
				len(d.Body),
				d.DeliveryTag,
				"d.Body",
				d.AppId,
			)

			err = compress.UnCompress(h.File, h.DestDir)
			if err != nil {
				fmt.Printf("Issue with uncompress\n")
				fmt.Println(err)
			}

		}
		d.Ack(false)

	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
