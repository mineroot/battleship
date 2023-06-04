package p2p

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	"sync"
	"time"
)

var (
	listener net.Listener
	conn     net.Conn
	wg       sync.WaitGroup
	done     = make(chan struct{})
)

func CreateServer() {
	wg.Add(1)
	defer wg.Done()
	var err error
	var lc net.ListenConfig
	var port int
	portMin, portMax := 49152, 65535
	for port = portMin; port <= portMax; port++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		listener, err = lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", port))
		cancel()
		if err == nil {
			break
		}
		log.Warn().Err(fmt.Errorf("p2p: unable to start listening on %d: %w", port, err)).Send()
	}
	if listener == nil {
		log.Error().Msg("p2p: unable to start listening")
	}
	defer listener.Close()
	log.Info().Msgf("p2p: server started on %d", port)
	for {
		conn, err = listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if conn != nil {
			log.Info().Msgf("p2p: connection accepted from %s", conn.RemoteAddr())
			return
		}
		log.Warn().Err(fmt.Errorf("p2p: unable to accept conn: %w", err)).Send()
	}
}

func CreateClient(addr string) {
	var err error
	for {
		select {
		case <-done:
			return
		default:
		}
		conn, err = net.DialTimeout("tcp", addr, time.Second)
		if err != nil {
			log.Warn().Err(fmt.Errorf("p2p: unable to dial: %w", err)).Send()
		}
		if conn != nil {
			log.Info().Msgf("connection established with %s", conn.RemoteAddr())
			break
		}
		time.Sleep(time.Second)
	}

	_, _ = conn.Write([]byte("hi\n")) // TODO: for test
}

func Close() {
	close(done)
	if listener != nil {
		listener.Close()
	}
	if conn != nil {
		conn.Close()
	}
	wg.Wait()
}
