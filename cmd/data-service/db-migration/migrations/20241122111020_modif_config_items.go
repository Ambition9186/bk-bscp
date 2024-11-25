/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package migrations

import (
	"gorm.io/gorm"

	"github.com/TencentBlueKing/bk-bscp/cmd/data-service/db-migration/migrator"
)

func init() {
	// add current migration to migrator
	migrator.GetMigrator().AddMigration(&migrator.Migration{
		Version: "20241122111020",
		Name:    "20241122111020_modif_config_items",
		Mode:    migrator.GormMode,
		Up:      mig20241122111020Up,
		Down:    mig20241122111020Down,
	})
}

// mig20241122111020Up for up migration
func mig20241122111020Up(tx *gorm.DB) error {
	// ConfigItems  : config_items
	type ConfigItems struct {
		FileState string `gorm:"column:file_state;type:varchar(20);default:'';NOT NULL"`
	}

	// ConfigItems add new column
	if !tx.Migrator().HasColumn(&ConfigItems{}, "file_state") {
		if err := tx.Migrator().AddColumn(&ConfigItems{}, "file_state"); err != nil {
			return err
		}
	}

	return nil
}

// mig20241122111020Down for down migration
func mig20241122111020Down(tx *gorm.DB) error {
	// ConfigItems  : config_items
	type ConfigItems struct {
		FileState string `gorm:"column:file_state;type:varchar(20);default:'';NOT NULL"`
	}

	// ConfigItems drop column
	if tx.Migrator().HasColumn(&ConfigItems{}, "file_state") {
		if err := tx.Migrator().DropColumn(&ConfigItems{}, "file_state"); err != nil {
			return err
		}
	}

	return nil
}
