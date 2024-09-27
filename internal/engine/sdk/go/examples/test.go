/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/vearch/vearch/v3/internal/engine/sdk/go/gamma"
	"github.com/vearch/vearch/v3/internal/pkg/vjson"
	"github.com/vearch/vearch/v3/internal/proto/vearchpb"
	"golang.org/x/exp/mmap"
)

type Options struct {
	Nprobe          int32
	DocID           uint64
	DocID2          uint64
	D               int
	MaxDocSize      int32
	AddDocNum       int
	SearchNum       uint64
	IndexingSize    int
	Filter          bool
	PrintDoc        bool
	SearchThreadNum int
	FieldsVec       []string
	FieldsType      []vearchpb.FieldType
	TableFieldsType []gamma.DataType
	Path            string
	LogDir          string
	VectorName      string
	RetrievalType   string
	StoreType       string
	Profiles        []string
	Feature         []byte
	ProfileFile     string
	FeatureFile     string
	FeatureMmap     *mmap.ReaderAt
	Engine          unsafe.Pointer
}

var opt *Options

func Init() {
	MaxDocSize := 10 * 10000
	FieldsVec := []string{"img", "_id", "field1", "field2", "field3"}
	opt = &Options{
		Nprobe:          256,
		DocID:           0,
		DocID2:          0,
		D:               512,
		MaxDocSize:      10 * 10000,
		AddDocNum:       10 * 10000,
		SearchNum:       10000,
		IndexingSize:    10 * 10000,
		Filter:          false,
		PrintDoc:        false,
		SearchThreadNum: 1,
		FieldsVec:       FieldsVec,
		FieldsType:      []vearchpb.FieldType{vearchpb.FieldType_STRING, vearchpb.FieldType_STRING, vearchpb.FieldType_STRING, vearchpb.FieldType_INT, vearchpb.FieldType_INT},
		TableFieldsType: []gamma.DataType{gamma.STRING, gamma.STRING, gamma.STRING, gamma.INT, gamma.INT},
		Path:            "./files",
		LogDir:          "./log",
		VectorName:      "abc",
		RetrievalType:   "IVFPQ",
		StoreType:       "RocksDB",
		Profiles:        make([]string, MaxDocSize*len(FieldsVec)),
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: [Program] [profile_file] [vectors_file]")
		os.Exit(-1)
	}
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	opt.ProfileFile = os.Args[1]
	opt.FeatureFile = os.Args[2]

	opt.FeatureMmap, _ = mmap.Open(opt.FeatureFile)

	config := struct {
		Path      string `json:"path"`
		SpaceName string `json:"space_name"`
		LogDir    string `json:"log_dir"`
	}{
		Path:      opt.Path,
		SpaceName: "test",
		LogDir:    opt.LogDir,
	}

	configJson, _ := vjson.Marshal(&config)
	opt.Engine = gamma.Init(configJson)
	fmt.Println("Inited")
}

func CreteTable() {
	kIVFPQParam :=
		string("{\"nprobe\" : 10, \"metric_type\" : \"InnerProduct\", \"ncentroids\" : 256,\"nsubvector\" : 64}")

	//kHNSWParam_str :=
	//	"{\"nlinks\" : 32, \"metric_type\" : \"InnerProduct\", \"efSearch\" : 64,\"efConstruction\" : 40}"
	//
	//kFLATParam_str := "{\"metric_type\" : \"InnerProduct\"}"
	table := gamma.Table{
		Name:        opt.VectorName,
		IndexType:   opt.RetrievalType,
		IndexParams: kIVFPQParam}

	for i := 0; i < len(opt.FieldsVec); i++ {
		isIndex := false
		if opt.Filter && (i == 0 || i == 2 || i == 3 || i == 4) {
			isIndex = true
		}
		fieldInfo := gamma.FieldInfo{
			Name:     opt.FieldsVec[i],
			DataType: opt.TableFieldsType[i],
			IsIndex:  isIndex,
		}
		table.Fields = append(table.Fields, fieldInfo)
	}

	vectorInfo := gamma.VectorInfo{
		Name:       opt.VectorName,
		DataType:   gamma.FLOAT,
		IsIndex:    true,
		Dimension:  int32(opt.D),
		StoreType:  opt.StoreType,
		StoreParam: string("{\"cache_size\": 2048}"),
	}
	table.VectorsInfos = append(table.VectorsInfos, vectorInfo)
	gamma.CreateTable(opt.Engine, &table)
}

