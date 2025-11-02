
package product

import (
	"database/sql"
	"myapi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
  
    rows, err := s.db.Query("SELECT id, name, description, price, created_at FROM products")
    if err!=nil{
        return nil,err
    }
    products :=make([]types.Product,0)
    for rows.Next(){
        product,err:=ScanProductIntoRows(rows)
        if err!=nil{
            return nil,err
        }
        products=append(products,*product)
    }
    if rows.Err() != nil {
        return nil, rows.Err()
    }
    return products, nil
}





func ScanProductIntoRows(rows *sql.Rows) (*types.Product, error) {
	u := new(types.Product)
	// Table columns: id, firstname, lastname, email, password, created_at
	err := rows.Scan(&u.Id, &u.Name, &u.Description, &u.Price, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
