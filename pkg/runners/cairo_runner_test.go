package runners_test

import (
	"testing"

	"github.com/lambdaclass/cairo-vm.go/pkg/lambdaworks"
	"github.com/lambdaclass/cairo-vm.go/pkg/parser"
	"github.com/lambdaclass/cairo-vm.go/pkg/runners"
	"github.com/lambdaclass/cairo-vm.go/pkg/vm"
	"github.com/lambdaclass/cairo-vm.go/pkg/vm/memory"
)

func TestNewCairoRunnerInvalidBuiltin(t *testing.T) {
	// Create a Program with one fake instruction
	program_data := make([]memory.MaybeRelocatable, 1)
	empty_identifiers := make(map[string]parser.Identifier, 0)
	program_data[0] = *memory.NewMaybeRelocatableFelt(lambdaworks.FeltOne())
	program := vm.Program{Data: program_data, Builtins: []string{"fake_builtin"}, Identifiers: &empty_identifiers}
	// Create CairoRunner
	_, err := runners.NewCairoRunner(program)
	if err == nil {
		t.Errorf("Expected creating a CairoRunner with fake builtin to fail")
	}
}
func TestInitializeRunnerNoBuiltinsNoProofModeEmptyProgram(t *testing.T) {
	// Create a Program with empty data
	program_data := make([]memory.MaybeRelocatable, 0)
	empty_identifiers := make(map[string]parser.Identifier, 0)
	program := vm.Program{Data: program_data, Identifiers: &empty_identifiers}
	// Create CairoRunner
	runner, err := runners.NewCairoRunner(program)
	if err != nil {
		t.Errorf("NewCairoRunner error in test: %s", err)
	}
	// Initialize the runner
	end_ptr, err := runner.Initialize()
	if err != nil {
		t.Errorf("Initialize error in test: %s", err)
	}
	if end_ptr.SegmentIndex != 3 || end_ptr.Offset != 0 {
		t.Errorf("Wrong end ptr value, got %+v", end_ptr)
	}

	// Check CairoRunner values
	if runner.ProgramBase.SegmentIndex != 0 || runner.ProgramBase.Offset != 0 {
		t.Errorf("Wrong ProgramBase value, got %+v", runner.ProgramBase)
	}

	// Check Vm's RunContext values
	if runner.Vm.RunContext.Pc.SegmentIndex != 0 || runner.Vm.RunContext.Pc.Offset != 0 {
		t.Errorf("Wrong Pc value, got %+v", runner.Vm.RunContext.Pc)
	}
	if runner.Vm.RunContext.Ap.SegmentIndex != 1 || runner.Vm.RunContext.Ap.Offset != 2 {
		t.Errorf("Wrong Ap value, got %+v", runner.Vm.RunContext.Ap)
	}
	if runner.Vm.RunContext.Fp.SegmentIndex != 1 || runner.Vm.RunContext.Fp.Offset != 2 {
		t.Errorf("Wrong Fp value, got %+v", runner.Vm.RunContext.Fp)
	}

	// Check memory

	// Program segment
	// 0:0 program_data[0] should be empty
	value, err := runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 0, Offset: 0})
	if err == nil {
		t.Errorf("Expected addr 0:0 to be empty for empty program, got: %+v", value)
	}

	// Execution segment
	// 1:0 return_fp
	value, err = runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 1, Offset: 0})
	if err != nil {
		t.Errorf("Memory Get error in test: %s", err)
	}
	rel, ok := value.GetRelocatable()
	if !ok || rel.SegmentIndex != 2 || rel.Offset != 0 {
		t.Errorf("Wrong value for address 1:0: %d", rel)
	}
	// 1:1 end_ptr
	value, err = runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 1, Offset: 1})
	if err != nil {
		t.Errorf("Memory Get error in test: %s", err)
	}
	rel, ok = value.GetRelocatable()
	if !ok || rel.SegmentIndex != 3 || rel.Offset != 0 {
		t.Errorf("Wrong value for address 1:1: %d", rel)
	}
}

func TestInitializeRunnerNoBuiltinsNoProofModeNonEmptyProgram(t *testing.T) {
	// Create a Program with one fake instruction
	program_data := make([]memory.MaybeRelocatable, 1)
	program_data[0] = *memory.NewMaybeRelocatableFelt(lambdaworks.FeltFromUint64(1))
	empty_identifiers := make(map[string]parser.Identifier, 0)
	program := vm.Program{Data: program_data, Identifiers: &empty_identifiers}
	// Create CairoRunner
	runner, err := runners.NewCairoRunner(program)
	if err != nil {
		t.Errorf("NewCairoRunner error in test: %s", err)
	}
	// Initialize the runner
	end_ptr, err := runner.Initialize()
	if err != nil {
		t.Errorf("Initialize error in test: %s", err)
	}
	if end_ptr.SegmentIndex != 3 || end_ptr.Offset != 0 {
		t.Errorf("Wrong end ptr value, got %+v", end_ptr)
	}

	// Check CairoRunner values
	if runner.ProgramBase.SegmentIndex != 0 || runner.ProgramBase.Offset != 0 {
		t.Errorf("Wrong ProgramBase value, got %+v", runner.ProgramBase)
	}

	// Check Vm's RunContext values
	if runner.Vm.RunContext.Pc.SegmentIndex != 0 || runner.Vm.RunContext.Pc.Offset != 0 {
		t.Errorf("Wrong Pc value, got %+v", runner.Vm.RunContext.Pc)
	}
	if runner.Vm.RunContext.Ap.SegmentIndex != 1 || runner.Vm.RunContext.Ap.Offset != 2 {
		t.Errorf("Wrong Ap value, got %+v", runner.Vm.RunContext.Ap)
	}
	if runner.Vm.RunContext.Fp.SegmentIndex != 1 || runner.Vm.RunContext.Fp.Offset != 2 {
		t.Errorf("Wrong Fp value, got %+v", runner.Vm.RunContext.Fp)
	}

	// Check memory

	// Program segment
	// 0:0 program_data[0]
	value, err := runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 0, Offset: 0})
	if err != nil {
		t.Errorf("Memory Get error in test: %s", err)
	}
	int, ok := value.GetFelt()
	if !ok || int != lambdaworks.FeltFromUint64(1) {
		t.Errorf("Wrong value for address 0:0: %d", int)
	}

	// Execution segment
	// 1:0 return_fp
	value, err = runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 1, Offset: 0})
	if err != nil {
		t.Errorf("Memory Get error in test: %s", err)
	}
	rel, ok := value.GetRelocatable()
	if !ok || rel.SegmentIndex != 2 || rel.Offset != 0 {
		t.Errorf("Wrong value for address 1:0: %d", rel)
	}
	// 1:1 end_ptr
	value, err = runner.Vm.Segments.Memory.Get(memory.Relocatable{SegmentIndex: 1, Offset: 1})
	if err != nil {
		t.Errorf("Memory Get error in test: %s", err)
	}
	rel, ok = value.GetRelocatable()
	if !ok || rel.SegmentIndex != 3 || rel.Offset != 0 {
		t.Errorf("Wrong value for address 1:1: %d", rel)
	}
}
