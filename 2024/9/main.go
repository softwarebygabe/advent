package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/softwarebygabe/advent/pkg/util"
)

type MemBlock struct {
	id    int
	empty bool
}

func (m MemBlock) String() string {
	if m.empty {
		return "."
	}
	return fmt.Sprintf("%d", m.id)
}

type HardDrive struct {
	diskMapRaw        string
	diskMap           []int
	memory            []MemBlock
	compacted         []MemBlock
	compactedChecksum int
}

func NewHardDrive(diskMapRaw string) *HardDrive {
	hd := &HardDrive{
		diskMapRaw: diskMapRaw,
	}
	for _, char := range strings.Split(diskMapRaw, "") {
		hd.diskMap = append(hd.diskMap, util.MustParseInt(char))
	}
	hd.generateMemory()
	// hd.compact() // Part 1
	// hd.compact2() // Part 2 (strings replace, takes too long)
	hd.compact3() // Part 2 (hopefully optimized)
	hd.compactedChecksum = computeChecksum(hd.compacted)
	return hd
}

func (h *HardDrive) String() string {
	return fmt.Sprintf("Hard Drive:\n\tDisk Map: %s\n\tMemory: %v\n\tCompacted: %v\n\tChecksum: %d", h.diskMapRaw, h.memory, h.compacted, h.compactedChecksum)
}

func stringToMemory(s string) []MemBlock {
	mem := []MemBlock{}
	for _, char := range strings.Split(s, "") {
		if char == "." {
			mem = append(mem, MemBlock{empty: true})
		} else {
			mem = append(mem, MemBlock{id: util.MustParseInt(char)})
		}
	}
	return mem
}

func (h *HardDrive) generateMemory() {
	var free bool
	var fidx int
	for _, n := range h.diskMap {
		for i := 0; i < n; i++ {
			if free {
				h.memory = append(h.memory, MemBlock{empty: true})
			} else {
				h.memory = append(h.memory, MemBlock{id: fidx})
			}
		}
		if !free {
			fidx++
		}
		free = !free
	}
}

func computeChecksum(mem []MemBlock) int {
	var sum int
	for i, b := range mem {
		sum += i * b.id
	}
	return sum
}

func (h *HardDrive) compact() {
	memStack := util.NewStack[MemBlock]()
	for _, block := range h.memory {
		if !block.empty {
			memStack.Push(block)
		}
	}
	var moved int
	for _, block := range h.memory {
		if block.empty {
			b, ok := memStack.Pop()
			if !ok {
				panic("stack empty")
			}
			h.compacted = append(h.compacted, b)
			moved++
		} else {
			h.compacted = append(h.compacted, block)
		}
	}
	for i := 1; i <= moved; i++ {
		h.compacted[len(h.compacted)-i] = MemBlock{empty: true}
	}
}

func memBlockString(mem []MemBlock) string {
	var res string
	for _, b := range mem {
		res += b.String()
	}
	return res
}

func createFileStack(mem []MemBlock) (util.Stack[[]MemBlock], map[int]int) {
	fileStack := util.NewStack[[]MemBlock]()
	posMap := map[int]int{}
	currFile := []MemBlock{}
	var currFileIdx int
	var pastFirstBlock bool
	for idx, block := range mem {
		if !block.empty {
			newBlock := idx > 0 && mem[idx-1].id != block.id
			if newBlock && len(currFile) > 0 {
				fileStack.Push(currFile)
				posMap[currFile[0].id] = currFileIdx
				currFileIdx = idx
				currFile = []MemBlock{block}
				pastFirstBlock = true
				continue
			}
			if currFileIdx == 0 && pastFirstBlock {
				currFileIdx = idx
			}
			currFile = append(currFile, block)
			continue
		}
		if len(currFile) > 0 {
			if !pastFirstBlock {
				pastFirstBlock = true
			}
			fileStack.Push(currFile)
			posMap[currFile[0].id] = currFileIdx
			currFileIdx = 0
			currFile = []MemBlock{}
		}
	}
	fileStack.Push(currFile)
	posMap[currFile[0].id] = currFileIdx
	return *fileStack, posMap
}

