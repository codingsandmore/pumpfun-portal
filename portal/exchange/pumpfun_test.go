package exchange

import (
	"github.com/joho/godotenv"
	log2 "github.com/rs/zerolog/log"
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

	sr, err := p.SwapSolForToken(0.005, "CotwVjzJnjhUEYw3xFLuhVw1B3xcfdJDEkNYnfcaf2Bb", 0.5, 0.01)

	assert.NoError(t, err)

	sr2, err2 := p.SwapTokenForSol(sr.Filled, "CotwVjzJnjhUEYw3xFLuhVw1B3xcfdJDEkNYnfcaf2Bb", 0.5, 0.01)

	assert.NoError(t, err2)

	log2.Info().Any("response", sr2).Msgf("test finished")

}

func TestPumpFun_SwapSolForTokenWithApiKeyAndValidTokenFailsDueToSimulationError(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.NoError(t, err)

	p := NewPumpFun(os.Getenv("PUMPFUN_API_KEY"), &http.Client{})

	_, err = p.SwapSolForToken(0.005, "8ReVeUanKktF6mAhCcf3JgCseAWjbo1rqUeNqwymuAQ1", 0.5, 0.01)

	assert.Error(t, err)
}

func TestPumpFun_SwapSolForTokenWithApiKeyAndEmptyToken(t *testing.T) {
	err := godotenv.Load("../../.env")

	assert.NoError(t, err)

	p := NewPumpFun(os.Getenv("PUMPFUN_API_KEY"), &http.Client{})

	_, err = p.SwapSolForToken(0.05, "", 2, 0.01)

	assert.Error(t, err)
}
