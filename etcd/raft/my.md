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
### leader election（领导选举）
  raft使用一种心跳机制来触发领导人选举  
  当服务器程序启动时，节点都是follower（跟随者）身份  
  如果一个跟随者在一段时间里没有接收到任何消息，也就是选举超时，然后它会认为系统中没有可用的领导者，将开始进行选举以选出新的领导者   
  要开始一次选举过程，follower会给当前term加1并且转换成candidate状态，然后它会并行向集群中的其他服务器节点发送请求投票的RPCs来给自己投票   
  候选人的状态维持直到在以下一个条件发生时：   
      它自己赢得了这次选举  
      其他服务器成为了领导者   
      一段时间之后没有任何一个获胜  
### Log replication（日志复制）  
  当选出leader后，它会开始接收客户端请求，每个请求会带有一个指令，可以被回放到状态机中  
  leader吧指令追加成一个log entry，然后通过appendentries RPC并行发送给其他的server，当该entry被多个server复制后，leader会把该entry回放到状态机中，然后把结果返回给客户端  
  当follower宕机或者运行较慢时，leader会无限重发appendentries 给这些follower，直到所有的follower都复制了这个log entry   
  raft的log replication要保证如果两个log entry有相同的index和term，那么它们存储相同的指令  
  leader在一个特定的term和index下，只会创建一个log entry   
  
