package service

import (
	"fmt"
	"time"
	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/config"
	"gorm.io/gorm"
)


func CreateOrder(user *domain.User, transaction *domain.TransactionRequest) error {
	var item domain.Item

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&item, transaction.ItemID).Error; err != nil {
			return err
		}
	
		if transaction.Total > item.Stock {
			return fmt.Errorf("Item out of stock")
		}
	
		store := &domain.Transaction{
			UserID: user.ID,
			GrossAmount: transaction.GrossAmount,
			Status: transaction.Status,
			OrderType: transaction.OrderType,
			TransactionDate: time.Now(),
		}
	
		if err := tx.Create(&store).Error; err != nil {
			tx.Rollback()
			return err
		}

		transactionItems := []domain.TransactionItem{
			{TransactionID: int(store.ID), ItemID: transaction.ItemID, Total: transaction.Total},
		}

		if err := tx.Create(&transactionItems).Error; err != nil {
			tx.Rollback()
			return err
		}

	
		item.Stock = item.Stock - transaction.Total
	
		if err :=  tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateOrderCart(user *domain.User, transaction *domain.TransactionCartRequest) error {
	var items []domain.TransactionItem

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		store := &domain.Transaction{
			UserID: user.ID,
			GrossAmount: transaction.GrossAmount,
			Status: transaction.Status,
			OrderType: transaction.OrderType,
			TransactionDate: time.Now(),
		}

		if err := tx.Create(&store).Error; err != nil {
			tx.Rollback()
			return err
		}

		for _, d := range(transaction.Items) {
			var itemData domain.Item

			// Double check if item still exist
			if err := tx.First(&itemData, d.ItemID).Error; err != nil {
				return err
			}
		
			if d.Total > itemData.Stock {
				tx.Rollback()
				return fmt.Errorf("Item out of stock")
			}

			// Update stock of item when checkout
			itemData.Stock = itemData.Stock - d.Total

			if err :=  tx.Save(&itemData).Error; err != nil {
				tx.Rollback()
				return err
			}

			// Delete item on cart
			if err := tx.Where("item_id = ?", d.ItemID).Where("user_id = ?", user.ID).Delete(&domain.Cart{}).Error; err != nil {
				tx.Rollback()
				return err
			}

			// Append data to slice
			item := domain.TransactionItem{
				TransactionID: int(store.ID), 
				ItemID: d.ItemID, 
				Total: d.Total,
			}

			items = append(items, item)
		}

		transactionItems := &items

		if err := tx.Create(&transactionItems).Error; err != nil {
			tx.Rollback()
			return err
		}

		// don't forgot to delete cart

		return nil
	})

	if err != nil {
		return	err
	}

	return nil
}


func DetailOrder(user *domain.User, transactionID string) (*domain.Transaction, error) {
    var transaction domain.Transaction
    err := config.DB.Preload("Items").Where("user_id = ?",user.ID).First(&transaction, "id = ?", transactionID).Error
    if err != nil {
        return nil, err
    }
    return &transaction, nil
}

func ListOrder(user *domain.User, page int, limit int) (*[]domain.Transaction, error) {
	var transaction []domain.Transaction

    err := config.DB.Preload("Items").
		Where("user_id = ?",user.ID).
		Scopes(domain.Paginate(page, limit)).
		Find(&transaction).Error

    if err != nil {
        return nil, err
    }
    return &transaction, nil
}