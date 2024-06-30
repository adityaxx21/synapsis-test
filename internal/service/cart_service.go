package service

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/config"

)


func StoreCart(user *domain.User, cart *domain.Cart) (*domain.Cart, error) {
	var item *domain.Item
	var store domain.Cart


	if err := config.DB.First(&item, cart.ItemID).Error; err != nil {
		return nil, err
	}
	
	result := config.DB.Where("user_id = ? AND item_id = ?", user.ID, cart.ItemID).First(&store)
	if result.Error == nil {
		// Update existing cart entry
		store.Total += cart.Total
		store.UpdatedAt = time.Now()
	} else if result.Error == gorm.ErrRecordNotFound {
		// Create new cart entry
		store = domain.Cart{
			UserID:    user.ID,
			ItemID:    cart.ItemID,
			Total:     cart.Total,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	} else {
		return nil, result.Error
	}

	// Validate Total against item quantity
	if store.Total > item.Stock {
		return  nil, fmt.Errorf("Total exceeds item quantity");
	}

	if err := config.DB.Save(&store).Error; err != nil {
		return nil, err
	}

	return &store, nil
}

func ListCart(user *domain.User, page int, limit int) ([]*domain.CartItem, error) {
	var cartItems []*domain.CartItem

	if err := config.DB.Model(&domain.Item{}).
		Select("items.id AS id, items.title, items.description, items.category, items.price, items.size, items.weight, items.stock, carts.user_id, carts.total").
		Joins("left join carts on carts.item_id = items.id").
		Where("carts.user_id = ?", user.ID).
		Scopes(domain.Paginate(page, limit)).
		Find(&cartItems).Error; err != nil {
			return nil, err
		}

	return cartItems, nil
}

func UpdateCart(id string, user_id uint, cart *domain.Cart)  (*domain.CartItem, error) {
	var item *domain.Item
	var cartItem *domain.CartItem

	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	if cart.Total > item.Stock {
		return  nil, fmt.Errorf("Total exceeds item quantity");
	}

	if err :=  config.DB.Model(&cart).Where("user_id = ?", user_id).Where("item_id = ?", id).Updates(cart).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&domain.Item{}).
		Select("items.id AS id, items.title, items.description, items.category, items.price, items.size, items.weight, items.stock, carts.user_id, carts.total").
		Joins("left join carts on carts.item_id = items.id").
		Where("carts.user_id = ?", user_id).
		First(&cartItem).Error; err != nil {
			return nil, err
		}

	return cartItem, nil
}

func DeleteCart(user_id uint, item_id string) error {
	var cart domain.Cart

	if err := config.DB.Where("user_id = ?", user_id).Where("item_id = ?", item_id).Delete(&cart).Error; err != nil {
		return err
	}

	return nil
}