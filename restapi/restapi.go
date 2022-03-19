package restapi

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"test01/datalayer"
	"time"

	"github.com/gorilla/mux"
)

func RunApi(endpoint string, db datalayer.SQLHandler) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx, db); err != nil {
		log.Printf("failed to serve:+%v\n", err)
		return err
	}

	return nil
}

func RunApiOnRouter(r *mux.Router, db datalayer.SQLHandler) {
	handler := newPromotionRestApiHandler(db)
	apiRouter := r.PathPrefix("/promotions").Subrouter()
	apiRouter.Methods("Get").Path("/{id}").HandlerFunc(
		handler.returnSinglePromotion)
}

func serve(ctx context.Context, db datalayer.SQLHandler) (err error) {
	handler := newPromotionRestApiHandler(db)
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/promotions/{id}", handler.returnSinglePromotion).Methods("GET")

	srv := &http.Server{
		Addr:    ":6999",
		Handler: myRouter,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
