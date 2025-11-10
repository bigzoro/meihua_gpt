# 梅花易数助手

一个前后端分离的梅花易数占卜示例项目，前端使用 Vue 3 + Vite + Pinia，后端基于 Gin + GORM + SQLite 实现快速起卦与历史记录功能。

## 项目结构

```
├── backend          # Go 服务端
│   ├── main.go      # 应用入口，Gin 路由与服务启动
│   ├── handlers     # HTTP 处理器
│   ├── models       # GORM 数据模型
│   ├── services     # 梅花易数推演逻辑
│   └── database     # 数据库初始化工具
└── frontend         # Vue 3 前端项目
    ├── src
    │   ├── api      # Axios 客户端
    │   ├── components
    │   ├── stores   # Pinia 状态管理
    │   └── styles
    ├── index.html
    └── vite.config.js
```

## 后端运行

```bash
cd backend
go run .
```

默认监听 `http://localhost:8080`，使用 SQLite 数据库 `meihua.db`（可通过环境变量 `MEIHUA_DB_PATH` 修改）。

## 前端运行

```bash
cd frontend
npm install
npm run dev
```

默认通过 Vite 启动在 `http://localhost:5173`，并调用后端 `/api` 接口。

## API 简介

- `POST /api/divinations`：创建新的占卦，入参包含 `subject`、`method`。
- `GET /api/divinations`：获取占卦历史列表（按时间倒序）。
- `GET /api/divinations/:id`：查看单条记录详情。

## 许可证

本项目仅作为示例，使用 MIT License 发布。
