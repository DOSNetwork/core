package main

import (
	"log"

	"github.com/DOSNetwork/core/examples/vss_p2p/internal"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
)

func packEncryptedDeal(encryptedDeal *vss.EncryptedDeal) proto.Message {
	p := &internal.EncryptedDeal{
		DHKey:     encryptedDeal.DHKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
	return proto.Message(p)
}

func unpackEncryptedDeal(encryptedDeal *internal.EncryptedDeal) *vss.EncryptedDeal {
	return &vss.EncryptedDeal{
		DHKey:     encryptedDeal.GetDHKey(),
		Signature: encryptedDeal.GetSignature(),
		Nonce:     encryptedDeal.GetNonce(),
		Cipher:    encryptedDeal.GetCipher(),
	}
}
func packEncryptedDeals(encryptedDeals []*vss.EncryptedDeal) proto.Message {
	size := len(encryptedDeals)
	deals := make([]*internal.EncryptedDeal, size)
	for i := 0; i < size; i++ {
		encryptedDeal := encryptedDeals[i]
		p := &internal.EncryptedDeal{
			DHKey:     encryptedDeal.DHKey,
			Signature: encryptedDeal.Signature,
			Nonce:     encryptedDeal.Nonce,
			Cipher:    encryptedDeal.Cipher,
		}
		deals[i] = p
	}
	msg := &internal.EncryptedDeals{
		Deals: deals,
	}
	return proto.Message(msg)
}
func unpackEncryptedDeals(encryptedDeals *internal.EncryptedDeals) []*vss.EncryptedDeal {
	size := len(encryptedDeals.Deals)
	deals := make([]*vss.EncryptedDeal, size)
	for i := 0; i < size; i++ {
		encryptedDeal := encryptedDeals.Deals[i]
		p := &vss.EncryptedDeal{
			DHKey:     encryptedDeal.GetDHKey(),
			Signature: encryptedDeal.GetSignature(),
			Nonce:     encryptedDeal.GetNonce(),
			Cipher:    encryptedDeal.GetCipher(),
		}
		deals[i] = p
	}

	return deals
}
func packResonse(response *vss.Response) proto.Message {
	p := &internal.Response{
		SessionID: response.SessionID,
		Index:     response.Index,
		Status:    response.Status,
		Signature: response.Signature,
	}
	return proto.Message(p)
}

func unpackResonse(response *internal.Response) *vss.Response {
	return &vss.Response{
		SessionID: response.GetSessionID(),
		Index:     response.GetIndex(),
		Status:    response.GetStatus(),
		Signature: response.GetSignature(),
	}
}
func packResonses(responses []*vss.Response) proto.Message {
	size := len(responses)
	r := make([]*internal.Response, size)
	for i := 0; i < size; i++ {
		response := responses[i]
		p := &internal.Response{
			SessionID: response.SessionID,
			Index:     response.Index,
			Status:    response.Status,
			Signature: response.Signature,
		}
		r[i] = p
	}
	msg := &internal.Responses{
		Responses: r,
	}
	return proto.Message(msg)
}
func unpackResonses(resposes *internal.Responses) []*vss.Response {
	size := len(resposes.Responses)
	r := make([]*vss.Response, size)
	for i := 0; i < size; i++ {
		response := resposes.Responses[i]
		p := &vss.Response{
			SessionID: response.GetSessionID(),
			Index:     response.GetIndex(),
			Status:    response.GetStatus(),
			Signature: response.GetSignature(),
		}
		r[i] = p
	}

	return r
}
func packJustification(justification *vss.Justification) proto.Message {
	dealBytes, err := justification.Deal.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	p := &internal.Justification{
		SessionID: justification.SessionID,
		Index:     justification.Index,
		Deal:      dealBytes,
		Signature: justification.Signature,
	}
	return proto.Message(p)
}

func unpackJustification(s vss.Suite, justification *internal.Justification) *vss.Justification {
	deal := &vss.Deal{}
	err := deal.UnmarshalBinary(s, justification.GetDeal())
	if err != nil {
		log.Fatal(err)
	}

	return &vss.Justification{
		SessionID: justification.GetSessionID(),
		Index:     justification.GetIndex(),
		Deal:      deal,
		Signature: justification.GetSignature(),
	}
}

func packPublicKey(publicKey kyber.Point) proto.Message {
	publicKeyBytes, err := publicKey.MarshalBinary()

	if err != nil {
		log.Fatal(err)
	}

	p := &internal.PublicKey{
		PublicKey: publicKeyBytes,
	}
	return proto.Message(p)
}

func unpackPublicKey(s vss.Suite, publicKeyBytes *internal.PublicKey) kyber.Point {
	publicKey := s.G2().Point()
	err := publicKey.UnmarshalBinary(publicKeyBytes.GetPublicKey())
	if err != nil {
		log.Fatal(err)
	}

	return publicKey
}
func packSignature(sig []byte, index int) proto.Message {
	p := &internal.Signature{
		Index:     uint32(index),
		Signature: sig,
	}
	return proto.Message(p)
}

func packPublicKeys(publicKeys []kyber.Point) proto.Message {

	publicKeysBytes := make([][]byte, 0)
	for _, publicKey := range publicKeys {
		publicKeyBytes, err := publicKey.MarshalBinary()
		if err != nil {
			log.Fatal(err)
		}

		publicKeysBytes = append(publicKeysBytes, publicKeyBytes)
	}

	p := &internal.PublicKeys{
		PublicKey: publicKeysBytes,
	}
	return proto.Message(p)
}

func unpackPublicKeys(s vss.Suite, publicKeysBytes *internal.PublicKeys) []kyber.Point {

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
