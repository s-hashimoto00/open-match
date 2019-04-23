/*
package set provides implementation of basic set operationg in golang.

Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package set

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/rs/xid"
)

var (
	a1 [5]string
	a2 [4]string
	i  [2]string
	u  [7]string
	d  [3]string
)

type stringOperation func([]string, []string) []string

func TestMain(m *testing.M) {

	a1 = [5]string{"a", "b", "c", "d", "e"}
	a2 = [4]string{"b", "c", "f", "g"}
	i = [2]string{"b", "c"}
	u = [7]string{"a", "b", "c", "d", "e", "f", "g"}
	d = [3]string{"a", "d", "e"}

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestStringOperations(t *testing.T) {
	t.Run("Difference: identical slices", testStringOperation(Difference, a1[:], a1[:], make([]string, 0)))
	t.Run("Difference: valid slices", testStringOperation(Difference, a1[:], a2[:], d[:]))
	t.Run("Difference: remove nil", testStringOperation(Difference, a1[:], nil, a1[:]))
	t.Run("Difference: remove from nil", testStringOperation(Difference, nil, a2[:], nil))
	t.Run("Union: valid slices", testStringOperation(Union, a1[:], a2[:], u[:]))
	t.Run("Union: nil first", testStringOperation(Union, nil, a2[:], a2[:]))
	t.Run("Union: nil second", testStringOperation(Union, a1[:], nil, a1[:]))
	t.Run("Intersection: valid slices", testStringOperation(Intersection, a1[:], a2[:], i[:]))
	t.Run("Intersection: nil first", testStringOperation(Intersection, a1[:], nil, nil))
	t.Run("Intersection: nil second", testStringOperation(Intersection, nil, a2[:], nil))
}

func testStringOperation(thisFunc stringOperation, in1 []string, in2 []string, outputExpected []string) func(*testing.T) {
	return func(t *testing.T) {
		t.Logf("func(%-15s , %-15s) = %v", fmt.Sprintf("%v", in1), fmt.Sprintf("%v", in2), outputExpected)
		outputActual := thisFunc(in1, in2)

		if len(outputActual) != len(outputExpected) {
			t.Errorf("Length of output string slice incorrect, got %v, want %v", len(outputActual), len(outputExpected))
		}
		for x := 0; x < len(outputActual); x++ {
			found := false
			for y := 0; y < len(outputExpected) && found != true; y++ {
				if outputActual[x] == outputExpected[y] {
					found = true
				}
			}
			if !found {
				t.Errorf("Element of output string slice incorrect, no %v in %v", outputActual[x], outputExpected)

			}

		}
	}
}

func makeBenchmarkData(total, size1, size2 int) ([]string, []string) {
	r := rand.New(rand.NewSource(1))

	ids := make([]string, 0, total)
	for i := 0; i < total; i++ {
		ids = append(ids, xid.New().String())
	}

	b1 := append([]string{}, ids[:size1]...)
	r.Shuffle(len(b1), func(i, j int) {
		b1[i], b1[j] = b1[j], b1[i]
	})

	b2 := append([]string{}, ids[size2:]...)
	r.Shuffle(len(b2), func(i, j int) {
		b2[i], b2[j] = b2[j], b2[i]
	})

	return b1, b2
}

func BenchmarkSet(b *testing.B) {
	s1, s2 := makeBenchmarkData(700, 500, 300)
	m1, m2 := makeBenchmarkData(7000, 5000, 3000)
	l1, l2 := makeBenchmarkData(70000, 50000, 30000)

	tbl := []struct {
		name string
		b1   []string
		b2   []string
	}{
		{
			name: "tiny",
			b1:   a1[:],
			b2:   a2[:],
		},
		{
			name: "small",
			b1:   s1,
			b2:   s2,
		},
		{
			name: "medium",
			b1:   m1,
			b2:   m2,
		},
		{
			name: "large",
			b1:   l1,
			b2:   l2,
		},
	}

	for _, v := range tbl {
		b.Run(fmt.Sprintf("Difference %s", v.name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Difference(v.b1, v.b2)
			}
		})
		b.Run(fmt.Sprintf("Union %s", v.name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Union(v.b1, v.b2)
			}
		})
		b.Run(fmt.Sprintf("Intersection %s", v.name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Intersection(v.b1, v.b2)
			}
		})
	}
}
