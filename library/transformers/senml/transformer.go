// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package senml

import (
	"github.com/jxncyjq/lib_stardust/core/errors"
	"github.com/jxncyjq/lib_stardust/library/messaging"
	"github.com/jxncyjq/lib_stardust/library/transformers"
)

const (
	// json represents SenML in cbor format content type.
	JsonDec = "application/senml+json"
	// CBOR represents SenML in CBOR format content type.
	CborDec = "application/senml+cbor"
)

var (
	errDecode    = errors.New("failed to decode senml", 705)
	errNormalize = errors.New("failed to normalize senml", 706)
)

var formats = map[string]Format{
	JsonDec: JSON,
	CborDec: CBOR,
}

type transformer struct {
	format Format
}

// New returns transformer service implementation for SenML messages.
func New(contentFormat string) transformers.Transformer {
	format, ok := formats[contentFormat]
	if !ok {
		format = formats[JsonDec]
	}

	return transformer{
		format: format,
	}
}

func (t transformer) Transform(msg messaging.Message) (interface{}, error) {
	raw, err := Decode(msg.Payload, t.format)
	if err != nil {
		return nil, errors.Wrap(errDecode, err)
	}

	normalized, err := Normalize(raw)
	if err != nil {
		return nil, errors.Wrap(errNormalize, err)
	}

	msgs := make([]Message, len(normalized.Records))
	for i, v := range normalized.Records {
		// Use reception timestamp if SenML messsage Time is missing
		t := v.Time
		if t == 0 {
			// Convert the Unix timestamp in nanoseconds to float64
			t = float64(msg.Created) / float64(1e9)
		}

		msgs[i] = Message{
			Channel:     msg.Channel,
			Subtopic:    msg.Subtopic,
			Publisher:   msg.Publisher,
			Protocol:    msg.Protocol,
			Name:        v.Name,
			Unit:        v.Unit,
			Time:        t,
			UpdateTime:  v.UpdateTime,
			Value:       v.Value,
			BoolValue:   v.BoolValue,
			DataValue:   v.DataValue,
			StringValue: v.StringValue,
			Sum:         v.Sum,
		}
	}

	return msgs, nil
}
