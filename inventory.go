package zaico 

import (
	"context"
	"fmt"
	"net/url"
)

// Inventory 在庫データ
type Inventory struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Quantity  string `json:"quantity"`
	Unit      string  `json:"unit"`
	Category  string  `json:"category"`
	State     string  `json:"state"`
	Place     string  `json:"place"`
	Etc       string  `json:"etc"`
	GroupTag  string  `json:"group_tag"`
	Code      string  `json:"code"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// InventoryListOptions 在庫データ一覧取得の検索オプション
type InventoryListOptions struct {
	Title                   string `url:"title,omitempty"`
	Category                string `url:"category,omitempty"`
	Place                   string `url:"place,omitempty"`
	Code                    string `url:"code,omitempty"`
	OptionalAttributesName  string `url:"optional_attributes_name,omitempty"`
	OptionalAttributesValue string `url:"optional_attributes_value,omitempty"`
	Page                    int    `url:"page,omitempty"`
}

// InventoryService 在庫データの操作を行うサービス
type InventoryService struct {
	client *Client
}

// List 在庫データの一覧を取得します
func (s *InventoryService) List(ctx context.Context, opts *InventoryListOptions) ([]Inventory, error) {
	path := "./inventories"
	if opts != nil {
		params, err := url.ParseQuery("")
		if err != nil {
			return nil, err
		}
		if opts.Title != "" {
			params.Add("title", opts.Title)
		}
		if opts.Category != "" {
			params.Add("category", opts.Category)
		}
		if opts.Place != "" {
			params.Add("place", opts.Place)
		}
		if opts.Code != "" {
			params.Add("code", opts.Code)
		}
		if opts.OptionalAttributesName != "" {
			params.Add("optional_attributes_name", opts.OptionalAttributesName)
		}
		if opts.OptionalAttributesValue != "" {
			params.Add("optional_attributes_value", opts.OptionalAttributesValue)
		}
		if opts.Page > 0 {
			params.Add("page", fmt.Sprintf("%d", opts.Page))
		}
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var inventories []Inventory
	_, err = s.client.Do(ctx, req, &inventories)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

// Get 在庫データを個別取得します
func (s *InventoryService) Get(ctx context.Context, id int) (*Inventory, error) {
	path := fmt.Sprintf("./inventories/%d", id)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var inventory Inventory
	_, err = s.client.Do(ctx, req, &inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

// Create 在庫データを作成します
func (s *InventoryService) Create(ctx context.Context, inventory *Inventory) (*Inventory, error) {
	req, err := s.client.NewRequest("POST", "./inventories", inventory)
	if err != nil {
		return nil, err
	}

	var created Inventory
	_, err = s.client.Do(ctx, req, &created)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

// Update 在庫データを更新します
func (s *InventoryService) Update(ctx context.Context, id int, inventory *Inventory) (*Inventory, error) {
	path := fmt.Sprintf("./inventories/%d", id)
	req, err := s.client.NewRequest("PUT", path, inventory)
	if err != nil {
		return nil, err
	}

	var updated Inventory
	_, err = s.client.Do(ctx, req, &updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete 在庫データを削除します
func (s *InventoryService) Delete(ctx context.Context, id int) error {
	path := fmt.Sprintf("./inventories/%d", id)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	return err
}
