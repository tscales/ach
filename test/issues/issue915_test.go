// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package issues

import (
	"bytes"
	"testing"

	"github.com/moov-io/ach"
)

func TestIssue915(t *testing.T) {
	rdr := ach.NewReader(bytes.NewReader(input))

	file, _ := rdr.Read()

	if n := len(file.Batches); n != 12 {
		t.Errorf("found %d batches", n)
	}
	for i := range file.Batches {
		entries := file.Batches[i].GetEntries()
		if n := len(entries); n != 1 {
			t.Errorf("Batch[%d] has %d entries", i, n)
		}
	}
}

var (
	input = []byte(`101CUSTOMNREC 0610001042103111747J094101CUSTOMN                ABC BANK, N.A.
5200ABC COMPANY                         9876543210PPDREVERSAL  0421162103110701123443210000001
622061000104101123400076543210007654321123456789012345TEST CONSUMER           0021070000087794
820000000100101000690000000000000000076543219876543210                         123443210000001
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789PPDTESTCREDIT1604212103110701121214140000002
622061000104101123400076543211234567890111111111111111TEST CONSUMER           0021070000087796
820000000100101000690000000000000012345678909123456789                         121214140000002
5200ABC COMPANY                         9876543210PPDREVERSAL  0421162103110701123443210000003
622061000104101123400076543210007654321123456789012345TEST CONSUMER           1021070000087803
705XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX00010087803
820000000200101000690000000000000000076543219876543210                         123443210000003
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789PPDTESTCREDIT1604212103110701121214140000004
622061000104101123400076543211234567890111111111111111TEST CONSUMER           1021070000087806
705XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX00010087806
820000000200101000690000000000000012345678909123456789                         121214140000004
5200ABC COMPANY                         9876543210WEBREVERSAL  0421162103110701123443210000005
622061000104101123400076543210007654321123456789012345TEST CONSUMER         S 0021070000087814
820000000100101000690000000000000000076543219876543210                         123443210000005
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789WEBTESTCREDIT1604212103110701121214140000006
622061000104101123400076543211234567890111111111111111TEST CONSUMER         S 0021070000087816
820000000100101000690000000000000012345678909123456789                         121214140000006
5200ABC COMPANY                         9876543210PPDTEST DEBIT0421162103110701123443210000007
627061000104101123400076543210007654321123456789012345TEST CONSUMER           0021070000087792
820000000100101000690000076543210000000000009876543210                         123443210000007
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789PPDREVERSAL  1604212103110701121214140000008
627061000104101123400076543211234567890111111111111111TEST CONSUMER           0021070000087798
820000000100101000690012345678900000000000009123456789                         121214140000008
5200ABC COMPANY                         9876543210PPDTEST DEBIT0421162103110701123443210000009
627061000104101123400076543211234567890123456789012345TEST CONSUMER           1021070000087800
705XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX00010087800
820000000200101000690012345678900000000000009876543210                         123443210000009
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789PPDREVERSAL  1604212103110701121214140000010
627061000104101123400076543211234567890111111111111111TEST CONSUMER           1021070000087809
705XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX00010087809
820000000200101000690012345678900000000000009123456789                         121214140000010
5200ABC COMPANY                         9876543210WEBTEST DEBIT0421162103110701123443210000011
627061000104101123400076543210007654321123456789012345TEST CONSUMER         S 0021070000087812
820000000100101000690000076543210000000000009876543210                         123443210000011
5200XYZ INCORPORATEDCOMPANY DESCRETIONAR9123456789WEBREVERSAL  1604212103110701121214140000012
627061000104101123400076543211234567890111111111111111TEST CONSUMER         S 0021070000087818
820000000100101000690012345678900000000000009123456789                         121214140000012
9000012000005000000160121200828004953580202003726666633
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999`)
)