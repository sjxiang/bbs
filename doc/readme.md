

# 企业里项目开发的流程

1. 编写需求文档
2. 原型与 UI 设计
3. 确定数据库的表、字段以及接口地址和数据
4. 同时进行：前端 mock 开发、后端接口开发
5. 接口开发完成后，将 mock 地址，换为接口地址
6. 测试、上线部署



# 数据库

## categories, 分类表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
name(名称) 
    varchar, 不为 null
rank(排序) 
    integer, 无符号, 不为 null, 默认1



## courses 课程表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
category_id(分类编号) 
    integer, 无符号, 不为 null, index 索引
user_id(用户编号)
    integer, 无符号, 不为 null, index 索引
name(名称) 
    varchar, 不为 null
image(课程图片)
    varchar
recommended(是否推荐)
    tinyint, 无符号, 不为 null, 默认0, 0 表示不推荐, 1 表示推荐, index 索引
introductory(是否为入门课程)
    tinyint, 无符号, 不为 null, 默认0, 0 表示不是入门课程, 1 表示是入门课程, index 索引
content(课程内容)
    text
likes_count(课程的点赞数)
    integer, 无符号, 不为 null, 默认0
chapters_count(课程的章节数量)
    integer, 无符号, 不为 null, 默认0


## chapters 章节表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
course_id(课程编号)
    integer, 无符号, 不为 null, index 索引
title(课程标题)
    varchar, 不为 null
content(课程内容)
    text
video(课程视频)
    varchar
rank(排序)
    integer, 无符号, 不为 null, 默认1


## users 用户表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
email(邮箱)
    varchar, 不为 null, unique 索引
username(用户名)
    varchar, 不为 null, unique 索引
nickname(昵称)
    varchar, 不为 null
password(密码)
    varchar, 不为 null
avatar(头像)
    varchar
sex(性别)
    tinyint, 无符号, 不为 null, 默认0, 0 表示不选择, 1 表示男, 2 表示女
company(公司.学校名)
    varchar
intro(简介)
    text
role(用户组)
    tinyint, 无符号, 不为 null, index 索引。默认0, 0 表示普通用户, 100 表示管理员



## likes 点赞表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
course_id(课程编号)
    integer, 无符号, 不为 null, index 索引
user_id(用户编号)
    integer, 无符号, 不为 null, index 索引



## settings 系统设置表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
name(项目名称)
    varchar
icp(备案号)
    varchar
copyright(版权信息)
    varchar



# posts 帖子表

id(编号)
    integer, 无符号, 不为 null, 主键, 自增
title(标题)
    varchar, 不为 null
content(内容)
    text
user_id(用户编号)
    integer, 无符号, 不为 null, index 索引
is_draft
    tinyint, 无符号, 不为 null, 默认0, 0 表示不是草稿, 1 表示是草稿
created_at    
    datetime, 不为 null, 默认当前时间
updated_at
    datetime, null