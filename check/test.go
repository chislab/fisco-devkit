package check

import (
	"context"
	"fisco/build/hello"
	"fmt"
	"github.com/chislab/go-fiscobcos/accounts/abi/bind"
	"github.com/chislab/go-fiscobcos/common"
	"github.com/chislab/go-fiscobcos/core/types"
)


var HelloAddr common.Address
var dHello *hello.Hello

func Test() error {
	ctx := context.Background()
	height, err := GethCli.BlockNumber(ctx)
	if err != nil {
		return err
	}
	auth := NewAuthFromPriKey(height)
	addr, tx, _dHello, err := hello.DeployHello(auth, GethCli)
	if err != nil {
		return err
	}
	HelloAddr = addr
	dHello = _dHello

	fmt.Println("deploy hello txHash", tx.Hash().String())
	msg, ok, err := dHello.Get(callOpts)
	if err != nil {
		return err
	}
	fmt.Println(msg, ok)

	ch := make(chan *hello.HelloEvtSet)
	go dHello.WatchEvtSet(&bind.WatchOpts{
		Start:   new(uint64),
		Context: ctx,
	}, ch)

	tx, err = dHello.Set(auth, "deng mimi")
	if err != nil {
		return err
	}
	fmt.Println("call hello.Set() txHash", tx.Hash().String())

	evt := <-ch
	fmt.Println("[msg from event]", evt.Msg)

	return nil
}


func parseHelloEvt(msg *types.Log) {
	switch msg.Topics[0] {
	case dHello.ABI.Events["EvtSet"].ID:
		evt, _ := dHello.ParseEvtSet(*msg)
		fmt.Println("== EvtRecordUploaded ==", evt.Msg, evt.From.String())
		// TODO mongodb, mysql cached.
	}
}