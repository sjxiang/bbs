docker run -d           --env COCKROACH_DATABASE=bbs           --env COCKROACH_USER=root           --env COCKROACH_PASSWORD=123456           --name=roach-single           -p 26257:26257 -p 8080:8080           -v "roach-single:/cockroach/cockroach-data"                  cockroachdb/cockroach:latest start-single-node


./cockroach sql --insecure --host=localhost:26257
docker exec -it 1e956c939715 bash




社交网站

注册 / 登录


时间线 timeline / Feed 流

帖子 Posts
    
    - 关注 
    - 全局


通知 Notifications


搜索 Search


评论 Comments


前端 https://watercss.kognise.dev/
