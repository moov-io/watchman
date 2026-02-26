// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"strings"
	"testing"
)

func TestReader_Read_success(t *testing.T) {
	// Build a minimal UN Consolidated List XML containing one individual and one entity
	xml := `<?xml version="1.0"?>
+<CONSOLIDATED_LIST>
+  <INDIVIDUALS>
+    <INDIVIDUAL>
+      <DATAID>id1</DATAID>
+      <FIRST_NAME>Jane</FIRST_NAME>
+      <SECOND_NAME>Doe</SECOND_NAME>
+      <THIRD_NAME></THIRD_NAME>
+      <FOURTH_NAME></FOURTH_NAME>
+    </INDIVIDUAL>
+  </INDIVIDUALS>
+  <ENTITIES>
+    <ENTITY>
+      <DATAID>id2</DATAID>
+      <FIRST_NAME>Acme Corp</FIRST_NAME>
+    </ENTITY>
+  </ENTITIES>
+</CONSOLIDATED_LIST>`

	reader := NewReader(strings.NewReader(xml))

	var inds []UNIndividual
	var ents []UNEntity

	err := reader.Read(
		func(p UNIndividual) { inds = append(inds, p) },
		func(e UNEntity) { ents = append(ents, e) },
	)
	if err != nil {
		t.Fatalf("unexpected error reading xml: %v", err)
	}

	if len(inds) != 1 {
		t.Fatalf("expected 1 individual, got %d", len(inds))
	}
	if inds[0].DataID != "id1" || inds[0].FirstName != "Jane" {
		t.Errorf("individual fields incorrect: %+v", inds[0])
	}

	if len(ents) != 1 {
		t.Fatalf("expected 1 entity, got %d", len(ents))
	}
	if ents[0].DataID != "id2" || ents[0].FirstName != "Acme Corp" {
		t.Errorf("entity fields incorrect: %+v", ents[0])
	}
}

func TestReader_Read_malformed(t *testing.T) {
	// truncated/invalid XML should return an error on Decode
	xml := `<CONSOLIDATED_LIST><INDIVIDUALS><INDIVIDUAL><DATAID>id1` +
		"</CONSOLIDATED_LIST>"

	reader := NewReader(strings.NewReader(xml))

	err := reader.Read(func(p UNIndividual) {}, func(e UNEntity) {})
	if err == nil {
		t.Fatal("expected error for malformed xml")
	}
}
