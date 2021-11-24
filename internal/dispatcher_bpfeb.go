// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadDispatcher returns the embedded CollectionSpec for dispatcher.
func loadDispatcher() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_DispatcherBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load dispatcher: %w", err)
	}

	return spec, err
}

// loadDispatcherObjects loads dispatcher and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *dispatcherObjects
//     *dispatcherPrograms
//     *dispatcherMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadDispatcherObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadDispatcher()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// dispatcherSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherSpecs struct {
	dispatcherProgramSpecs
	dispatcherMapSpecs
}

// dispatcherSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherProgramSpecs struct {
	Dispatcher *ebpf.ProgramSpec `ebpf:"dispatcher"`
}

// dispatcherMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dispatcherMapSpecs struct {
	Bindings           *ebpf.MapSpec `ebpf:"bindings"`
	DestinationMetrics *ebpf.MapSpec `ebpf:"destination_metrics"`
	Destinations       *ebpf.MapSpec `ebpf:"destinations"`
	Sockets            *ebpf.MapSpec `ebpf:"sockets"`
}

// dispatcherObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherObjects struct {
	dispatcherPrograms
	dispatcherMaps
}

func (o *dispatcherObjects) Close() error {
	return _DispatcherClose(
		&o.dispatcherPrograms,
		&o.dispatcherMaps,
	)
}

// dispatcherMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherMaps struct {
	Bindings           *ebpf.Map `ebpf:"bindings"`
	DestinationMetrics *ebpf.Map `ebpf:"destination_metrics"`
	Destinations       *ebpf.Map `ebpf:"destinations"`
	Sockets            *ebpf.Map `ebpf:"sockets"`
}

func (m *dispatcherMaps) Close() error {
	return _DispatcherClose(
		m.Bindings,
		m.DestinationMetrics,
		m.Destinations,
		m.Sockets,
	)
}

// dispatcherPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadDispatcherObjects or ebpf.CollectionSpec.LoadAndAssign.
type dispatcherPrograms struct {
	Dispatcher *ebpf.Program `ebpf:"dispatcher"`
}

func (p *dispatcherPrograms) Close() error {
	return _DispatcherClose(
		p.Dispatcher,
	)
}

func _DispatcherClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed dispatcher_bpfeb.o
var _DispatcherBytes []byte
