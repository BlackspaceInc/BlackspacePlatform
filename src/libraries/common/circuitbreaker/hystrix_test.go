package circuitbreaker

import (
	"net/http"
	"testing"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/h2non/gock.v1"

	"github.com/BlackspaceInc/common/testutil"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000",
	})
	logrus.SetLevel(logrus.DebugLevel)
	Client = &http.Client{}
}

func TestCallUsingResilienceAllFails(t *testing.T) {
	defer gock.Off()

	buildGockMatcherTimes(500, 4)
	hystrix.Flush()

	bytes, err := CallUsingCircuitBreaker(context.TODO(), "TEST", "http://quotes-service", "GET")

	if err == nil {
		return
	}

	if bytes == nil {
		return
	}
}

func TestCallUsingResilienceLastSucceeds(t *testing.T) {
	defer gock.Off()
	retries = 3
	buildGockMatcherTimes(500, 2)
	body := []byte("Some response")
	buildGockMatcherWithBody(200, string(body))
	hystrix.Flush()

	bytes, err := CallUsingCircuitBreaker(context.TODO(), "TEST", "http://quotes-service", "GET")

	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, bytes)
	testutil.AssertEqual(t, string(body), string(bytes))
}

func TestCallHystrixOpensAfterThresholdPassed(t *testing.T) {
	defer gock.Off()
	for a := 0; a < 6; a++ {
		buildGockMatcher(500)
	}
	hystrix.Flush()

	retries = 0
	hystrix.ConfigureCommand("TEST", hystrix.CommandConfig{
		RequestVolumeThreshold: 5,
	})
	for a := 0; a < 6; a++ {
		CallUsingCircuitBreaker(context.TODO(), "TEST", "http://quotes-service", "GET")
	}

	cb, _, _ := hystrix.GetCircuit("TEST")
	testutil.AssertTrue(t, cb.IsOpen())
}

func buildGockMatcherTimes(status int, times int) {
	for a := 0; a < times; a++ {
		buildGockMatcher(status)
	}
}

func buildGockMatcherWithBody(status int, body string) {
	gock.New("http://quotes-service").
		Reply(status).BodyString(body)
}

func buildGockMatcher(status int) {
	buildGockMatcherWithBody(status, "")
}
