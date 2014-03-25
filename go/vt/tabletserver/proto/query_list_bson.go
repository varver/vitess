// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/youtube/vitess/go/bson"
	"github.com/youtube/vitess/go/bytes2"
)

// MarshalBson bson-encodes QueryList.
func (queryList *QueryList) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	// []BoundQuery
	{
		bson.EncodePrefix(buf, bson.Array, "Queries")
		lenWriter := bson.NewLenWriter(buf)
		for _i, _v1 := range queryList.Queries {
			_v1.MarshalBson(buf, bson.Itoa(_i))
		}
		lenWriter.Close()
	}
	bson.EncodeInt64(buf, "SessionId", queryList.SessionId)
	bson.EncodeInt64(buf, "TransactionId", queryList.TransactionId)

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into QueryList.
func (queryList *QueryList) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for QueryList", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Queries":
			// []BoundQuery
			if kind != bson.Null {
				if kind != bson.Array {
					panic(bson.NewBsonError("unexpected kind %v for queryList.Queries", kind))
				}
				bson.Next(buf, 4)
				queryList.Queries = make([]BoundQuery, 0, 8)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					bson.SkipIndex(buf)
					var _v1 BoundQuery
					_v1.UnmarshalBson(buf, kind)
					queryList.Queries = append(queryList.Queries, _v1)
				}
			}
		case "SessionId":
			queryList.SessionId = bson.DecodeInt64(buf, kind)
		case "TransactionId":
			queryList.TransactionId = bson.DecodeInt64(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}