func (h *HardDrive) compact2() {
	fileStack, _ := createFileStack(h.memory)
	h.compacted = h.memory
	for fileStack.Len() > 0 {
		fmt.Println("stack length:", fileStack.Len())
		start := time.Now()
		file, _ := fileStack.Pop()
		compactedString := memBlockString(h.compacted)
		// fmt.Println(compactedString)
		fileString := memBlockString(file)
		// fmt.Println(fileString)
		freeBlockString := strings.Repeat(".", len(file))
		// fmt.Println(freeBlockString)
		compactedStringSplit := strings.Split(compactedString, fileString)
		leftOfFile := compactedStringSplit[0]
		// find index of free block in left
		idx := strings.Index(leftOfFile, freeBlockString)
		if idx < 0 {
			continue // skip, no free space
		}
		// do replacement
		leftOfFileArr := strings.Split(leftOfFile, "")
		for i := idx; i < idx+len(freeBlockString); i++ {
			leftOfFileArr[i] = file[0].String()
		}
		newCompactedStringLeft := strings.Join(leftOfFileArr, "")
		if newCompactedStringLeft != leftOfFile {
			// moved
			if len(compactedStringSplit) > 1 {
				newCompactedString := newCompactedStringLeft + freeBlockString + compactedStringSplit[1]
				h.compacted = stringToMemory(newCompactedString)
			} else {
				newCompactedString := newCompactedStringLeft + freeBlockString
				h.compacted = stringToMemory(newCompactedString)
			}
		}
		fmt.Println("dur:", time.Since(start))
	}
}

func (h *HardDrive) compact3() {
	fileStack, filePosMap := createFileStack(h.memory)
	freeSpaceLens := []int{}
	freeSpacePos := map[int][]int{}

	var currlen, currpos int
	for idx, mem := range h.memory {
		if idx == 0 {
			continue
		}
		if mem.empty {
			if currlen < 1 {
				currpos = idx
			}
			currlen++
		} else if currpos != 0 {
			freeSpaceLens = append(freeSpaceLens, currlen)
			freeSpacePos[currlen] = append(freeSpacePos[currlen], currpos)
			currpos = 0
			currlen = 0
		}
	}
	if currlen > 0 {
		freeSpaceLens = append(freeSpaceLens, currlen)
		freeSpacePos[currlen] = append(freeSpacePos[currlen], currpos)
	}
	// fmt.Println(freeSpaceLens)
	// fmt.Println(freeSpacePos)
	// fmt.Println(filePosMap)

	h.compacted = append(h.compacted, h.memory...)
	for fileStack.Len() > 0 {
		fmt.Println("stack len:", fileStack.Len())
		start := time.Now()
		file, _ := fileStack.Pop()
		// fmt.Println(file)
		// fmt.Println(h.compacted)
		// fmt.Println(freeSpaceLens)
		// fmt.Println(freeSpacePos)
		// find the first free length that will work
		for lidx, l := range freeSpaceLens {
			fileLen := len(file)
			// fmt.Println("hello", l, fileLen)
			if l >= fileLen {
				// we found a length
				positions, ok := freeSpacePos[l]
				if !ok {
					// there are no more
					break
				}
				if len(positions) < 1 {
					// there are no more
					break
				}
				pos := positions[0]
				// fmt.Println("pos", pos)
				// fileIdx := strings.Index(memBlockString(h.compacted), memBlockString(file))
				fileIdx, ok := filePosMap[file[0].id]
				if !ok {
					panic("could not find file pos")
				}
				// fmt.Println("pos", pos)
				// fmt.Println("fileIdx", fileIdx)
				// fmt.Println("fileLen", fileLen)
				if pos < fileIdx {
					// found a suitable position
					// move it there
					for i := pos; i < pos+fileLen; i++ {
						h.compacted[i] = file[0]
					}
					// update map
					freeSpacePos[l] = positions[1:]
					l2 := l - fileLen
					freeSpaceLens[lidx] = l2
					if l > fileLen {
						// reduce l, re-map
						p2 := pos + fileLen
						newPos := append(freeSpacePos[l2], p2)
						sort.Ints(newPos)
						freeSpacePos[l2] = newPos
					}
					// now replace old spot with free space
					for i := fileIdx; i < fileIdx+fileLen; i++ {
						h.compacted[i] = MemBlock{empty: true}
					}
					break
				}
			}
		}
		fmt.Println("dur:", time.Since(start))
	}

}

func Run(filename string) {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	diskMap := lines[0]
	hd := NewHardDrive(diskMap)
	fmt.Println(hd.String())
}

func main() {
	// Run("input_ex.txt")
	Run("input_1.txt")
}
