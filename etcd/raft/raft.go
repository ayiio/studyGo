package raft

import (
	"sync"
)

//1.模拟三节点选举
//2.改造成分布式选举，加入RPC调用
//3.自动选主，日志复制

//定义三节点常量
const raftCount = 3

//leader对象
type Leader struct {
	Term     int //任期
	LeaderId int //leader编号
}

//Raft声明
type Raft struct {
	mu              sync.Mutex //锁
	me              int        //节点编号
	currentTerm     int        //当前任期
	votedFor        int        //为哪个节点投票
	state           int        //状态，0-follower，1-candidate， 2-leader
	lastMessageTime int64      //发送最后一条数据的时间
	currentLeader   int        //当前节点的leader
	message         chan bool  //节点间信息通道
	electCh         chan bool  //选举通道
	heartBeat       chan bool  //心跳信号通道
	heartbeatRe     chan bool  //返回心跳信号通道
	timeout         int        //超时时间，随机
}

// Term-0未上任， LeaderId -1没有编号
var leader = Leader{
	Term:     0,
	LeaderId: -1,
}

func Vote() {

}
