package goEssential

/*
Create constants for file permissions using iota with bit shifting:
    Read = 1 << iota (should equal 1)
    Write (should equal 2)
    Execute (should equal 4)
Then create combined permissions:
    ReadWrite (should equal 3)
    AllPermissions (should equal 7)
*/

import "fmt"

const (
	Read    = 1 << iota // 1 << 0 = 1 (binary: 00000001)
	Write               // 1 << 1 = 2 (binary: 00000010)
	Execute             // 1 << 2 = 4 (binary: 00000100)
)

// Combined permissions using bitwise OR
const (
	ReadWrite      = Read | Write           // 1 | 2 = 3
	AllPermissions = Read | Write | Execute // 1 | 2 | 4 = 7
)

// << its behave as bit shifter: push 1 to position 0 of byte
// Each value has exactly one bit set at a different position, which allows us to combine them using bitwise OR (|) without overlapping
// so Read | Write: 00000001 | 00000010 = 00000011 = 3
func functionTwenty() {

	// Test your implementation
	fmt.Println("=== Basic Permissions ===")
	fmt.Printf("Read: %d\n", Read)
	fmt.Printf("Write: %d\n", Write)
	fmt.Printf("Execute: %d\n", Execute)

	fmt.Println("\n=== Combined Permissions ===")
	fmt.Printf("ReadWrite: %d\n", ReadWrite)
	fmt.Printf("AllPermissions: %d\n", AllPermissions)

	// Bonus: Verify the combinations work
	fmt.Printf("Read | Write = %d (should equal ReadWrite: %v)\n",
		Read|Write, Read|Write == ReadWrite)
	fmt.Printf("Read | Write | Execute = %d (should equal AllPermissions: %v)\n",
		Read|Write|Execute, Read|Write|Execute == AllPermissions)
}
