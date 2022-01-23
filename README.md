# Two Choices

あなたはどっち派？という2択を選択する暇つぶしアプリです。
https://www.twochoices.site

実務で経験した技術の復習及び、Github Actions、Kubernetes等の技術を勉強することを目的として作成しました。
学習の記録として、ソースコード内各所に不要なコメントが残っておりますが、ご容赦下さい。

## 使用技術

#### 【サーバ】
- Golang
  - protocol buffersを使用したAPIの定義ファイルを作成
  - protocol buffersの定義に従いAPIを実装
  - Firebase Authenticationを使用した認証機能の実装
    (学習目的であり、ログイン後の機能は未実装)
- GRPC
#### 【クライアント】
- Nuxt.js
  - Vuetifyを使用し、画面フレームUIを実装
  - メイン画面実装
  - Typescript, Sass等を使用し、ページ内の動的なデザイン変更、及びデータ取得処理を実装
- Typescript
#### 【データ】
- MySQL
- Redis
#### 【CI/CD】
- Github Actions
  - Golangのサーバ側テスト実行
  - サーバ、クライアント、DBのコンテナイメージをビルドし、GCP Container Registoryへpush
- GCP
- Kubernetes
  - 複数環境の構築を想定し、kustomizeを使用
#### 【その他】
- Firebase Authentication
- SendGrid

## インフラ構成
- DBは仮の為、ServerのPod内にMySQLコンテナを立てており、データの永続化は行っておりません
![two_choices](https://user-images.githubusercontent.com/64344576/147831621-8c9d2653-d07b-4af5-acda-f0437f393404.png)
