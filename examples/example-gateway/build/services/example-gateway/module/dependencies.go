// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package module

import (
	barendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"
	bazendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"
	contactsendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"
	googlenowendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"
	multiendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi"
	panicendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic"
	baztchannelendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz"
	panictchannelendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/panic"
	quuxendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/quux"
	withexceptionsendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/withexceptions"

	zanzibar "github.com/uber/zanzibar/runtime"
)

// Dependencies contains dependencies for the example-gateway service module
type Dependencies struct {
	Default  *zanzibar.DefaultDependencies
	Endpoint *EndpointDependencies
}

// EndpointDependencies contains endpoint dependencies
type EndpointDependencies struct {
	Bar            barendpointgenerated.Endpoint
	Baz            bazendpointgenerated.Endpoint
	Contacts       contactsendpointgenerated.Endpoint
	Googlenow      googlenowendpointgenerated.Endpoint
	Multi          multiendpointgenerated.Endpoint
	Panic          panicendpointgenerated.Endpoint
	BazTChannel    baztchannelendpointgenerated.Endpoint
	PanicTChannel  panictchannelendpointgenerated.Endpoint
	Quux           quuxendpointgenerated.Endpoint
	Withexceptions withexceptionsendpointgenerated.Endpoint
}
