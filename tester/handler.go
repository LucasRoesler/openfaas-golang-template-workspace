package function

import (
	"strings"

	"handler/function/pulp"
	"handler/function/yo"

	"github.com/sirupsen/logrus"
)

// Handle a serverless request
func Handle(req []byte) string {

	logrus.Debug("Received request")

	b := strings.Builder{}
	b.WriteString(pulp.What(req))
	b.WriteString("\n")
	b.WriteString(yo.Yo(req))

	return b.String()
}
