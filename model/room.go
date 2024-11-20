/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2024 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file room.go
 * @package model
 * @author Dr.NP <np@herewe.tech>
 * @since 11/20/2024
 */

package model

import (
	"context"
	"time"

	"github.com/go-sicky/sicky/driver"
	"github.com/go-sicky/sicky/logger"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Room struct {
	bun.BaseModel `bun:"table:room"`

	ID uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero" json:"-"`
}

func InitRoom(ctx context.Context, dropTable bool) error {
	if dropTable {
		// Drop exist table
		_, err := driver.DB.NewDropTable().Model((*Room)(nil)).Exec(ctx)
		if err != nil {
			logger.Logger.ErrorContext(ctx, err.Error())

			return err
		}
	}

	// Create table
	_, err := driver.DB.NewCreateTable().Model((*Room)(nil)).Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	// Indexes
	_, err = driver.DB.NewCreateIndex().
		Model((*Room)(nil)).
		Index("idx_room_crated_at").
		Column("created_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*Room)(nil)).
		Index("idx_room_updated_at").
		Column("updated_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*Room)(nil)).
		Index("idx_room_deleted_at").
		Column("deleted_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	return err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
