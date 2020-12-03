# Make Struct

make a struct with default value.
it can avoid return nil in http json body

## Result display

>there is a order list prepare to return

```shell
type OrderList struct {
    OrderDetails []*OrderDetail   `json:"order_details"`
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
```
>make before

```shell
{
    "order_details":nil,
    "total":nil,
    "meta":nil
}
```

```shell
{
    "order_details":[
        {
            "id":"123456",
            "goods":null,
            "payment_info":{
                "pay_list":null
            }
        }
    ],
    "total":0,
    "meta":[

    ]
}
```

>make after

```shell
{
    "order_details":[],
    "total":0,
    "meta":[]
}

{
    "order_details":[
        {
            "id":"123456",
            "goods":{
                "names":[

                ]
            },
            "payment_info":{
                "pay_list":[

                ]
            }
        }
    ],
    "total":0,
    "meta":[

    ]
}
```

## Example

```shell
func GetOrderList() *OrderList {
	result := &OrderList{}
	makeStruct.MakeStruct(result)
	return result
}
```

## FEAT

> add tag to control fill the field or not
