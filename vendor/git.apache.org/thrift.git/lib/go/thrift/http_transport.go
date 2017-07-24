/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE ***REMOVED***le
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this ***REMOVED***le
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this ***REMOVED***le except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * speci***REMOVED***c language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"strings"
)

// NewThriftHandlerFunc is a function that create a ready to use Apache Thrift Handler function
func NewThriftHandlerFunc(processor TProcessor,
	inPfactory, outPfactory TProtocolFactory) func(w http.ResponseWriter, r *http.Request) {

	return gz(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/x-thrift")

		transport := NewStreamTransport(r.Body, w)
		processor.Process(inPfactory.GetProtocol(transport), outPfactory.GetProtocol(transport))
	})
}

// NewThriftHandlerFunc2 is same as NewThriftHandlerFunc but requires a Context as its ***REMOVED***rst param.
func NewThriftHandlerFunc2(ctx context.Context, processor TProcessor2,
	inPfactory, outPfactory TProtocolFactory) func(w http.ResponseWriter, r *http.Request) {

	return gz(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/x-thrift")

		transport := NewStreamTransport(r.Body, w)
		processor.Process(ctx, inPfactory.GetProtocol(transport), outPfactory.GetProtocol(transport))
	})
}

// gz transparently compresses the HTTP response if the client supports it.
func gz(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler(gzw, r)
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
