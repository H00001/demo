package component_test

import (
	"context"
	"io/ioutil"
	"testing"

	http_client "github.com/jukylin/esim/http"
	"github.com/jukylin/esim/log"
	"github.com/stretchr/testify/assert"
)

// go test
func TestControllers_Esim(t *testing.T) {
	logger := log.NewLogger()

	client := http_client.NewClient()
	ctx := context.Background()
	resp, err := client.Get(ctx, "http://localhost:8080")

	if err != nil {
		logger.Errorf(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf(err.Error())
	}
	logger.Debugf(string(body))
	assert.Equal(t, 200, resp.StatusCode)
}
