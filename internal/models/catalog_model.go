package models

import (
	"time"
)

type AttributeDataType string

const (
	AttrTypeText    AttributeDataType = "text"
	AttrTypeNumber  AttributeDataType = "number"
	AttrTypeBoolean AttributeDataType = "boolean"
	AttrTypeSelect  AttributeDataType = "select"
)

type ProductAuthorRole string

const (
	RoleMainAuthor ProductAuthorRole = "main_author"
	RoleTranslator ProductAuthorRole = "translator"
	RoleEditor     ProductAuthorRole = "editor"
	RoleIllustrator ProductAuthorRole = "illustrator"
)

type ProductType struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Slug        string    `gorm:"size:120;uniqueIndex;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	Icon        string    `gorm:"size:255" json:"icon"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Product struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	ProductTypeID uint         `gorm:"not null" json:"product_type_id"`
	ProductType   ProductType  `gorm:"foreignKey:ProductTypeID" json:"product_type"`
	Name          string       `gorm:"size:500;not null" json:"name"`
	Slug          string       `gorm:"size:550;uniqueIndex;not null" json:"slug"`
	Description   string       `gorm:"type:text" json:"description"`
	Price         float64      `gorm:"type:decimal(12,2);not null" json:"price"`
	OriginalPrice float64      `gorm:"type:decimal(12,2)" json:"original_price"`
	CoverImage    string       `gorm:"size:500" json:"cover_image"`
	IsActive      bool         `gorm:"not null;default:true" json:"is_active"`
	CreatedAt     time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	Images        []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Categories    []Category     `gorm:"many2many:product_categories;" json:"categories"`
	Authors       []Author       `gorm:"many2many:product_authors;" json:"authors"`
}

type ProductImage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	URL       string    `gorm:"size:500;not null" json:"url"`
	AltText   string    `gorm:"size:255" json:"alt_text"`
	SortOrder int8      `gorm:"not null;default:0" json:"sort_order"`
	IsPrimary bool      `gorm:"not null;default:false" json:"is_primary"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Category struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ParentID    *uint      `json:"parent_id"`
	Parent      *Category  `gorm:"foreignKey:ParentID" json:"parent"`
	Name        string     `gorm:"size:200;not null" json:"name"`
	Slug        string     `gorm:"size:220;uniqueIndex;not null" json:"slug"`
	Description string     `gorm:"type:text" json:"description"`
	Image       string     `gorm:"size:500" json:"image"`
	SortOrder   int        `json:"sort_order"`
	IsActive    bool       `gorm:"not null;default:true" json:"is_active"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Children    []Category `gorm:"foreignKey:ParentID" json:"children"`
}

type Author struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:300;not null" json:"name"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Avatar      string    `gorm:"size:500" json:"avatar"`
	Nationality string    `gorm:"size:100" json:"nationality"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AttributeDefinition struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ProductTypeID uint      `gorm:"not null" json:"product_type_id"`
	Name          string            `gorm:"size:200;not null" json:"name"`
	Key           string            `gorm:"size:100;not null" json:"key"`
	DataType      AttributeDataType `gorm:"type:varchar(20);not null" json:"data_type"` // text, number, boolean, select
	Unit          string            `gorm:"size:50" json:"unit"`
	IsRequired    bool      `gorm:"not null;default:false" json:"is_required"`
	IsFilterable  bool      `gorm:"not null;default:false" json:"is_filterable"`
	SortOrder     int       `json:"sort_order"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AttributeSelectOption struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AttributeID uint      `gorm:"not null" json:"attribute_id"`
	Label       string    `gorm:"size:200;not null" json:"label"`
	Value       string    `gorm:"size:100;not null" json:"value"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProductAttributeValue struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ProductID    uint      `gorm:"not null;uniqueIndex:idx_product_attribute" json:"product_id"`
	AttributeID  uint      `gorm:"not null;uniqueIndex:idx_product_attribute" json:"attribute_id"`
	ValueText    string    `gorm:"type:text" json:"value_text"`
	ValueNumber  float64   `gorm:"type:decimal(12,4)" json:"value_number"`
	ValueBoolean bool      `json:"value_boolean"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type VariantOptionType struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type VariantOptionValue struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OptionTypeID uint      `gorm:"not null" json:"option_type_id"`
	Value        string    `gorm:"size:200;not null" json:"value"`
	SortOrder    int       `json:"sort_order"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProductVariant struct {
	ID            uint                 `gorm:"primaryKey" json:"id"`
	ProductID     uint                 `gorm:"not null" json:"product_id"`
	SKU           string               `gorm:"size:100;uniqueIndex;not null" json:"sku"`
	PriceOverride float64              `gorm:"type:decimal(12,2)" json:"price_override"`
	ImageID       *uint                `json:"image_id"`
	IsActive      bool                 `gorm:"not null;default:true" json:"is_active"`
	CreatedAt     time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
	OptionValues  []VariantOptionValue `gorm:"many2many:variant_option_assignments;" json:"option_values"`
}

type Bundle struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:300;not null" json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	BundlePrice float64      `gorm:"type:decimal(12,2);not null" json:"bundle_price"`
	IsActive    bool         `gorm:"not null;default:true" json:"is_active"`
	StartsAt    *time.Time   `json:"starts_at"`
	ExpiresAt   *time.Time   `json:"expires_at"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	Items       []BundleItem `gorm:"foreignKey:BundleID" json:"items"`
}

type BundleItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	BundleID  uint    `gorm:"not null" json:"bundle_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	VariantID *uint   `json:"variant_id"`
	Quantity  int     `gorm:"not null;default:1" json:"quantity"`
}

type ProductCategory struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}

type ProductAuthor struct {
	ProductID uint              `gorm:"primaryKey"`
	AuthorID  uint              `gorm:"primaryKey"`
	Role      ProductAuthorRole `gorm:"type:varchar(50);not null"` // main_author, translator, editor, illustrator
	SortOrder int8              `json:"sort_order"`
}

type VariantOptionAssignment struct {
	VariantID     uint `gorm:"primaryKey"`
	OptionValueID uint `gorm:"primaryKey"`
}
