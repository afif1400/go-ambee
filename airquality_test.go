package go_ambee

import (
	"context"
	"errors"
	"os"
	"testing"
)

type AuthTest struct {
	Key string
}

// function to initialize the test auth
func (a *AuthTest) Init() (*AuthTest, error) {
	key, ok := os.LookupEnv("AMBEE_KEY")
	if !ok {
		ErrNoKey := errors.New("AMBEE_KEY not found")
		return nil, ErrNoKey
	}

	return &AuthTest{
		Key: key,
	}, nil

}

func TestGetAirQualityByCity(t *testing.T) {

	auth, err := (&AuthTest{}).Init()
	if err != nil {
		t.Error(err)
	}
	c := NewClient(auth.Key)
	ctx := context.Background()
	res, err := c.GetAirQualityByCity(ctx, "Raichur")

	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("res: %v", res)
}

func TestGetAirQualityByPostalCode(t *testing.T) {

	auth, err := (&AuthTest{}).Init()
	if err != nil {
		t.Error(err)
	}

	c := NewClient(auth.Key)
	ctx := context.Background()
	res, err := c.GetAirQualityByPostalCode(ctx, "584101", "IN")

	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("res: %v", res)
}
