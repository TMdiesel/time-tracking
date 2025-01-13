# 🚀 Time Tracker CLI

**Time Tracker CLI** は、タスクと時間を効率的に管理するための CLI ベースの時間管理ツールです。

## 主な機能

- プロジェクト・タスク管理
  プロジェクトとタスクの作成をする。

- 時間計測 (Time Tracking)
  タスクの開始・停止操作で作業時間を記録する。1 つのタスクのみをアクティブに保つことで重複記録を防止。

- 履歴の集計・表示
  日付範囲での作業履歴を確認する。

## インストール

### 必要条件

- Go: 1.21 以上
- OS: macOS, Linux, Windows

### インストール方法

```bash
git clone https://github.com/your-username/time-tracker-cli.git
cd time-tracker-cli
go build -o time-tracker cmd/main.go
echo 'export PATH=$PATH:$(pwd)' >> ~/.bashrc  # bash の場合
source ~/.bashrc
```

## 使い方

### 1. プロジェクトの作成

```bash
time-tracker project create --name "Rocket Development" --description "Building a next-gen rocket"
```

### 2. タスクの作成

```bash
time-tracker task create --name "Design engine" --description "Designing the propulsion system"
```

### 3. タスクの開始

```bash
time-tracker task start
```

- 同時にアクティブにできるタスクは 1 つのみです。

### 4. タスクの停止

```bash
time-tracker task stop
```

### 5. 作業履歴の表示

```bash
time-tracker entry list --from 2024-01-01 --to 2024-01-31
```
