package main

import "hdfs/hdfs"

// import "tidydfs/tdfs"

// "fmt"
// "tdfs"
// "runtime"
// "sync"

const DN3_DIR string = "./datanode"
const DN3_LOCATION string = "http://localhost:11093"
const DN3_CAPACITY int = 100

func main() {
	var dn3 hdfs.DataNode
	dn3.DATANODE_DIR = DN3_DIR
	
	dn3.Reset()
	dn3.SetConfig(DN3_LOCATION, DN3_CAPACITY)

	dn3.Run()
}