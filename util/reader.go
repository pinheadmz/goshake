package util


func ReadU8(data []byte, offset *uint64) uint8 {
	ret := data[*offset]
	*offset += 1
	return ret
}

func ReadU16(data []byte, offset *uint64) uint16 {
	ret := uint16(data[*offset] | data[*offset + 1] << 8)
	*offset += 2
	return ret
}

func ReadU32(data []byte, offset *uint64) uint32 {
	var ret uint32
	ret |= uint32(data[*offset])
	ret |= uint32(data[*offset + 1]) << 8
	ret |= uint32(data[*offset + 2]) << 16
	ret |= uint32(data[*offset + 3]) << 24
	*offset += 4
	return ret
}

func ReadU64(data []byte, offset *uint64) uint64 {
	lo := ReadU32(data, offset)
	hi := ReadU32(data, offset)
	ret := uint64(lo | hi << 16)
	return ret
}

func ReadVarInt(data []byte, offset *uint64) uint64 {
	switch data[*offset] {
		case 0xff:
			*offset += 1
			return ReadU64(data, offset)
		case 0xfe:
			*offset += 1
			return uint64(ReadU32(data, offset))
		case 0xfd:
			*offset += 1
			return uint64(ReadU16(data, offset))
		default:
			ret := uint64(data[*offset])
			*offset += 1
			return ret
	}
}

func ReadBytes(data []byte, offset *uint64, size uint64) []byte {
	ret := data[*offset:*offset + size]
	*offset += size
	return ret
}