/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package remoteapp

import "fmt"

type chunkState int

const (
	chunkState_Invalid chunkState = iota
	chunkState_Active
	chunkState_Local
	chunkState_Uploading
	chunkState_UploadError
	chunkState_Cleaning
	chunkState_Remote
	chunkState_Downloading
	chunkState_DownloadError
)

var chunkStateNames = []string{
	"Invalid",
	"Active",
	"Local",
	"Uploading",
	"UploadError",
	"Cleaning",
	"Remote",
	"Downloading",
	"DownloadError",
}

func (s chunkState) String() string {
	if s < 0 || int(s) >= len(chunkStateNames) {
		return fmt.Sprintf("chunkState(%d)", s)
	}
	return chunkStateNames[s]
}
