package transaction

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetTransactions() ([]Transaction, error) {
	Transactions, err := usecase.Repo.GetTransactions()
	return Transactions, err
}

func (usecase Usecase) GetTransaction(id string) (*Transaction, error) {
	transaction, err := usecase.Repo.GetTransaction(id)
	return transaction, err
}

func (usecase Usecase) CreateTransaction(transaction *Transaction) error {
	err := usecase.Repo.CreateTransaction(transaction)
	return err
}
