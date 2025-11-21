package postgres

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// WithTx executes a function within a database transaction
// If the function returns an error, the transaction is rolled back
// Otherwise, the transaction is committed
func WithTx(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %w", err, rbErr)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// WithTxContext executes a function within a database transaction with context support
func WithTxContext(ctx context.Context, db *gorm.DB, fn func(context.Context, *gorm.DB) error) error {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after rollback
		}
	}()

	if err := fn(ctx, tx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %w", err, rbErr)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// SavePoint creates a savepoint within a transaction
func SavePoint(tx *gorm.DB, name string) error {
	return tx.Exec(fmt.Sprintf("SAVEPOINT %s", name)).Error
}

// RollbackTo rolls back to a savepoint
func RollbackTo(tx *gorm.DB, name string) error {
	return tx.Exec(fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", name)).Error
}

// ReleaseSavepoint releases a savepoint
func ReleaseSavepoint(tx *gorm.DB, name string) error {
	return tx.Exec(fmt.Sprintf("RELEASE SAVEPOINT %s", name)).Error
}

// InTx checks if the DB instance is in a transaction
func InTx(db *gorm.DB) bool {
	committer, ok := db.Statement.ConnPool.(gorm.TxCommitter)
	return ok && committer != nil
}

// TxFunc is a function that operates within a transaction
type TxFunc func(*gorm.DB) error

// TxFuncContext is a function that operates within a transaction with context
type TxFuncContext func(context.Context, *gorm.DB) error

// RunInTxIfNotInTx executes the function in a transaction only if not already in one
func RunInTxIfNotInTx(db *gorm.DB, fn TxFunc) error {
	if InTx(db) {
		// Already in transaction, just execute
		return fn(db)
	}
	// Not in transaction, create one
	return WithTx(db, fn)
}

// RunInTxIfNotInTxContext executes the function in a transaction only if not already in one (with context)
func RunInTxIfNotInTxContext(ctx context.Context, db *gorm.DB, fn TxFuncContext) error {
	if InTx(db) {
		// Already in transaction, just execute
		return fn(ctx, db)
	}
	// Not in transaction, create one
	return WithTxContext(ctx, db, fn)
}