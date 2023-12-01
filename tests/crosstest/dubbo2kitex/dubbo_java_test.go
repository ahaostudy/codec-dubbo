/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dubbo2kitex

import (
	"testing"

	"github.com/kitex-contrib/codec-dubbo/tests/util"
)

func TestDubboJava(t *testing.T) {
	util.RunAndTestDubboJavaClient(t, "../../dubbo-java", "org.apache.dubbo.tests.client.Application",
		nil,
		[]string{
			// comment lines mean dubbo-java can not support
      "EchoBool":     false,
      "EchoByte":     false,
      "EchoInt16":    false,
      "EchoInt32":    false,
      "EchoInt64":    false,
      "EchoFloat":    false,
      "EchoDouble":   false,
      "EchoString":   false,
      "EchoBinary":   false,
      "EchoBoolList": false,
      //"EchoByteList":   false,
      //"EchoInt16List":  false,
      "EchoInt32List": false,
      "EchoInt64List": false,
      //"EchoFloatList":  false,
      "EchoDoubleList": false,
      "EchoStringList": false,
      // hessian2 can not support encoding [][]byte
      // dubbo-java can not support
      //"EchoBinaryList":   false,
      "EchoBool2BoolMap": false,
      //"EchoBool2ByteMap":   false,
      //"EchoBool2Int16Map":  false,
      "EchoBool2Int32Map": false,
      "EchoBool2Int64Map": false,
      //"EchoBool2FloatMap":  false,
      "EchoBool2DoubleMap": false,
      "EchoBool2StringMap": false,
      //"EchoBool2BinaryMap": false,
      "EchoMultiBool":   false,
      "EchoMultiByte":   false,
      "EchoMultiInt16":  false,
      "EchoMultiInt32":  false,
      "EchoMultiInt64":  false,
      "EchoMultiFloat":  false,
      "EchoMultiDouble": false,
      "EchoMultiString": false,
      "EchoMethodA":     false,
      "EchoMethodB":     false,
      "EchoMethodC":     false,
      "EchoMethodD":     false,
      //"EchoOptionalBool":           false,
      //"EchoOptionalInt32":          false,
      //"EchoOptionalString":         false,
      "EchoOptionalBoolList":            false,
      "EchoOptionalInt32List":           false,
      "EchoOptionalStringList":          false,
      "EchoOptionalBool2BoolMap":        false,
      "EchoOptionalBool2Int32Map":       false,
      "EchoOptionalBool2StringMap":      false,
      "EchoOptionalStruct":              false,
      "EchoOptionalMultiBoolRequest":    false,
      "EchoOptionalMultiInt32Request":   false,
      "EchoOptionalMultiStringRequest":  false,
      "EchoOptionalMultiBoolResponse":   false,
      "EchoOptionalMultiInt32Response":  false,
      "EchoOptionalMultiStringResponse": false,
		},
	)
}
