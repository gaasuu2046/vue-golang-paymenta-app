package handler

import (
	"context"
	"net/http"
	"strconv"
	"vue-golang-payment-app/backend-api/db"
	"vue-golang-payment-app/backend-api/domain"
	gpay "vue-golang-payment-app/payment-service/proto"

	"google.golang.org/grpc"
)

var addr = "localhost:50051"

// context handleRawConn
func Charge(c Context) {
	t := domain.Payment{}
	c.Bind(&t)

	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// idから item情報取得
	res, err := db.SelectItem(int64(identifier))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	//grpc サーバーに送るrequestを作成
	greq := &gpay.PayRequest{
		Id:          int64(identifier),
		Token:       t.Token,
		Name:        res.Name,
		Description: res.Description,
		Amount:      res.Amount,
	}

	// IPアドレスとport(50051)を指定して、サーバーと通信する。
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)

	//gRPCマイクロサービスにrequest送信をする
	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
