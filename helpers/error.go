package helpers

import "github.com/sirupsen/logrus"

func FuncError(err error, msg string) {
	logger := logrus.New()
	if err != nil {
		logger.WithField("Message:", err).Error(msg)
	}
}
