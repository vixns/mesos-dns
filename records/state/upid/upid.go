/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package upid

import (
	"fmt"
	"net"
	"strings"
)

// UPID is a equivalent of the UPID in libprocess.
type UPID struct {
	ID   string
	Host string
	Port string
}

// Parse parses the UPID from the input string.
func Parse(input string) (*UPID, error) {
	var err error
	upid := new(UPID)

	splits := strings.Split(input, "@")
	if len(splits) != 2 {
		return nil, fmt.Errorf("Expect one `@' in the input")
	}
	upid.ID = splits[0]

	// Using network "tcp" allows us to resolve ipv4 and ipv6
	if _, err = net.ResolveTCPAddr("tcp", splits[1]); err != nil {
		return nil, err
	}
	upid.Host, upid.Port, err = net.SplitHostPort(splits[1])
	return upid, err
}
