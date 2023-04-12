# 練習使用 docker 去建立連接資料庫

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

這邊主要練習的部分為如何將 golang api 讓他人可以直接 clone 下來使用，其中含將資料庫的框架也處理好，會需要用到 sql init sqripts 來初始化資料庫的 table 和一些 user 權限的問題

## 使用環境及工具

採用Golang語言的Echo框架開發RESTful API，使用PostgreSQL作為資料庫

## 如何運行該專案(使用docker-compose)

### Clone 專案

```
git clone https://github.com/jerry88819/docker_database.git
```

### 運行專案

```
# 在本專案的根目錄下執行以下指令即可
# -d 代表背景運行(Optional)
docker-compose up -d
```

## License

This project is open source and available under the [MIT License](https://opensource.org/licenses/MIT).
