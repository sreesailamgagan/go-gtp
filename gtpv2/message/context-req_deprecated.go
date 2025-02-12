// Copyright 2019-2022 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import "log"

// Serialize serializes ContextRequest into bytes.
//
// Deprecated: use ContextRequest.Marshal instead.
func (c *ContextRequest) Serialize() ([]byte, error) {
	log.Println("ContextRequest.Serialize is deprecated. use ContextRequest.Marshal instead")
	return c.Marshal()
}

// SerializeTo serializes ContextRequest into bytes given as b.
//
// Deprecated: use ContextRequest.MarshalTo instead.
func (c *ContextRequest) SerializeTo(b []byte) error {
	log.Println("ContextRequest.SerializeTo is deprecated. use ContextRequest.MarshalTo instead")
	return c.MarshalTo(b)
}

// DecodeContextRequest decodes bytes as ContextRequest.
//
// Deprecated: use ParseContextRequest instead.
func DecodeContextRequest(b []byte) (*ContextRequest, error) {
	log.Println("DecodeContextRequest is deprecated. use ParseContextRequest instead")
	return ParseContextRequest(b)
}

// DecodeFromBytes decodes bytes as ContextRequest.
//
// Deprecated: use ContextRequest.UnmarshalBinary instead.
func (c *ContextRequest) DecodeFromBytes(b []byte) error {
	log.Println("ContextRequest.DecodeFromBytes is deprecated. use ContextRequest.UnmarshalBinary instead")
	return c.UnmarshalBinary(b)
}

// Len returns the actual length of ContextRequest.
//
// Deprecated: use ContextRequest.MarshalLen instead.
func (c *ContextRequest) Len() int {
	log.Println("ContextRequest.Len is deprecated. use ContextRequest.MarshalLen instead")
	return c.MarshalLen()
}
