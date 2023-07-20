### Raft介绍
  Raft提供了一种在计算系统集群中分布状态机的通用方法，确保集群中的每个节点都同意一系列相同的状态转换。    
  它有许多开源参考实现，具有Go，C++，Java等语言的完整规范实现。    
  一个Raft集群包含了若干个服务器节点，通常是5个，这允许整个系统容忍2个节点的失败，每个节点处于以下三种状态之一：    
      **follower（跟随者）**，所有节点都以follower的状态开始，如果没有收到leader消息则会变成candidate状态    
      **candidate（候选人）**，会向其他节点“拉投票”，如果得到大部分的票则成为leader，这个过程叫做leader选举（leader election）    
      **leader（领导者）**，所有对系统的修改都会先经过leader    
### Raft一致性算法
  Raft通过选出一个leader来简化日志副本的管理，例如日志项（log entry）只允许从leader流向follower  
  基于leader的方法，raft算法可以分解成三个子问题：   
      **leader election**（领导选举），原来的leader挂掉后，必须选出一个新的leader   
      **log replication**（日志复制），leader从客户端接收日志，并复制到整个集群中   
      **safety**（安全性），如果有任意的server将日志项放回到状态机中，那么其他的server只会回放相同的日志项   
### raft动画演示
  地址：http://thesecretlivesofdata.com/raft   
  介绍简单版的领导者选取和日志复制的过程 + 介绍详细版的领导者选举和日志复制的过程 + 如果遇到网络分区（脑裂），raft算法如何恢复网络一致性
