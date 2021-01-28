package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"code.cfops.it/sys/tubular/internal"
)

func bind(e *env, args ...string) error {
	set := e.newFlagSet("bind", `<label> <protocol> <ip[/mask]> <port>

Bind a given prefix, port and protocol to a label.
`)
	if err := set.Parse(args); errors.Is(err, flag.ErrHelp) {
		return nil
	} else if err != nil {
		return err
	}

	bind, err := bindingFromArgs(set.Args())
	if err != nil {
		return err
	}

	dp, err := e.openDispatcher()
	if err != nil {
		return err
	}
	defer dp.Close()

	return dp.AddBinding(bind)
}

func unbind(e *env, args ...string) error {
	set := e.newFlagSet("unbind", `<label> <protocol> <ip[/mask]> <port>

Remove a previously created binding.
`)
	if err := set.Parse(args); errors.Is(err, flag.ErrHelp) {
		return nil
	} else if err != nil {
		return err
	}

	bind, err := bindingFromArgs(set.Args())
	if err != nil {
		return err
	}

	dp, err := e.openDispatcher()
	if err != nil {
		return err
	}
	defer dp.Close()

	if err := dp.RemoveBinding(bind); err != nil {
		return err
	}

	e.stdout.Log("Removed", bind)
	return nil
}

func bindingFromArgs(args []string) (*internal.Binding, error) {
	if n := len(args); n != 4 {
		return nil, fmt.Errorf("expected label, protocol, ip and port but got %d arguments", n)
	}

	var proto internal.Protocol
	switch args[1] {
	case "udp":
		proto = internal.UDP
	case "tcp":
		proto = internal.TCP
	}

	port, err := strconv.ParseUint(args[3], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %s", err)
	}

	return internal.NewBinding(args[0], proto, args[2], uint16(port))
}
