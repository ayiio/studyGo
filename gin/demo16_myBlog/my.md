## 项目目录规范
   + blogger
     + controller：页面控制
     + dao：数据层(sql)
     + model：实体层(结构体)
     + service：业务逻辑
     + static：css/js
     + utils：工具包
     + views：HTML模板
     + main.go：项目入口，定义路由
## 数据表
   + article 文章表
     ```sql
     create table article (
        id bigint(20) primary key not null auto, --- 文章ID
        category_id bitint(20) not null, --- 分类ID
        content longtext not null, --- 文章内容
        title varchar(1024) not null, --- 文章标题
        view_count int(255) not null, --- 阅读次数
        comment_count int(255) not null, --- 评论次数
        username varchar(128) not null, --- 作者
        status int(10) not null 1, --- 状态，正常为1
        summary varchar(255) not null, --- 文章摘要
        create_time timestamp CURRENT_TIME, --- 发布时间
        update_time timestamp --- 更新时间
     ) engine=InnoDB, charset=utf8;
     ```
   + category 分类表
     ```sql
     create table category (
        id bigint(20) primary key not null auto, --- 分类ID
        category_name varchar(255) not null, --- 分类名字
        category_no int(10) not null,  --- 分类排序
        create_time timestamp CURRENT_TIME,
        update_time timestamp
     ) engine=InnoDB, charset=utf8;
     ```
   + comment 评论表
     ```sql
     create table common (
        id bigint(20) primary key not null auto, --- 评论ID
        content text not null, --- 评论内容
        username varchar(64) not null, --- 评论作者
        create_time timestamp CURRENT_TIME not null, --- 评论发布时间
        status int(255) not null 1, --- 评论状态：0， 删除：1 
        article_id bigint(20) --- 文章ID
     ) engine=InnoDB, charset=utf8;
     ```
   + leavemsg 留言表
     ```sql
     create table leavemsg (
        id bigint(20) primary key not null auto, --- 留言ID
        username varchar(255) not null, --- 留言作者
        email varchar(255) not null, --- 作者邮箱
        content text not null, --- 留言内容
        create_time timestamp CURRENT_TIME,
        update_time timestamp
     ) engine=InnoDB, charset=utf8;
     ```
## 实体类
   + 结构体定义
      + 分类的结构体
      + 文章结构体(不包含大文本的文章内容)
      + 文章详情页结构体
      + 文章上下页
## 数据层
   + 数据库相关
      + init() 数据库初始化函数
      + 分类相关的操作(添加，查询一或多或所有)
      + 文章相关的操作(投稿添加文章，首页查询所有文章，根据文章id查看内容)
