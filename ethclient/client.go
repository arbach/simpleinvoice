package ethclient

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type Client struct {
	client *ethclient.Client
}

type EthereumWallet struct {
	Address    string
	PrivateKey string
}

var Web3 *Client

func NewClient(nodeURL string) (*Client, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}
	Web3 = &Client{
		client: client,
	}
	return Web3, nil
}

func (web3 *Client) CreateAccount() (*EthereumWallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	walletKey := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	walletAddress := hexutil.Encode(hash.Sum(nil)[12:])

	wallet := EthereumWallet{
		Address:    walletAddress,
		PrivateKey: walletKey,
	}
	return &wallet, nil
}

func (web3 *Client) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := web3.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return big.NewInt(0), err
	}

	return balance, nil
}
