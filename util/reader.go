package util

type Reader struct {
	data []byte
	offset uint64
}

func (reader *Reader) Init(data []byte) {
	reader.data = data
	reader.offset = 0
}

func (reader *Reader) ReadU8() uint8 {
	ret := reader.data[reader.offset]
	reader.offset += 1
	return ret
}

func (reader *Reader) ReadU16() uint16 {
	ret := uint16(reader.data[reader.offset] | reader.data[reader.offset + 1] << 8)
	reader.offset += 2
	return ret
}

func (reader *Reader) ReadU32() uint32 {
	var ret uint32
	ret |= uint32(reader.data[reader.offset])
	ret |= uint32(reader.data[reader.offset + 1]) << 8
	ret |= uint32(reader.data[reader.offset + 2]) << 16
	ret |= uint32(reader.data[reader.offset + 3]) << 24
	reader.offset += 4
	return ret
}

func (reader *Reader) ReadU64() uint64 {
	lo := reader.ReadU32()
	hi := reader.ReadU32()
	var ret uint64
	ret |= uint64(lo)
	ret |= uint64(hi) << 32
	return ret
}

func (reader *Reader) ReadVarInt() uint64 {
	switch reader.data[reader.offset] {
		case 0xff:
			reader.offset += 1
			return reader.ReadU64()
		case 0xfe:
			reader.offset += 1
			return uint64(reader.ReadU32())
		case 0xfd:
			reader.offset += 1
			return uint64(reader.ReadU16())
		default:
			ret := uint64(reader.data[reader.offset])
			reader.offset += 1
			return ret
	}
}

func (reader *Reader) ReadBytes(size uint64) []byte {
	ret := reader.data[reader.offset:reader.offset + size]
	reader.offset += size
	return ret
}