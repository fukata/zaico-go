package main 

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestInventoryService_List(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")

		inventories := []Inventory{
			{
				ID:        1,
				Title:     "テスト在庫1",
				Quantity:  10,
				Unit:      "個",
				Category:  "テストカテゴリ",
				State:     "新品",
				Place:     "倉庫A",
				CreatedAt: "2024-01-01T00:00:00Z",
				UpdatedAt: "2024-01-01T00:00:00Z",
			},
		}
		json.NewEncoder(w).Encode(inventories)
	})

	ctx := context.Background()
	inventories, err := client.Inventory.List(ctx, nil)
	if err != nil {
		t.Errorf("Inventory.List returned error: %v", err)
	}

	if len(inventories) != 1 {
		t.Errorf("Inventory.List returned %d inventories, want 1", len(inventories))
	}
}

func TestInventoryService_ListWithOptions(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")

		// クエリパラメータの検証
		if got := r.URL.Query().Get("title"); got != "テスト在庫" {
			t.Errorf("title query = %v, want %v", got, "テスト在庫")
		}
		if got := r.URL.Query().Get("category"); got != "テストカテゴリ" {
			t.Errorf("category query = %v, want %v", got, "テストカテゴリ")
		}
		if got := r.URL.Query().Get("place"); got != "倉庫A" {
			t.Errorf("place query = %v, want %v", got, "倉庫A")
		}
		if got := r.URL.Query().Get("code"); got != "123456" {
			t.Errorf("code query = %v, want %v", got, "123456")
		}
		if got := r.URL.Query().Get("optional_attributes_name"); got != "担当者" {
			t.Errorf("optional_attributes_name query = %v, want %v", got, "担当者")
		}
		if got := r.URL.Query().Get("optional_attributes_value"); got != "山田" {
			t.Errorf("optional_attributes_value query = %v, want %v", got, "山田")
		}
		if got := r.URL.Query().Get("page"); got != "2" {
			t.Errorf("page query = %v, want %v", got, "2")
		}

		inventories := []Inventory{
			{
				ID:        1,
				Title:     "テスト在庫1",
				Quantity:  10,
				Unit:      "個",
				Category:  "テストカテゴリ",
				State:     "新品",
				Place:     "倉庫A",
				CreatedAt: "2024-01-01T00:00:00Z",
				UpdatedAt: "2024-01-01T00:00:00Z",
			},
		}
		json.NewEncoder(w).Encode(inventories)
	})

	ctx := context.Background()
	opts := &InventoryListOptions{
		Title:                   "テスト在庫",
		Category:                "テストカテゴリ",
		Place:                   "倉庫A",
		Code:                    "123456",
		OptionalAttributesName:  "担当者",
		OptionalAttributesValue: "山田",
		Page:                    2,
	}

	inventories, err := client.Inventory.List(ctx, opts)
	if err != nil {
		t.Errorf("Inventory.List returned error: %v", err)
	}

	if len(inventories) != 1 {
		t.Errorf("Inventory.List returned %d inventories, want 1", len(inventories))
	}
}

func TestInventoryService_Get(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")

		inventory := Inventory{
			ID:        1,
			Title:     "テスト在庫1",
			Quantity:  10,
			Unit:      "個",
			Category:  "テストカテゴリ",
			State:     "新品",
			Place:     "倉庫A",
			CreatedAt: "2024-01-01T00:00:00Z",
			UpdatedAt: "2024-01-01T00:00:00Z",
		}
		json.NewEncoder(w).Encode(inventory)
	})

	ctx := context.Background()
	inventory, err := client.Inventory.Get(ctx, 1)
	if err != nil {
		t.Errorf("Inventory.Get returned error: %v", err)
	}

	if inventory.ID != 1 {
		t.Errorf("Inventory.Get returned ID %d, want 1", inventory.ID)
	}
}

func TestInventoryService_Create(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")

		var inventory Inventory
		json.NewDecoder(r.Body).Decode(&inventory)

		inventory.ID = 1
		inventory.CreatedAt = "2024-01-01T00:00:00Z"
		inventory.UpdatedAt = "2024-01-01T00:00:00Z"
		json.NewEncoder(w).Encode(inventory)
	})

	ctx := context.Background()
	newInventory := &Inventory{
		Title:    "テスト在庫1",
		Quantity: 10,
		Unit:     "個",
		Category: "テストカテゴリ",
		State:    "新品",
		Place:    "倉庫A",
	}

	created, err := client.Inventory.Create(ctx, newInventory)
	if err != nil {
		t.Errorf("Inventory.Create returned error: %v", err)
	}

	if created.ID != 1 {
		t.Errorf("Inventory.Create returned ID %d, want 1", created.ID)
	}
}

func TestInventoryService_Update(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")

		var inventory Inventory
		json.NewDecoder(r.Body).Decode(&inventory)

		inventory.UpdatedAt = "2024-01-02T00:00:00Z"
		json.NewEncoder(w).Encode(inventory)
	})

	ctx := context.Background()
	updateInventory := &Inventory{
		Title:    "更新された在庫",
		Quantity: 20,
		Unit:     "個",
		Category: "テストカテゴリ",
		State:    "中古",
		Place:    "倉庫B",
	}

	updated, err := client.Inventory.Update(ctx, 1, updateInventory)
	if err != nil {
		t.Errorf("Inventory.Update returned error: %v", err)
	}

	if updated.Title != "更新された在庫" {
		t.Errorf("Inventory.Update returned title %q, want %q", updated.Title, "更新された在庫")
	}
}

func TestInventoryService_Delete(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Authorization", "Bearer test-token")
		testHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.Inventory.Delete(ctx, 1)
	if err != nil {
		t.Errorf("Inventory.Delete returned error: %v", err)
	}
}

func TestInventoryService_Error(t *testing.T) {
	client, mux, close := setup()
	defer close()

	mux.HandleFunc("/inventories/999", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    404,
			Status:  "error",
			Message: "Inventory not found",
		})
	})

	ctx := context.Background()
	_, err := client.Inventory.Get(ctx, 999)
	if err == nil {
		t.Error("Expected error to be returned")
	}

	if errResp, ok := err.(*ErrorResponse); !ok {
		t.Errorf("Expected ErrorResponse, got %T", err)
	} else if errResp.Code != 404 {
		t.Errorf("Expected error code 404, got %d", errResp.Code)
	}
}
