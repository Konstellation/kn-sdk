package keybase

import (
	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func SaveCoinKey(clientRoot, keyName, keyPass, keyMnemonic string, overwrite bool) (sdk.AccAddress, string, error) {
	// TODO
	// kn-sdk? appName
	kr, err := keyring.New("kn-sdk", keyring.BackendPass, clientRoot, nil)

	var info keyring.Info
	if keyMnemonic == "" {
		// generate a private key, with recovery phrase
		info, keyMnemonic, err = kr.NewMnemonic(keyName, keyring.English, keyPass, hd.Secp256k1)
	} else {
		// account 0 "Account number for HD derivation"
		// index 0 "Address index number for HD derivation"
		// TODO
		// 118? coinType
		path := hd.CreateHDPath(118, 0, 0).String()
		info, err = kr.NewAccount(keyName, keyMnemonic, keyPass, path, hd.Secp256k1)
	}

	if err != nil {
		return []byte{}, "", err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), keyMnemonic, nil
}
