// Copyright 2020 Luca Stasio. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package core implements core components of gogitsu lib.
//
// controller.go defines the Controller interface and the
// base controller implementation struct from which derive concrete
// controllers into apps.
package core

import "github.com/jinzhu/gorm"

// Repository is the contract interface for a db repository.
// This interface can be implemented with any db library.
// According to the specific db library, the DB type will be
// the library db type.
// A.e. for a gorm implementation it will be a *gorm.DB, or
// for a database/sql implementation it will be a *sql.DB.
type Repository interface {
	DB() interface{}
	SetDB(db interface{})
	InTransaction(tf TxFunction) error
}

// TxFunction is a function to be executed in a transaction.
// Depending on the repository implementation, the tx param
// will be the transaction object specific for the db library
// used in the implementation.
type TxFunction func(tx interface{}) error

// GormRepository is the base repository implemented with gorm.
// This should not be instantiated but composed in your
// concrete model repository.
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository returns a new gorm repository instance.
// This is useful in struct composition to create specific repository
// into apps.
// You can compose the GormRepository struct into your repository
// and use the NewGormRepository func to initializ the composition.
func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db}
}

// DB returns the internal, lib specific, db manager pointer.
func (r *GormRepository) DB() interface{} {
	return r.db
}

// SetDB sets into the repository the lib specific db manager pointer.
// The input db param is casted to a *gorm.DB.
func (r *GormRepository) SetDB(db interface{}) {
	r.db = db.(*gorm.DB)
}

// InTransaction executes the tf function in a db transaction.
// The tf function param will receive a tx instance that will
// have to be casted to a *gorm.DB instance.
// All of the db operation in the tf function will have to be
// executed against the tx passed in. Then, here, the transaction
// will be committed or rolled back for you.
func (r *GormRepository) InTransaction(tf TxFunction) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tf(tx); err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return err
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
