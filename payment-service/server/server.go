
package main

import (
	"context"
	"os"
	gpay "vue-golang-payment-app/payment-service/proto"
	payjp "github.com/payjp/payjp-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement sa
type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {
	// PAIの初期化
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)

	// 支払いをします。第一引数に支払い金額、第二引数に支払いの方法や設定を入れます。
	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		// 現在はjpyのみサポート
		Currency: "jpy",
		CardToken: req.Token,
		Capture: true,
		Description: req.Name + ":" + req.Description,

	}

}
