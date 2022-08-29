package main

import (
	"os"

	"github.com/kelein/gokit/zlog"
)

func testStandardLogger() {
	defer zlog.Sync()
	zlog.Debug("test pkg zlog")
	zlog.Info("test pkg zlog")
	zlog.Warn("test pkg zlog")
	zlog.Error("test pkg zlog")
	// zlog.Fatal("test pkg zlog")
}

func testRotateLogger() {
	f1, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	f2, err := os.OpenFile("fatal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	tops := []zlog.TeeOption{
		{
			Rotation: zlog.DefaultRotation(f2.Name()),
			LevelFunc: func(l zlog.Level) bool {
				return l >= zlog.ErrorLevel
			},
		},
		{
			Rotation: zlog.DefaultRotation(f1.Name()),
			LevelFunc: func(l zlog.Level) bool {
				return l >= zlog.DebugLevel && l < zlog.ErrorLevel
			},
		},
	}

	logger := zlog.NewRotateLogger(tops)
	zlog.ResetDefault(logger)
	defer zlog.Sync()

	zlog.Debug("test pkg zlog")
	zlog.Info("test pkg zlog")
	zlog.Warn("test pkg zlog")
	zlog.Error("test pkg zlog")
	// zlog.Fatal("test pkg zlog")
}

func main() {
	zlog.Info("testStandardLogger ...")
	testStandardLogger()

	// zlog.Sugar().Info("test sugarLogger ...")
	// zlog.Sugar().Infow("test sugarLogger", "err", errors.New("test sugar"))

	// zlog.Info("testRotateLogger ...", zlog.String("ts", time.Now().String()))
	// testRotateLogger()
}
