package msgParser

import (
	"log"
	"testing"

	"github.com/DOSNetwork/core/examples/ethereum_vss_p2p/internalMsg"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	suite := suites.MustFind("bn256")

	deal := &vss.EncryptedDeal{
		DHKey:[]byte("DHKey"),
		Signature:[]byte("Signature"),
		Nonce:[]byte("Nonce"),
		Cipher:[]byte("Cipher"),
	}

	switch content := PackEncryptedDeal(deal).(type) {
	case *internalMsg.EncryptedDeal:
		recoverDeal := UnpackEncryptedDeal(content)
		require.Equal(t, deal, recoverDeal)
	default:
		log.Fatal("Type mismatch")
	}


	deals := []*vss.EncryptedDeal{deal, deal}

	switch content := PackEncryptedDeals(deals).(type) {
	case *internalMsg.EncryptedDeals:
		recoverDeals := UnpackEncryptedDeals(content)
		require.Equal(t, deals, recoverDeals)
	default:
		log.Fatal("Type mismatch")
	}


	response := & vss.Response{
		SessionID:[]byte("SessionID"),
		Index:0,
		Status:true,
		Signature:[]byte("Signature"),
	}

	switch content := PackResponse(response).(type) {
	case *internalMsg.Response:
		recoverResponse := UnpackResponse(content)
		require.Equal(t, response, recoverResponse)
	default:
		log.Fatal("Type mismatch")
	}


	responses := make([]*vss.Response, 0)
	responses = append(responses, response, response)

	switch content := PackResponses(responses).(type) {
	case *internalMsg.Responses:
		recoverResponses := UnpackResponses(content)
		require.Equal(t, responses, recoverResponses)
	default:
		log.Fatal("Type mismatch")
	}


	justification := &vss.Justification{
		SessionID:[]byte("SessionID"),
		Index:0,
		Deal:&vss.Deal{
			SessionID:[]byte("SessionID"),
		},
		Signature:[]byte("Signature"),
	}

	switch content := PackJustification(justification).(type) {
	case *internalMsg.Justification:
		recoverJustification := UnpackJustification(suite, content)
		require.Equal(t, justification, recoverJustification)
	default:
		log.Fatal("Type mismatch")
	}

	index := uint32(1)
	queryId := "queryId"
	sigContent := []byte("content")
	sig := []byte("Signature")

	switch content := PackSignature(index, queryId, sigContent, sig).(type) {
	case *internalMsg.Signature:
		recoverIndex, recoverId, recoverContent, recoverSig:= UnpackSignature(content)
		require.Equal(t, index, recoverIndex)
		require.Equal(t, queryId, recoverId)
		require.Equal(t, sigContent, recoverContent)
		require.Equal(t, sig, recoverSig)
	default:
		log.Fatal("Type mismatch")
	}


	publicKey := suite.G2().Point()

	switch content := PackPublicKey(publicKey).(type) {
	case *internalMsg.PublicKey:
		recoverPublicKey := UnpackPublicKey(suite, content)
		require.Equal(t, publicKey, *recoverPublicKey)
	default:
		log.Fatal("Type mismatch")
	}


	publicKeys := []kyber.Point{suite.G2().Point(), suite.G2().Point()}

	switch content := PackPublicKeys(publicKeys).(type) {
	case *internalMsg.PublicKeys:
		recoverPublicKeys := UnpackPublicKeys(suite, content)
		require.Equal(t, publicKeys, *recoverPublicKeys)
	default:
		log.Fatal("Type mismatch")
	}
}
