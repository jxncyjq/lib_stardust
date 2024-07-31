// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package transformers

import "git.u-linke.com/ulink/commons/library/messaging"

// Transformer specifies API form Message transformer.
type Transformer interface {
	// Transform Mainflux message to any other format.
	Transform(msg messaging.Message) (interface{}, error)
}
