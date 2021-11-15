package rdb

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// DB設定情報
type Config struct {
	MasterHost      string
	SlaveHost       string
	Name            string
	Port            int
	User            string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// RDBKind - RDBの接続種別
type RDBKind uint32

const (
	// RDBKindMaster - マスター
	RDBKindMaster RDBKind = iota
	// RDBKindSlave - スレーブ
	RDBKindSlave
)

// Exector - DB操作
type Exector struct {
	config *Config
}

// DBとロジックを分離する
type SessionScope struct {
	Conn *gorm.DB
}

// NewExector - 生成
func NewExector(config *Config) *Exector {
	return &Exector{
		config: config,
	}
}

// Process - 通常のアクセス
func (s *Exector) Process(kind RDBKind, f func(*SessionScope) error) error {
	conn, err := s.tryConnect(kind)
	if err != nil {
		return err
	}
	// Close：DBとのコネクション開放
	// deferでreturnする直前に実行し、Close忘れを防止
	defer conn.Close()

	// セッションスコープ使用でビジネスロジックとDBを切り離し
	ss := &SessionScope{Conn: conn}
	if err := f(ss); err != nil {
		return err
	}
	return nil
}

// ProcessTX - トランザクションを利用したアクセス
func (s *Exector) ProcessTX(kind RDBKind, f func(*SessionScope) error) error {
	conn, err := s.tryConnect(kind)
	if err != nil {
		return err
	}
	defer conn.Close()

	// トランザクション開始
	conn = conn.Begin()
	processed := false
	// これがあることによって関数実行中にパニックが発生してもRollbackするようになっている
	defer func() {
		if !processed {
			log.Fatal("process tx fail")
			conn = conn.Rollback()
		}
	}()

	ss := &SessionScope{Conn: conn}
	if err := f(ss); err != nil {
		// エラーが発生した場合はロールバック
		conn = conn.Rollback()
		processed = true
		return err
	}
	// 正常な場合は関数実行後にコミットする
	conn = conn.Commit()
	processed = true
	return nil
}

// tryConnect - 接続
func (s *Exector) tryConnect(kind RDBKind) (*gorm.DB, error) {
	host := s.config.MasterHost
	if kind == RDBKindSlave {
		host = s.config.SlaveHost
	}
	// コネクション確立
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.config.User, s.config.Password, host, s.config.Port, s.config.Name))
	if err != nil {
		return nil, err
	}
	// https://tutuz-tech.hatenablog.com/entry/2020/03/24/170159
	// 最大の接続数
	// 最大接続数を超えるリクエストがあると、どれかの接続がアイドル状態に戻るまで待機する
	db.DB().SetMaxOpenConns(s.config.MaxOpenConns)
	// アイドル状態ので残しておく接続数
	// 0にすると、毎回接続を新規作成するので、パフォーマンスが悪くなる
	db.DB().SetMaxIdleConns(s.config.MaxIdleConns)
	// 接続を再利用できる最大の期間を設定する
	db.DB().SetConnMaxLifetime(s.config.ConnMaxLifetime)
	return db, nil
}
