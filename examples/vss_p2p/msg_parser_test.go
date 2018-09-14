package vss_p2p

import (
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser(t *testing.T) {
	suite := suites.MustFind("bn256")

	deal := &vss.EncryptedDeal{
		DHKey:[]byte("DHKey"),
		Signature:[]byte("Signature"),
		Nonce:[]byte("Nonce"),
		Cipher:[]byte("Cipher"),
	}
	recoverDeal := unpackEncryptedDeal(packEncryptedDeal(deal))
	require.Equal(t, deal, recoverDeal)

	response := & vss.Response{
		SessionID:[]byte("SessionID"),
		Index:0,
		Status:true,
		Signature:[]byte("Signature"),
	}
	recoverResponse := unpackResonse(packResonse(response))
	require.Equal(t, response, recoverResponse)

	justification := &vss.Justification{
		SessionID:[]byte("SessionID"),
		Index:0,
		Deal:&vss.Deal{
			SessionID:[]byte("SessionID"),
		},
		Signature:[]byte("Signature"),
	}
	recoverJustification := unpackJustification(suite, packJustification(justification))
	require.Equal(t, justification, recoverJustification)

	publicKey := suite.G2().Point()
	recoverPublicKey := unpackPublicKey(suite, packPublicKey(publicKey))
	require.Equal(t, publicKey, recoverPublicKey)

	publicKeys := []kyber.Point{suite.G2().Point(), suite.G2().Point()}
	recoverPublicKeys := unpackPublicKeys(suite, packPublicKeys(publicKeys))
	require.Equal(t, publicKeys, recoverPublicKeys)
}
