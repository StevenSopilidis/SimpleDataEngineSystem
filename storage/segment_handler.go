package storage

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	hk "github.com/StevenSopilidis/SimpleDataEngineSystem/internal"

	"github.com/google/uuid"
)

const (
	segmentLocation = "storage/segments"
)

type SegmentHandler struct {
}

/*
*
Segment Structure
512 bits -> key (its in hex so its 64 bytes)
16 bits -> length of data
len bits -> data
*/
type SegmentEntry struct {
	Key  string
	Len  uint16
	Data []byte
}

func getAllFilePaths(dirPath string) []string {
	var filePaths []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		filePaths = append(filePaths, path)
		print(path)
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return filePaths
}

// getCreationTime returns the creation time of a file
func getCreationTime(filePath string) time.Time {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}

	return fileInfo.ModTime()
}

func (f *SegmentHandler) FindItem(key []byte) *SegmentEntry {
	filePaths := getAllFilePaths("storage/segments")
	sort.Slice(filePaths, func(i, j int) bool {
		return getCreationTime(filePaths[i]).Before(getCreationTime(filePaths[j]))
	})
	for _, filePath := range filePaths {
		fileName := filepath.Base(filePath)
		segment, err := f.GetSegment(fileName)
		if err != nil {
			return nil
		}
		item, found := segment[hk.HashStringToSHA256(key)]
		if found {
			return &item
		}
	}
	return nil
}

// reads a whole segment from fs
func (f *SegmentHandler) GetSegment(id string) (map[string]SegmentEntry, error) {
	fileLocation := fmt.Sprintf("%s/%s.dat", segmentLocation, id)
	data, err := os.ReadFile(fileLocation)
	entries := make(map[string]SegmentEntry, 0)
	if err != nil {
		return entries, err
	}
	i := 0
	for i < len(data) {
		keyBuff := data[i : i+64]
		key := string(keyBuff)
		lenBuff := data[i+64 : i+66]
		len := binary.LittleEndian.Uint16(lenBuff)
		dataBuff := data[i+66 : i+66+int(len)]
		i += int(66 + len)
		entries[key] = SegmentEntry{
			Key:  key,
			Len:  len,
			Data: dataBuff,
		}
	}

	return entries, err
}

// gets the data from RemoveAll() function in memtable
func (f *SegmentHandler) AppendSegment(data map[string][]byte) error {
	segmentId := uuid.New()
	filename := fmt.Sprintf("%s/%s.dat", segmentLocation, segmentId.String())
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	segmentData := make([]byte, 0)
	for k := range data {
		var tmp []byte
		if err != nil {
			return err
		}
		keyBuff := []byte(k)
		tmp = append(tmp, keyBuff...)
		datalen := len(data[k])
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(datalen))
		tmp = append(tmp, buf...)
		tmp = append(tmp, data[k]...)
		segmentData = append(segmentData, tmp...)
	}
	_, err = file.Write(segmentData)
	return err
}
