// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"context"

	"github.com/google/tink/go/tink"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// keyID uniquely identifies a keyset.
type keyID struct {
	instance    int64
	directoryID string
	appID       string
}

// KeySets implements storage.UserManagerTable in memory.
type KeySets struct {
	keysets map[keyID]*tink.KeysetHandle
}

// NewKeySets produces a fake implementation of storage.UserManagerTable.
func NewKeySets() *KeySets {
	return &KeySets{
		keysets: make(map[keyID]*tink.KeysetHandle),
	}
}

// Get returns the requested keyset.
func (k *KeySets) Get(ctx context.Context, instance int64, directoryID, appID string) (*tink.KeysetHandle, error) {
	ks, ok := k.keysets[keyID{
		instance:    instance,
		directoryID: directoryID,
		appID:       appID,
	}]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "KeySet %v/%v/%v not found", instance, directoryID, appID)
	}
	return ks, nil
}
