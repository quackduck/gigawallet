package dogecoin

import (
	giga "github.com/dogecoinfoundation/gigawallet/pkg"

	"github.com/jaxlotl/go-libdogecoin"
)

var _ giga.DogecoinL1 = DogecoinL1Libdogecoin{}

/* Returns a Mocked giga.DogecoinL1 implementor */
func NewL1Libdogecoin(config giga.Config) (DogecoinL1Libdogecoin, error) {
	return DogecoinL1Libdogecoin{}, nil
}

type DogecoinL1Libdogecoin struct {
}

func (d DogecoinL1Libdogecoin) Send(txn giga.Txn) error {
	return nil
}

func (d DogecoinL1Libdogecoin) MakeAddress() (giga.Address, error) {
	libdogecoin.W_context_start()
	priv, pub := libdogecoin.W_generate_hd_master_pub_keypair(false)
	libdogecoin.W_context_stop()
	return giga.Address{priv, pub}, nil
}
