// Copyright © 2015 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package semtech

import (
	"fmt"
	"github.com/thethingsnetwork/core"
	"github.com/thethingsnetwork/core/components"
	"github.com/thethingsnetwork/core/semtech"
	"net"
)

type semtechAckNacker struct {
	conn      chan udpMsg
	recipient core.Recipient
}

func (an semtechAckNacker) Ack(p core.Packet) error {
	txpk, err := components.ConvertToTXPK(p)
	if err != nil {
		return err
	}
	raw, err := semtech.Marshal(semtech.Packet{
		Version:    semtech.VERSION,
		Identifier: semtech.PULL_RESP,
		Payload:    &semtech.Payload{TXPK: &txpk},
	})

	addr, ok := an.recipient.Address.(*net.UDPAddr)
	if !ok {
		return fmt.Errorf("Recipient address was invalid. Expected UDPAddr but got: %v", an.recipient.Address)
	}
	an.conn <- udpMsg{raw: raw, addr: addr}
	return nil
}

func (an semtechAckNacker) Nack(p core.Packet) error {
	return nil
}