func AddDocToEngine(docNum int) {
	for i := 0; i < docNum; i++ {
		var doc gamma.Doc
		for j := 0; j < len(opt.FieldsVec); j++ {
			field := &vearchpb.Field{
				Name: opt.FieldsVec[j],
				Type: opt.FieldsType[j],
			}
			profileStr := opt.Profiles[i*len(opt.FieldsVec)+j]
			if field.Type == vearchpb.FieldType_INT {
				tmp, _ := strconv.Atoi(profileStr)
				bytesBuffer := bytes.NewBuffer([]byte{})
				binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
				field.Value = bytesBuffer.Bytes()
			} else if field.Type == vearchpb.FieldType_LONG {
				tmp, _ := strconv.ParseInt(profileStr, 10, 64)
				bytesBuffer := bytes.NewBuffer([]byte{})
				binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
				field.Value = bytesBuffer.Bytes()
			} else {
				field.Value = []byte(profileStr)
			}
			doc.Fields = append(doc.Fields, field)
		}

		field := &vearchpb.Field{
			Name: opt.VectorName,
			Type: vearchpb.FieldType_VECTOR,
		}
		field.Value = make([]byte, opt.D*4)
		_, _ = opt.FeatureMmap.ReadAt(field.Value, int64(i*opt.D)*4)
		doc.Fields = append(doc.Fields, field)

		//for i := 0; i < opt.D; i++ {
		//	a := Float64frombytes(field.Value[i*4:])
		//	fmt.Println(a)
		//}

		byteArray := make([][]byte, 1)
		byteArray[0] = doc.Serialize()
		gamma.AddOrUpdateDocs(opt.Engine, byteArray)
	}
}

func Add() {
	file, err := os.OpenFile(opt.ProfileFile, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for idx := 0; idx < opt.AddDocNum; idx++ {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}

		docs := strings.Split(line, "\t")
		i := 0
		for _, doc := range docs {
			opt.Profiles[idx*len(opt.FieldsVec)+i] = doc
			i++
			if i > len(opt.FieldsVec)-1 {
				break
			}
		}
	}
	AddDocToEngine(opt.AddDocNum)
}

func Load() {
	gamma.Load(opt.Engine)
}

func Dump() {
	gamma.Dump(opt.Engine)
}

func GetDoc() {
	//tmp, _ := strconv.ParseInt("100", 10, 64)
	//bytesBuffer := bytes.NewBuffer([]byte{})
	//binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
	//docID := bytesBuffer.Bytes()
	docID := []byte("100")
	var doc gamma.Doc
	gamma.GetDocByID(opt.Engine, docID, &doc)
	fmt.Printf("doc [%v]", doc)
}

func Float64frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Search() {
	request := &vearchpb.SearchRequest{
		ReqNum:          1,
		TopN:            100,
		IsBruteSearch:   0,
		MultiVectorRank: 0,
		IndexParams:     "{\"metric_type\" : \"InnerProduct\", \"recall_num\" : 100, \"nprobe\" : 10, \"ivf_flat\" : 0}",
		L2Sqrt:          false,
	}
	request.VecFields = make([]*vearchpb.VectorQuery, 1)
	request.VecFields[0] = &vearchpb.VectorQuery{
		Name:     opt.VectorName,
		MinScore: 0,
		MaxScore: 1000,
	}

	request.VecFields[0].Value = make([]byte, opt.D*4)
	_, _ = opt.FeatureMmap.ReadAt(request.VecFields[0].Value, int64(0*opt.D)*4)

	//for i := 0; i < opt.D; i++ {
	//	a := Float64frombytes(request.VecFields[0].Value[i*4:])
	//	fmt.Println(a)
	//}

	response := &vearchpb.SearchResponse{}
	reqByte := gamma.SearchRequestSerialize(request)
	respByte, status := gamma.Search(opt.Engine, reqByte)

	if status.Code != 0 {
		fmt.Printf("code %v\n", status.Code)
		return
	}
	gamma.DeSerialize(respByte, response)

	fmt.Printf("%v\n", *(response.Results[0]))
}

func main() {
	Init()
	defer gamma.Close(opt.Engine)
	//Load()
	//return
	fmt.Println(opt.FeatureFile)
	CreteTable()
	Add()
	GetDoc()
	time.Sleep(time.Duration(1) * time.Second)
	status := gamma.GetEngineStatus(opt.Engine)
	fmt.Println(status)
	Search()

	Dump()
	opt.FeatureMmap.Close()
}
