package database

import (
	"fmt"
	"log"

	"github.com/IvanAbramovichWork/family-inventory-app/app/config"
	"github.com/IvanAbramovichWork/family-inventory-app/app/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(cfg *config.Config) *sqlx.DB {
	var db *sqlx.DB
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func createFamily(db *sqlx.DB, family *models.Family) error {
	_, err := db.NamedExec(`INSERT INTO family (name, created_at, updated_at) VALUES (:name, :created_at, :updated_at)`, map[string]interface{}{
		"name":       family.Name,
		"created_at": family.CreatedAt,
		"updated_at": family.UpdatedAt,
	})

	if err != nil {
		log.Printf("Error inserting family: %v", err)
		return err
	}

	return nil
}

func createFamilyMember(db *sqlx.DB, familyMember *models.FamilyMember) error {
	_, err := db.NamedExec(`INSERT INTO family_members (family_id, user_id, role) VALUES (:family_id, :user_id, :role)`, map[string]interface{}{
		"family_id": familyMember.FamilyId,
		"user_id":   familyMember.UserId,
		"role":      familyMember.Role,
	})

	if err != nil {
		log.Printf("Error inserting family member: %v", err)
		return err
	}

	return nil
}

func insertInventory(db *sqlx.DB, inventory *models.Inventory) error {
	_, err := db.NamedExec(`INSERT INTO inventory (family_id, product_id, quantity, expiration_date, created_at, updated_at) VALUES (:family_id, :product_id, :quantity, :expiration_date, :created_at, :updated_at)`, map[string]interface{}{
		"family_id":       inventory.FamilyId,
		"product_id":      inventory.ProductId,
		"quantity":        inventory.Quantity,
		"expiration_date": inventory.ExpirationDate,
		"created_at":      inventory.CreatedAt,
		"updated_at":      inventory.UpdatedAt,
	})

	if err != nil {
		log.Printf("Error inserting inventory: %v", err)
		return err

	}

	return nil
}

func createProduct(db *sqlx.DB, product *models.Product) error {
	_, err := db.NamedExec(`INSERT INTO products (name, category, description, photo_url, barcode, created_at, updated_at) VALUES (:name, :category, :description, :photo_url, :barcode, :created_at, :updated_at)`, map[string]interface{}{
		"name":        product.Name,
		"category":    product.Category,
		"description": product.Description,
		"photo_url":   product.Photo,
		"barcode":     product.Barcode,
		"created_at":  product.CreatedAt,
		"updated_at":  product.UpdatedAt,
	})

	if err != nil {
		log.Printf("Error inserting product: %v", err)
		return err

	}

	return nil
}

func createUser(db *sqlx.DB, user models.User) error {
	query := `INSERT INTO users (name, email, password_hash, role) VALUES (:name, :email, :password, :role)`
	_, err := db.NamedExec(query, user)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}

func getUserByID(db *sqlx.DB, userID int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.Get(&user, query, userID)
	return &user, err
}
