package primitives

import (
	"fmt"

	"goshake/util"
)

// Input

type Input struct {
	prevHash []byte
	prevIndex uint32
	sequence uint32
}

func (input *Input) Read(data []byte, offset *uint64) {
	input.prevHash = util.ReadBytes(data, offset, 32)
	input.prevIndex = util.ReadU32(data, offset)
	input.sequence = util.ReadU32(data, offset)
}

func (input Input) Print() {
	fmt.Printf("  txid: %032x \n", input.prevHash)
	fmt.Printf("  index: %08d \n", input.prevIndex)
	fmt.Printf("  sequence: %08x \n", input.sequence)
}

// Covenant

type Covenant struct {
	covType uint8
	itemsCount uint64
	items []*Item
}

func (cov *Covenant) Read(data []byte, offset *uint64) {
	cov.covType = util.ReadU8(data, offset)
	cov.itemsCount = util.ReadVarInt(data, offset)
	for i := uint64(0); i < cov.itemsCount; i++ {
		item := new(Item)
		item.Read(data, offset)
		cov.items = append(cov.items, item)
	}
}

func (cov Covenant) Print() {
	fmt.Printf("  covenant type: %02d \n", cov.covType)
	fmt.Printf("  covenant item count: %02d \n", cov.itemsCount)

	for i := uint64(0); i < cov.itemsCount; i++ {
		item := cov.items[i]
		fmt.Printf("   item #%d: \n", i)
		fmt.Printf("    item size: %02x (%d bytes) \n", item.size, item.size)
		fmt.Printf("    item: %x \n", item.data)
	}
}

// Output

type Output struct {
	value uint64
	addrVersion uint8
	addrHashSize uint8
	addrHash []byte
	covenant *Covenant
}

func (output *Output) Read(data []byte, offset *uint64) {
	output.value = util.ReadU64(data, offset)
	output.addrVersion = util.ReadU8(data, offset)
	output.addrHashSize = util.ReadU8(data, offset)
	output.addrHash = util.ReadBytes(data, offset, uint64(output.addrHashSize))
	output.covenant = new(Covenant)
	output.covenant.Read(data, offset)
}

func (output Output) Print() {
	fmt.Printf("  value: %d \n", output.value)
	fmt.Printf("  address version: %02x \n", output.addrVersion)
	fmt.Printf("  address hash size: %02x (%d bytes) \n", output.addrHashSize, output.addrHashSize)
	fmt.Printf("  address hash: %x \n", output.addrHash)
	output.covenant.Print()
}

// Witness

type Witness struct {
	itemsCount uint64
	items []*Item
}

func (wit *Witness) Read(data []byte, offset *uint64) {
	wit.itemsCount = util.ReadVarInt(data, offset)
	for i := uint64(0); i < wit.itemsCount; i++ {
		item := new(Item)
		item.Read(data, offset)
		wit.items = append(wit.items, item)
	}
}

func (wit Witness) Print() {
	fmt.Printf("   stack size: %02d \n", wit.itemsCount)

	for i := uint64(0); i < wit.itemsCount; i++ {
		fmt.Printf("    item #%d: \n", i)
		item := wit.items[i]
		fmt.Printf("     item size: %02x (%d bytes) \n", item.size, item.size)
		fmt.Printf("     item: %x \n", item.data)
	}
}

// Item (generic, for witness stack and covenant items)

type Item struct {
	size uint64
	data []byte
}

func (item *Item) Read(data []byte, offset *uint64) {
	item.size = util.ReadVarInt(data, offset)
	item.data = util.ReadBytes(data, offset, uint64(item.size))
}

// TX

type TX struct {
	version uint32
	inputs []*Input
	outputs []*Output
	locktime uint32
	witness []*Witness
}

func (tx *TX) Read(data []byte, offset *uint64) {
	tx.version = util.ReadU32(data, offset)

	incount := util.ReadVarInt(data, offset)
	for i := uint64(0); i < incount; i++ {
		input := new(Input)
		input.Read(data, offset)
		tx.inputs = append(tx.inputs, input)
	}

	outcount := util.ReadVarInt(data, offset)
	for i := uint64(0); i < outcount; i++ {
		output := new(Output)
		output.Read(data, offset)
		tx.outputs = append(tx.outputs, output)
	}

	tx.locktime = util.ReadU32(data, offset)

	for i := uint64(0); i < incount; i++ {
		witness := new(Witness)
		witness.Read(data, offset)
		tx.witness = append(tx.witness, witness)
	}
}

func (tx TX) Print() {
	fmt.Printf("version: %08x \n", tx.version)

	fmt.Printf("input count: %02d \n", len(tx.inputs))
	for i, input := range tx.inputs {
		fmt.Printf(" input #%d: \n", i)
		input.Print()
	}

	fmt.Printf("output count: %02d \n", len(tx.outputs))
	for i, output := range tx.outputs {
		fmt.Printf(" output #%d: \n", i)
		output.Print()
	}

	fmt.Printf("locktime: %08x \n", tx.locktime)

	fmt.Printf("witness: \n")
	for i, witness := range tx.witness {
		fmt.Printf(" witness for input #%d: \n", i)
		witness.Print()
	}
}
