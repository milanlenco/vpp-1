// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/punt"
)

////////// type-safe key-value pair with metadata //////////

type PuntToHostKVWithMetadata struct {
	Key      string
	Value    *punt.ToHost
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type PuntToHostDescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, oldValue, newValue *punt.ToHost) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Add                func(key string, value *punt.ToHost) (metadata interface{}, err error)
	Delete             func(key string, value *punt.ToHost, metadata interface{}) error
	Modify             func(key string, oldValue, newValue *punt.ToHost, oldMetadata interface{}) (newMetadata interface{}, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *punt.ToHost, metadata interface{}) bool
	Update             func(key string, value *punt.ToHost, metadata interface{}) error
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *punt.ToHost) []Dependency
	DerivedValues      func(key string, value *punt.ToHost) []KeyValuePair
	Dump               func(correlate []PuntToHostKVWithMetadata) ([]PuntToHostKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type PuntToHostDescriptorAdapter struct {
	descriptor *PuntToHostDescriptor
}

func NewPuntToHostDescriptor(typedDescriptor *PuntToHostDescriptor) *KVDescriptor {
	adapter := &PuntToHostDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *PuntToHostDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castPuntToHostValue(key, oldValue)
	typedNewValue, err2 := castPuntToHostValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *PuntToHostDescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castPuntToHostValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *PuntToHostDescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castPuntToHostValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castPuntToHostValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castPuntToHostMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *PuntToHostDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castPuntToHostValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castPuntToHostMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *PuntToHostDescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castPuntToHostValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castPuntToHostValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castPuntToHostMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *PuntToHostDescriptorAdapter) Update(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castPuntToHostValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castPuntToHostMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Update(key, typedValue, typedMetadata)
}

func (da *PuntToHostDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castPuntToHostValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *PuntToHostDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castPuntToHostValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *PuntToHostDescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []PuntToHostKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castPuntToHostValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castPuntToHostMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			PuntToHostKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castPuntToHostValue(key string, value proto.Message) (*punt.ToHost, error) {
	typedValue, ok := value.(*punt.ToHost)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castPuntToHostMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
