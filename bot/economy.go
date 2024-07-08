package bot


import (
	parquet "github.com/parquet-go/parquet-go"
)

type RowType struct { 
	UserID string 
	Balance int 
}


func CheckBalance(messageAuthorID string){
	// Check the balance of the user
	

}


func InitialiseEconomy(){
	// Initialise the economy
	if err := parquet.WriteFile("data/economy.parquet", []RowType{
		{UserID: "testUser", Balance: 100},
	}); err != nil {
		// ...
	}
}