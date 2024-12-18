package helper

import (
	"log"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			log.Printf("Error during rollback: %v", rollbackErr)
		}
		log.Printf("Transaction error: %v", err)
		panic(err)
	} else {
		if commitErr := tx.Commit().Error; commitErr != nil {
			log.Printf("Error during commit: %v", commitErr)
			tx.Rollback()
			panic(commitErr)
		}
	}
}
