package vss_p2p

import (
	"log"

	"github.com/DOSNetwork/core/examples/vss_p2p/pb"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/dedis/kyber"
)

func packEncryptedDeal(encryptedDeal *vss.EncryptedDeal) *msg.EncryptedDeal {
	return &msg.EncryptedDeal{
		DHKey: encryptedDeal.DHKey,
		Signature: encryptedDeal.Signature,
		Nonce: encryptedDeal.Nonce,
		Cipher: encryptedDeal.Cipher,
	}
}

func unpackEncryptedDeal(encryptedDeal *msg.EncryptedDeal) *vss.EncryptedDeal {
	return &vss.EncryptedDeal{
		DHKey: encryptedDeal.GetDHKey(),
		Signature: encryptedDeal.GetSignature(),
		Nonce: encryptedDeal.GetNonce(),
		Cipher: encryptedDeal.GetCipher(),
	}
}

func packResonse (response *vss.Response) *msg.Response {
	return &msg.Response{
		SessionID: response.SessionID,
		Index: response.Index,
		Status: response.Status,
		Signature: response.Signature,
	}
}

func unpackResonse (response *msg.Response) *vss.Response {
	return &vss.Response{
		SessionID: response.GetSessionID(),
		Index: response.GetIndex(),
		Status: response.GetStatus(),
		Signature: response.GetSignature(),
	}
}

func packJustification(justification *vss.Justification) *msg.Justification {
	dealBytes, err := justification.Deal.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	return &msg.Justification{
		SessionID:justification.SessionID,
		Index:justification.Index,
		Deal:dealBytes,
		Signature:justification.Signature,
	}
}

func unpackJustification(s vss.Suite, justification *msg.Justification) *vss.Justification {
	deal := &vss.Deal{}
	err := deal.UnmarshalBinary(s, justification.GetDeal())
	if err != nil {
		log.Fatal(err)
	}

	return &vss.Justification{
		SessionID:justification.GetSessionID(),
		Index:justification.GetIndex(),
		Deal:deal,
		Signature:justification.GetSignature(),
	}
}

func packPublicKey(publicKey kyber.Point) *msg.PublicKey {
	publicKeyBytes, err := publicKey.MarshalBinary()

	if err != nil {
		log.Fatal(err)
	}

	return &msg.PublicKey{
		PublicKey: publicKeyBytes,
	}
}

func unpackPublicKey(s vss.Suite, publicKeyBytes *msg.PublicKey) kyber.Point {
	publicKey := s.G2().Point()
	err := publicKey.UnmarshalBinary(publicKeyBytes.GetPublicKey())
	if err != nil {
		log.Fatal(err)
	}

	return publicKey
}

func packPublicKeys(publicKeys []kyber.Point) *msg.PublicKeys {

	publicKeysBytes := make([][]byte, 0)
	for _, publicKey := range publicKeys {
		publicKeyBytes, err := publicKey.MarshalBinary()
		if err != nil {
			log.Fatal(err)
		}

		publicKeysBytes = append(publicKeysBytes, publicKeyBytes)
	}

	return &msg.PublicKeys{
		PublicKey:publicKeysBytes,
	}
}

func unpackPublicKeys(s vss.Suite, publicKeysBytes *msg.PublicKeys) []kyber.Point {

	publicKeys := make([]kyber.Point, 0)
	for _, publicKeyBytes := range publicKeysBytes.GetPublicKey() {
		publicKey := s.G2().Point()
		err := publicKey.UnmarshalBinary(publicKeyBytes)
		if err != nil {
			log.Fatal(err)
		}

		publicKeys = append(publicKeys, publicKey)
	}

	return publicKeys
}