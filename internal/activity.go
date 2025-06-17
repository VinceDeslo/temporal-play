package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
    log.Printf("Withdrawing $%d from account %s.\n\n",
        data.Amount,
        data.SourceAccount,
    )

    referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
    bank := BankingService{"bank-api.example.com"}
    confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
    return confirmation, err
}

func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
    log.Printf("Depositing $%d into account %s.\n\n",
        data.Amount,
        data.TargetAccount,
    )

    referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
    bank := BankingService{"bank-api.example.com"}
    // Uncomment the next line and comment the one after that to simulate an unknown failure
    // confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
    confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
    return confirmation, err
}

func Refund(ctx context.Context, data PaymentDetails) (string, error) {
    log.Printf("Refunding $%v back into account %v.\n\n",
        data.Amount,
        data.SourceAccount,
    )

    referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
    bank := BankingService{"bank-api.example.com"}
    confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
    return confirmation, err
}

func Explain(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Requiring payment details explanation from llm.\n\n")

	llm, err := NewLLMService(ctx)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	prompt := fmt.Sprintf("Explain in text only human readable format the transaction described by the following data: %s", string(bytes))

	result, err := llm.Prompt(ctx, prompt)
	if err != nil {
		return "", err
	}
	return result, err	
}
