package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
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

func main() {
	//过程，有3个节点，最初都是follower
	//如果candidate状态，进行投票拉票
	//产生Leader

	//创建3个节点
	for i := 0; i < raftCount; i++ {
		Make(i)
	}
	for {

	}

}

//创建节点
func Make(me int) *Raft {
	rf := &Raft{
		me:            me,
		votedFor:      -1, //节点刚创建，-1谁都不投
		state:         0,  //0-follower
		timeout:       0,  //无超时
		currentLeader: 0,  //无领导
	}
	rf.setTerm(0) //节点任期
	//初始化通道
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	//设置投票用的随机种子
	rand.Seed(time.Now().UnixNano())

	//选举协程
	go rf.election()

	//心跳检测协程
	go rf.sendLeaderHeartBeat()

	return rf
}

//设置任期
func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

//选举
func (rf *Raft) election() {
	//设置标记，判断是否选出Leader
	var result bool
	for {
		//设置超时
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		case <-time.After(time.Duration(timeout) * time.Millisecond): //延迟等待一毫秒
			fmt.Println("当前节点状态：", rf.state)
		}
		result = false
		for !result {
			//选主逻辑
			result = rf.election_one_rand(&leader)
		}
	}
}

//150-300随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

//获取当前时间 - 发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond) //转换毫秒数
}

//实现选主逻辑
func (rf *Raft) election_one_rand(leader *Leader) bool {
	//定义超时
	var timeout int64
	timeout = 100
	//投票数量
	var vote int
	//心跳是否开始
	var triggerHeartBeat bool
	//时间
	var last = millisecond()
	//用于返回
	success := false

	//当前节点变为candidate
	rf.mu.Lock()
	rf.toCandidate() //修改状态为candidate
	rf.mu.Unlock()
	fmt.Println("become the candidate, start electing leader")
	for {
		//遍历所有节点，拉取选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me { //自己已投
				go func() {
					if leader.LeaderId < 0 { //没有Leader时才投票
						//设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		//设置投票数量
		vote = 1 //有自己的1票
		//遍历节点加选票
		for i := 0; i < raftCount; i++ {
			//计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					//投票数量加1
					vote++
					//若选票个数大于节点个数一半，才能成功成为Leader
					success = vote > raftCount/2
					if success && !triggerHeartBeat { //有临时决定的leader且没发起心跳信号
						//变换成主节点，选主成功
						rf.mu.Lock()
						rf.toLeader()
						rf.mu.Unlock()
						//开始触发心跳信号检测
						triggerHeartBeat = true
						//由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println("leader已选出,为: ", rf.me)
						fmt.Println("leader开始发出心跳信号")
					}
				}
			}
		}

		//leader选出后，再做检验
		//若不超时，且票数大于一半，则选举没有问题
		if timeout+last < millisecond() || (vote > raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			//等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

//修改状态
func (rf *Raft) toCandidate() {
	rf.state = 1                   //状态
	rf.setTerm(rf.currentTerm + 1) //任期
	rf.votedFor = rf.me            //给自己投票
	rf.currentLeader = -1          //当前领导
}

//设置领导
func (rf *Raft) toLeader() {
	rf.state = 2             //状态
	rf.currentLeader = rf.me //领导
}

//设置leader节点发送心跳信号
//同时保证数据一致性 - 同步， 检测从者的状态
func (rf *Raft) sendLeaderHeartBeat() {
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntryImpl()
		}
	}
}

//返回leader确认信号
func (rf *Raft) sendAppendEntryImpl() {
	//主节点不返回确认信号给自己
	if rf.currentLeader == rf.me {
		//此时是leader
		var success_count = 0 //记录确认信号的节点个数
		for i := 0; i < raftCount; i++ {
			//设置确认信号
			if i != rf.me {
				go func() {
					rf.heartbeatRe <- true
				}()
			}
		}
		//计算确认信号的返回个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_count++
					if success_count > raftCount/2 {
						fmt.Println("选举投票成功，心跳信号ok")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}

