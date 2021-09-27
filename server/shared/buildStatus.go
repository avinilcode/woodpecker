// Copyright 2019 mhmxs.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shared

import (
	"time"

	"github.com/woodpecker-ci/woodpecker/server/model"
)

type UpdateBuildStore interface {
	UpdateBuild(*model.Build) error
}

func UpdateToStatusRunning(store UpdateBuildStore, build model.Build, started int64) (*model.Build, error) {
	build.Status = model.StatusRunning
	build.Started = started
	return &build, store.UpdateBuild(&build)
}

func UpdateToStatusPending(store UpdateBuildStore, build model.Build, reviewer string) (*model.Build, error) {
	build.Reviewer = reviewer
	build.Status = model.StatusPending
	build.Reviewed = time.Now().Unix()
	return &build, store.UpdateBuild(&build)
}

func UpdateToStatusDeclined(store UpdateBuildStore, build model.Build, reviewer string) (*model.Build, error) {
	build.Reviewer = reviewer
	build.Status = model.StatusDeclined
	build.Reviewed = time.Now().Unix()
	return &build, store.UpdateBuild(&build)
}

func UpdateStatusToDone(store UpdateBuildStore, build model.Build, status string, stopped int64) (*model.Build, error) {
	build.Status = status
	build.Finished = stopped
	return &build, store.UpdateBuild(&build)
}

func UpdateToStatusError(store UpdateBuildStore, build model.Build, err error) (*model.Build, error) {
	build.Error = err.Error()
	build.Status = model.StatusError
	build.Started = time.Now().Unix()
	build.Finished = build.Started
	return &build, store.UpdateBuild(&build)
}

func UpdateToStatusKilled(store UpdateBuildStore, build model.Build) (*model.Build, error) {
	build.Status = model.StatusKilled
	build.Finished = time.Now().Unix()
	return &build, store.UpdateBuild(&build)
}
