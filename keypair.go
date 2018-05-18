/* SPDX-License-Identifier: GPL-2.0
 *
 * Copyright (C) 2017-2018 Jason A. Donenfeld <Jason@zx2c4.com>. All Rights Reserved.
 * Copyright (C) 2017-2018 Mathias N. Hall-Andersen <mathias@hall-andersen.dk>.
 */

package main

import (
	"crypto/cipher"
	"sync"
	"time"
)

/* Due to limitations in Go and /x/crypto there is currently
 * no way to ensure that key material is securely ereased in memory.
 *
 * Since this may harm the forward secrecy property,
 * we plan to resolve this issue; whenever Go allows us to do so.
 */

type Keypair struct {
	sendNonce    uint64
	send         cipher.AEAD
	receive      cipher.AEAD
	replayFilter ReplayFilter
	isInitiator  bool
	created      time.Time
	localIndex   uint32
	remoteIndex  uint32
}

type Keypairs struct {
	mutex    sync.RWMutex
	current  *Keypair
	previous *Keypair
	next     *Keypair
}

func (kp *Keypairs) Current() *Keypair {
	kp.mutex.RLock()
	defer kp.mutex.RUnlock()
	return kp.current
}

func (device *Device) DeleteKeypair(key *Keypair) {
	if key != nil {
		device.indexTable.Delete(key.localIndex)
	}
}
