/*
package apisrv provides an implementation of the gRPC server defined in ../../../api/protobuf-spec/mmlogic.proto.
Most of the documentation for what these calls should do is in that file!

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

// Intersection returns the intersection of two sets.
func Intersection(a []string, b []string) (out []string) {

	small, large := a, b
	if len(small) > len(large) {
		small, large = large, small
	}

	hash := make(map[string]struct{}, len(small))
	out = make([]string, 0, len(small))

	for _, v := range small {
		hash[v] = struct{}{}
	}

	// Iterate through large set, adding items found in small set.
	for _, v := range large {
		if _, found := hash[v]; found {
			out = append(out, v)
		}
	}

	return out

}

// Union returns the union of two sets.
func Union(a []string, b []string) (out []string) {

	small, large := a, b
	if len(small) > len(large) {
		small, large = large, small
	}

	hash := make(map[string]struct{}, len(small))
	out = make([]string, 0, len(small)+len(large))

	for _, v := range small {
		hash[v] = struct{}{}
	}

	// Iterate through large set, adding items not found in small set.
	out = append(out, small...)
	for _, v := range large {
		if _, found := hash[v]; !found {
			out = append(out, v)
		}
	}

	return out

}

// Difference returns the items in the first argument that are not in the
// second (set 'a' - set 'b')
func Difference(a []string, b []string) (out []string) {

	hash := make(map[string]struct{}, len(b))
	out = make([]string, 0, len(a))

	for _, v := range b {
		hash[v] = struct{}{}
	}

	// Iterate through a, adding items not found in b.
	for _, v := range a {
		if _, found := hash[v]; !found {
			out = append(out, v)
		}
	}

	return out
}
