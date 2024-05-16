package exchange

import (
	"github.com/joho/godotenv"
	"github.com/test-go/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestPumpFun_SwapSolForTokenNoApiKey(t *testing.T) {

	p := NewPumpFun("", &http.Client{})

	_, err := p.SwapSolForToken(0.05, "", 2, 0.01)

	assert.Error(t, err)
}
func TestPumpFun_SwapSolForTokenWithApiKey(t *testing.T) {
	err := godotenv.Load("../../.env")

	assert.NoError(t, err)

	p := NewPumpFun(os.Getenv("PUMPFUN_API_KEY"), &http.Client{})

	_, err = p.SwapSolForToken(0.05, "", 2, 0.01)

	assert.Error(t, err)
}

func TestPumpFun_SwapSolForTokenWithApiKeyAndValidToken(t *testing.T) {
	err := godotenv.Load("../../.env")

	assert.NoError(t, err)

	p := NewPumpFun(os.Getenv("PUMPFUN_API_KEY"), &http.Client{})

	_, err = p.SwapSolForToken(0.05, "5QqWdXWvJ1GUwLVQCfDxgeTcR5LhRMqnRT49HWMd23oX", 2, 0.01)

	assert.NoError(t, err)
}

func TestPumpFun_SwapSolForTokenWithApiKeyAndEmptyToken(t *testing.T) {
	err := godotenv.Load("../../.env")

	assert.NoError(t, err)

	p := NewPumpFun(os.Getenv("PUMPFUN_API_KEY"), &http.Client{})

	_, err = p.SwapSolForToken(0.05, "", 2, 0.01)

	assert.Error(t, err)
}
