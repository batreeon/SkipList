package skipList

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

// 讲skiplist的key:value编码为二进制数据
func (sl *skipList[K, V]) Pack() ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	m := make(map[K]V)
	for node := sl.Front(); node != nil; node = node.Next() {
		m[node.key] = node.value
	}
	err := encoder.Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// 将二进制数据解码为skiplist的key:value
func (sl *skipList[K, V]) Unpack(data []byte) (map[K]V, error) {
	buffer := bytes.NewReader(data)
	decoder := gob.NewDecoder(buffer)

	m := make(map[K]V)
	err := decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 将skiplist中的键值对数据存入文件
func (sl *skipList[K, V]) DumpFile(path string) error {
	dataBytes, err := sl.Pack()
	if err != nil {
		return err
	}
	f, _ := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0666)
	defer f.Close()
	_, err = f.Write(dataBytes)
	if err != nil {
		return err
	}
	return nil
}

// 从文件中还原出skiplist节点
func (sl *skipList[K, V]) LoadFile(path string) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}
	binaryData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	m, err := sl.Unpack(binaryData)
	if err != nil {
		return err
	}
	for k, v := range m {
		sl.Insert(k, v)
	}
	return nil
}