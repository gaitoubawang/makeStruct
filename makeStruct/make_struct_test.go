package makeStruct

import (
	"encoding/json"
	"reflect"
	"testing"
)

type OrderList struct {
	OrderDetails []*OrderDetail `json:"order_details"`
	Total        *int           `json:"total"`
	Meta         []string       `json:"meta"`
}

type OrderDetail struct {
	ID          string       `json:"id"`
	Goods       *GoodsDetail `json:"goods"`
	PaymentInfo PaymentInfo  `json:"payment_info"`
}

type GoodsDetail struct {
	Names []string `json:"names"`
}

type PaymentInfo struct {
	PayList []string `json:"pay_list"`
}

func TestCasesOfMakeStruct(t *testing.T) {
	var total = 0
	type args struct {
		target interface{}
	}
	type want struct {
		result interface{}
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "TestCasesOfMakeStruct - normal",
			args: args{
				target: &OrderList{},
			},
			want: want{
				result: &OrderList{
					OrderDetails: []*OrderDetail{},
					Total:        &total,
					Meta:         make([]string, 0),
				},
			},
		},
		{
			name: "TestCasesOfMakeStruct - had object",
			args: args{
				target: &OrderList{
					OrderDetails: []*OrderDetail{
						{
							ID: "123456",
						},
					},
					Total: &total,
					Meta:  make([]string, 0),
				},
			},
			want: want{
				result: &OrderList{
					OrderDetails: []*OrderDetail{
						{
							ID: "123456",
							Goods: &GoodsDetail{
								Names: make([]string, 0),
							},
							PaymentInfo: PaymentInfo{
								PayList: make([]string, 0),
							},
						},
					},
					Total: &total,
					Meta:  make([]string, 0),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := json.Marshal(tt.args.target)
			t.Logf("make before: %s", string(data))
			MakeStruct(tt.args.target)
			data, _ = json.Marshal(tt.args.target)
			t.Logf("make after: %s", string(data))
			if !reflect.DeepEqual(tt.args.target, tt.want.result) {
				t.Errorf("%s , not equal", tt.name)
				data, _ = json.Marshal(tt.want.result)
				t.Logf("make want: %s", string(data))
			}
		})
	}
}
