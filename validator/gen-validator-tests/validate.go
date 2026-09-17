//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"encoding/json"
	"encoding/xml"

	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
	jsonparser "katydid.org.za/go/parser-go-json/json"
	xmlparser "katydid.org.za/go/parser-go-xml/xml"
	"katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/validator-go/validator/ast"
	"katydid.org.za/go/validator-go/validator/combinator"
)

type Validator struct {
	Name       string
	CodecName  string
	Grammar    *ast.Grammar
	Expected   bool
	Bytes      []byte
	SchemaName string
	Extension  string
}

var Validators = []Validator{}

var duplicates = map[string]struct{}{}

func checkDuplicates(name, codecName string) {
	n := name + "." + codecName
	if _, ok := duplicates[n]; ok {
		panic("duplicate validator: " + n)
	}
	duplicates[n] = struct{}{}
}

type ProtoMessage interface {
	proto.Message
	Description() *descriptor.FileDescriptorSet
}

func ValidateProtoEtc(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	ValidateReflect(name, grammar, m, expected)
	ValidateJson(name, grammar, m, expected)
	ValidateProto(name, grammar, m, expected)
}

// ValidateJsonProto only handles Json and Proto, since Reflect's fields are reordered when unmarshaled from JSON.
func ValidateJsonProto(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	ValidateJson(name, grammar, m, expected)
	ValidateProto(name, grammar, m, expected)
}

func ValidateProto(name string, g combinator.G, m ProtoMessage, expected bool) {
	schemaName := registerProto(m)
	checkDuplicates(name, "pb")
	Validators = append(Validators, Validator{
		Name:       name,
		CodecName:  "pb",
		Grammar:    g.Grammar(),
		Expected:   expected,
		Bytes:      must(proto.Marshal(m)),
		SchemaName: schemaName,
		Extension:  schemaName + ".pb",
	})
}

func ValidateJsonString(name string, g combinator.G, s string, expected bool) {
	checkDuplicates(name, "json")
	ValidateJSONHedge(name, g, s, expected)
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "json",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     []byte(s),
		Extension: "json",
	})
}

func ValidateJson(name string, g combinator.G, m any, expected bool) {
	checkDuplicates(name, "json")
	ValidateJSONHedge(name, g, string(must(json.MarshalIndent(m, "", "\t"))), expected)
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "json",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     must(json.MarshalIndent(m, "", "\t")),
		Extension: "json",
	})
}

func ValidateReflect(name string, g combinator.G, m any, expected bool) {
	checkDuplicates(name, "goreflect")
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "goreflect",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     must(json.MarshalIndent(m, "", "\t")),
		Extension: "goreflect",
	})
}

func ValidateXMLString(name string, g combinator.G, s string, expected bool) {
	checkDuplicates(name, "xml")
	ValidateXMLHedge(name, g, s, expected)
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "xml",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     []byte(s),
		Extension: "xml",
	})
}

func ValidateXML(name string, g combinator.G, m any, expected bool) {
	checkDuplicates(name, "xml")
	ValidateXMLHedge(name, g, string(must(xml.MarshalIndent(m, "", "\t"))), expected)
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "xml",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     must(xml.MarshalIndent(m, "", "\t")),
		Extension: "xml",
	})
}

func ValidateJSONHedge(name string, g combinator.G, s string, expected bool) {
	checkDuplicates(name, "hedge")
	p := jsonparser.NewParser()
	p.Init([]byte(s))
	h, err := hedge.ParseInto(p)
	if err != nil {
		panic(err)
	}
	data := must(json.Marshal(h))
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "hedge",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     data,
		Extension: "hedge",
	})
}

func ValidateXMLHedge(name string, g combinator.G, s string, expected bool) {
	checkDuplicates(name, "hedge")
	p := xmlparser.NewParser()
	p.Init([]byte(s))
	h, err := hedge.ParseInto(p)
	if err != nil {
		panic(err)
	}
	data := must(json.Marshal(h))
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "hedge",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     data,
		Extension: "hedge",
	})
}

func must[A any](bs A, err error) A {
	if err != nil {
		panic(err)
	}
	return bs
}
