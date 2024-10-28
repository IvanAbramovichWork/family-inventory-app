package database

import (
	"fmt"
	"log"
	"time"

	"github.com/IvanAbramovichWork/family-inventory-app/app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Family struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type FamilyMember struct {
	Id       int    `db:"id"`
	FamilyId int    `db:"family_id"`
	UserId   int    `db:"user_id"`
	Role     string `db:"role"`
}

type Product struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Category    string    `db:"category"`
	Description string    `db:"description"`
	Photo       string    `db:"photo_url"`
	Barcode     string    `db:"barcode"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Inventory struct {
	Id             int       `db:"id"`
	FamilyId       int       `db:"family_id"`
	ProductId      int       `db:"product_id"`
	Quantity       int       `db:"quantity"`
	ExpirationDate time.Time `db:"expiration_date"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
type Transaction struct {
	Id        int       `db:"id"`
	FamilyId  int       `db:"family_id"`
	ProductId int       `db:"product_id"`
	UserId    int       `db:"user_id"`
	Action    string    `db:"action"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
}

type Notification struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	ProductId int       `db:"product_id"`
	Message   string    `db:"message"`
	SentAt    time.Time `db:"sent_at"`
	IsRead    bool      `db:"is_read"`
}

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

func createFamily(db *sqlx.DB, family *Family) error {
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

func createFamilyMember(db *sqlx.DB, familyMember *FamilyMember) error {
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

func insertInventory(db *sqlx.DB, inventory *Inventory) error {
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

func createProduct(db *sqlx.DB, product *Product) error {
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

func createUser(db *sqlx.DB, user User) error {
	query := `INSERT INTO users (name, email, password_hash, role) VALUES (:name, :email, :password, :role)`
	_, err := db.NamedExec(query, user)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}

func getUserByID(db *sqlx.DB, userID int) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.Get(&user, query, userID)
	return &user, err
}
