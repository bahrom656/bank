package types

//Money представляет собой денежную сумму в минимальных единицах {центы, копейки, дирамы, и т.д.}
type Money int64

//Status представляет собой статус платежа
type Status string

//Category представляет собой категорию, в котором был совершен платёж
type Category string

//Предопределенные статусы платежей
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
