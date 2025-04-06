# zaico-go

Zaico APIのGo言語クライアントライブラリです。

## 概要

このライブラリは、[Zaico](https://zaico.co.jp/)の在庫管理APIをGo言語から簡単に利用するためのクライアントライブラリです。

## インストール

```bash
go get github.com/fukata/zaico-go
```
## 使用方法

### クライアントの初期化

```go
import "github.com/fukata/zaico-go"

// APIトークンを使用してクライアントを初期化
client := zaico.NewClient("your-api-token")
```

### 在庫データの操作

#### 在庫データの一覧取得

```go
// 基本的な一覧取得
inventories, err := client.Inventory.List(ctx, nil)

// 検索条件を指定して一覧取得
opts := &zaico.InventoryListOptions{
    Title:    "テスト在庫",
    Category: "テストカテゴリ",
    Place:    "倉庫A",
    Page:     1,
}
inventories, err := client.Inventory.List(ctx, opts)
```

#### 在庫データの個別取得

```go
inventory, err := client.Inventory.Get(ctx, 1)
```

#### 在庫データの作成

```go
newInventory := &zaico.Inventory{
    Title:    "新しい在庫",
    Quantity: 10,
    Unit:     "個",
    Category: "カテゴリA",
    State:    "新品",
    Place:    "倉庫A",
}
created, err := client.Inventory.Create(ctx, newInventory)
```

#### 在庫データの更新

```go
updateInventory := &zaico.Inventory{
    Title:    "更新された在庫",
    Quantity: 20,
    Unit:     "個",
    Category: "カテゴリA",
    State:    "中古",
    Place:    "倉庫B",
}
updated, err := client.Inventory.Update(ctx, 1, updateInventory)
```

#### 在庫データの削除

```go
err := client.Inventory.Delete(ctx, 1)
```

### エラーハンドリング

```go
inventory, err := client.Inventory.Get(ctx, 999)
if err != nil {
    if errResp, ok := err.(*zaico.ErrorResponse); ok {
        // エラーレスポンスの詳細を取得
        fmt.Printf("Error: %s (Code: %d)\n", errResp.Message, errResp.Code)
    } else {
        // その他のエラー
        fmt.Printf("Error: %v\n", err)
    }
}
```

## ライセンス

MIT License

## 作者

[fukata](https://github.com/fukata) 