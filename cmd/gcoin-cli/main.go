package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/superstas/gcoin/cmd"
	"github.com/superstas/gcoin/gcoin/amount"
	cli_rpc "github.com/superstas/gcoin/gcoin/network/cli"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var (
	to       string
	a        float64
	nodeAddr string
)

func main() {
	app := cli.NewApp()
	app.Name = cmd.GCoinASCIILogo() + "\ngcoin cli tool"
	app.Version = "0.1"
	app.Usage = "command-line client for gcoin"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "node-addr",
			Usage:       "an hostname:port of running node",
			Value:       "localhost:10000",
			Destination: &nodeAddr,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "this command sends coins",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "to",
					Usage:       "an address",
					Destination: &to,
				},
				cli.Float64Flag{
					Name:        "amount",
					Usage:       "amount of coins",
					Destination: &a,
				},
			},
			Action: func(c *cli.Context) error {
				if err := send(); err != nil {
					return cli.NewExitError(err, 1)
				}
				return nil
			},
		},
		{
			Name:    "getbalance",
			Aliases: []string{"gblnc"},
			Usage:   "this command returns all known addresses with UTXOs",
			Action: func(c *cli.Context) error {
				if err := getBalance(); err != nil {
					return cli.NewExitError(err, 1)
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}

func getBalance() error {
	b, err := connect().GetBalance(context.Background(), &cli_rpc.GetBalanceRequest{})
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address", "Balance"})
	var t amount.Amount
	for _, balance := range b.Addresses {
		a := amount.Amount(balance.Balance)
		table.Append([]string{balance.Address, a.String()})
		t += a
	}
	table.SetFooter([]string{"TOTAL", t.String()})
	table.Render()
	return nil
}

func send() error {
	if a == 0 {
		return errors.New("amount is zero")
	}

	a, err := amount.NewAmount(a)
	if err != nil {
		return err
	}

	b, err := connect().Send(context.Background(), &cli_rpc.SendRequest{
		ToAddress: to,
		Amount:    int64(a),
	})

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Transaction ID"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor})
	table.Append([]string{b.TxId})
	table.Render()
	return nil
}

func connect() cli_rpc.CliServiceClient {
	conn, err := grpc.Dial(nodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return cli_rpc.NewCliServiceClient(conn)
}
