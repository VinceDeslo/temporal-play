package main

import (
	"log"

	"github.com/VinceDeslo/temporal-play/internal"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client.", err)
    }
    defer c.Close()

    w := worker.New(c, internal.MoneyTransferTaskQueueName, worker.Options{})

    // This worker hosts both Workflow and Activity functions.
    w.RegisterWorkflow(internal.MoneyTransfer)
    w.RegisterActivity(internal.Withdraw)
    w.RegisterActivity(internal.Deposit)
    w.RegisterActivity(internal.Refund)

    // Start listening to the Task Queue.
    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("unable to start Worker", err)
    }
}

