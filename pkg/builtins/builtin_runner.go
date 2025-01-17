package builtins

import "github.com/lambdaclass/cairo-vm.go/pkg/vm/memory"

type BuiltinRunner interface {
	// Returns the first address of the builtin's memory segment
	Base() memory.Relocatable
	// Returns the name of the builtin
	Name() string
	// Creates a memory segment for the builtin and initializes its base
	InitializeSegments(*memory.MemorySegmentManager)
	// Returns the builtin's initial stack
	InitialStack() []memory.MaybeRelocatable
	// Attempts to deduce the value of a memory cell given by its address. Can return either a nil pointer and an error, if an error arises during the deduction,
	// a valid pointer and nil if the deduction was succesful, or a nil pointer and nil if there is no deduction for the memory cell
	DeduceMemoryCell(memory.Relocatable, *memory.Memory) (*memory.MaybeRelocatable, error)
	// Adds a validation rule to the memory
	// Validation rules are applied when a value is inserted into the builtin's segment
	AddValidationRule(*memory.Memory)
	// TODO: Later additions -> Some of them could depend on a Default Implementation
	// // Most of them depend on Layouts being implemented
	// // Use cases:
	// // I. PROOF_MODE
	// // Returns the builtin's ratio, can be nil if the layout is dynamic
	// Ratio() *uint // proof-mode end_run logic
	// // Returns the builtin's allocated memory units
	// GetAllocatedMemoryUnits(*vm.VirtualMachine) (uint, error) // proof-mode end_run logic
	// // Returns the list of memory addresses used by the builtin
	// GetMemoryAccesses(*memory.MemorySegmentManager) ([]memory.Relocatable, error) // proof-mode end_run logic
	// GetUsedCells(*memory.MemorySegmentManager) (uint, error)                      // proof-mode end_run logic
	// GetRangeCheckUsage(*memory.Memory) (*uint, *uint)                             // proof-mode end_run logic
	// GetUsedPermRangeCheckLimits(*vm.VirtualMachine) (uint, error)                 // proof-mode end_run logic
	// GetUsedDilutedCheckUnits(diluted_spacing uint, diluted_n_bits uint) uint      // proof-mode end_run logic
	// GetUsedCellsAndAllocatedSizes(*vm.VirtualMachine) (uint, uint, error)         // proof-mode end_run logic + finalize_segments
	// // II. SECURITY (secure-run flag cairo-run || verify-secure flag run_from_entrypoint)
	// RunSecurityChecks(*vm.VirtualMachine) error // verify_secure_runner logic
	// // Returns the base & stop_ptr, stop_ptr can be nil
	// GetMemorySegmentAddresses() (memory.Relocatable, *memory.Relocatable) //verify_secure_runner logic
	// // III. STARKNET-SPECIFIC
	// GetUsedInstances(*memory.MemorySegmentManager) (uint, error) // get_execution_resources (starknet use case)
	// // IV. GENERAL CASE (but not critical)
	// FinalStack(*memory.MemorySegmentManager, memory.Relocatable) (memory.Relocatable, error) // read_return_values
}
