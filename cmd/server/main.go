package main

import (
	delivery_grpc "clean-arch-go-grpc/internal/delivery/grpc"
	"clean-arch-go-grpc/internal/delivery/product_pb"
	"clean-arch-go-grpc/internal/repository"
	"clean-arch-go-grpc/internal/usecase"
	"clean-arch-go-grpc/pkg/gorm"
	"clean-arch-go-grpc/pkg/logrus"
	viperPkg "clean-arch-go-grpc/pkg/viper"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/spf13/viper"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	logger "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	ctxKeyWaitGroup = "waitGroup"
)

var (
	Ctx    context.Context
	Cancel context.CancelFunc
)

func init() {
	Ctx, Cancel = context.WithCancel(context.WithValue(context.Background(), ctxKeyWaitGroup, new(sync.WaitGroup)))
}

func GetWaitGroupInCtx(ctx context.Context) *sync.WaitGroup {
	if wg, ok := ctx.Value(ctxKeyWaitGroup).(*sync.WaitGroup); ok {
		return wg
	}

	return nil
}

func main() {
	viper := viperPkg.NewViper()
	logrus := logrus.NewLogger()
	gorm := gorm.NewDatabase(viper, logrus)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pr := repository.NewProductRepository(gorm, logrus)

	pu := usecase.NewProductUsecase(logrus, pr)

	delivery_grpc.NewProductServerGrpc(server, logrus, pu)

	logrus.Infof("server listening :%d", 3000)
	go func() {
		err = server.Serve(lis)
		if err != nil {
			fmt.Println("Unexpected Error", err)
		}
	}()

	go StartSwaggoServer(Ctx, viper)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Infof("shutdown server ...")

}

func StartSwaggoServer(ctx context.Context, viper *viper.Viper) {
	// waitGroup add
	wg := GetWaitGroupInCtx(ctx)
	wg.Add(1)
	defer wg.Done()

	// New Mux
	gwmux := runtime.NewServeMux()
	// var opts []grpc.DialOption
	// if need certificate
	// if cfg.System.OpenPem {
	// 	opts = []grpc.DialOption{grpc.WithTransportCredentials(config.GetServerCred(&cfg.Cert))}
	// } else {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// }

	host := viper.GetString("swaggo.host")
	port := viper.GetString("swaggo.port")
	baseDir := viper.GetString("swaggo.base-dir")

	endpoint := strings.Join([]string{host, port}, ":")

	// registry handler endpoint
	err := product_pb.RegisterProductHandlerHandlerFromEndpoint(ctx, gwmux, endpoint, opts)
	if err != nil {
		logger.Fatal().Msgf("registry http server error:%v", err)
	}

	// add swagger api
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	dir := filepath.Join(baseDir, "internal/delivery/swagger")
	handler := http.FileServer(http.Dir(dir))
	mux.Handle("/api/", http.StripPrefix("/api/", handler))
	logger.Info().Msg("add swagger api path:" + dir)

	go func() {
		logger.Info().Msg("starting swaggo server at..." + port)
		// if cfg.System.OpenPem {
		// 	err = http.ListenAndServeTLS(endpoint, cfg.Cert.ServerPemPath, cfg.Cert.ServerKeyPath, handler)
		// 	// err = http.ListenAndServeTLS(endpoint, insecure.certPEM, insecure.keyPEM, handler)
		// } else {
		err = http.ListenAndServe(":"+port, mux)
		// }
		if err != nil {
			logger.Fatal().Msgf("swaggo server listen error:%v", err)
		}
	}()

	<-ctx.Done()
}
