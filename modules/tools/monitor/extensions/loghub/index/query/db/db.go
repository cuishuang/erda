// Copyright (c) 2021 Terminus, Inc.
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

package db

import (
	"github.com/jinzhu/gorm"
)

// DB .
type DB struct {
	*gorm.DB
	LogDeployment        LogDeploymentDB
	LogServiceInstanceDB LogServiceInstanceDB
	LogInstanceDB        LogInstanceDB
}

// New .
func New(db *gorm.DB) *DB {
	return &DB{
		DB:                   db,
		LogDeployment:        LogDeploymentDB{db},
		LogServiceInstanceDB: LogServiceInstanceDB{db},
		LogInstanceDB:        LogInstanceDB{db},
	}
}

// Begin .
func (db *DB) Begin() *DB {
	return New(db.DB.Begin())
}