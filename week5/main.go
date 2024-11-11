package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"week5/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Failed to run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// Interrupt signal 신호 처리를 위한 context 생성
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("failed to listen port %v", err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with %s", url)
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
