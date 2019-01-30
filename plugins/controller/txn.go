// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/contiv/vpp/plugins/controller/api"
	scheduler_api "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
)

// kvSchedulerTxn implements Transaction interface for KVScheduler.
type kvSchedulerTxn struct {
	kvScheduler scheduler_api.KVScheduler

	groupings map[string][]string // key -> groupings

	// values set via Put or Delete
	values api.KeyValuePairs

	// injected by Controller to merge external with internal configuration
	merged api.KeyValuePairs
}

// newTransaction creates new transaction to be executed via KVScheduler.
func newTransaction(kvScheduler scheduler_api.KVScheduler) *kvSchedulerTxn {
	return &kvSchedulerTxn{
		kvScheduler: kvScheduler,
		groupings:   make(map[string][]string),
		values:      make(api.KeyValuePairs),
		merged:      make(api.KeyValuePairs),
	}
}

// Commit applies the requested transaction changes.
func (txn *kvSchedulerTxn) Commit(ctx context.Context) (seqNum uint64, err error) {
	schedTxn := txn.kvScheduler.StartNBTransaction()
	for key, value := range txn.values {
		if value != nil {
			// put
			schedTxn.SetValue(key, value, txn.getGroupings(key))
		} else {
			// delete
			schedTxn.SetValue(key, nil, txn.getGroupings(key))
		}
	}
	for key, value := range txn.merged {
		schedTxn.SetValue(key, value, txn.getGroupings(key))
	}
	return schedTxn.Commit(ctx)
}

// Put add request to the transaction to add or modify a value.
// <value> cannot be nil.
func (txn *kvSchedulerTxn) Put(key string, value proto.Message) {
	if value == nil {
		panic(fmt.Sprintf("Put nil value for key '%s'", key))
	}
	txn.values[key] = value
}

// Delete adds request to the transaction to delete an existing value.
func (txn *kvSchedulerTxn) Delete(key string) {
	txn.values[key] = nil
}

// SetGroupings allows to set potentially nested groups to cluster related
// key-value pairs.
func (txn *kvSchedulerTxn) SetGroupings(key string, groupings []string) {
	txn.groupings[key] = groupings
}

// Get is used to obtain value already prepared to be applied by this transaction.
// Until the transaction is committed, provided values can still be changed.
func (txn *kvSchedulerTxn) Get(key string) proto.Message {
	value, _ := txn.values[key]
	return value
}

func (txn *kvSchedulerTxn) getGroupings(key string) []string {
	groupings, hasAny := txn.groupings[key]
	if hasAny {
		return groupings
	}
	return nil
}