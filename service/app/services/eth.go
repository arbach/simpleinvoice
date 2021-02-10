package services

import (
	"math"
	"math/big"

	"github.com/arbach/simpleinvoice/ethclient"
)

type Service struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *Service {
	s := &Service{
		client: client,
	}
	return s
}

func (s *Service) GenerateAddress() (string, error) {
	account, err := s.client.CreateAccount()
	if err != nil {
		return "", err
	}

	return account.Address, nil
}

func (s *Service) GetBalanceInEther(address string) (float64, error) {
	balance, err := s.client.GetBalance(address)
	if err != nil {
		return 0.0, err
	}

	var a, b, result big.Rat
	a.SetInt(balance)
	b.SetFloat64(math.Pow10(18))

	result.Quo(&a, &b)
	c, _ := result.Float64()
	return c, nil
}
