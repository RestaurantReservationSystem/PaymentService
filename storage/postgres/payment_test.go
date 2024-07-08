package postgres

import (
	"fmt"
	"log"
	"log/slog"
	pb "payment_service/genproto"
	"reflect"
	"testing"
)

//	func TestConnection(t *testing.T) {
//		db, err := Connection()
//		if err != nil {
//			panic(err)
//		}
//
// }
func TestPaymentRepository_CreatePayment(t *testing.T) {
	db, err := Connection()
	if err != nil {
		log.Fatal("error is connection db")
	}
	repo := NewPaymentRepository(db)
	request := pb.CreatePaymentRequest{PaymentStatus: "pending", PaymentMethod: "card", Amount: 4.5, ReservationId: "1234alk;sdl;kf"}
	response, err := repo.CreatePayment(&request)
	if err != nil {
		fmt.Println("+++++++", err)
		//log.Fatalf("error internal server error ->", err.Error())
	}
	expected := &pb.Void{}
	if reflect.DeepEqual(response, expected) {
		slog.StringValue("message is success")
func TestPaymentRepository_CreatePayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	pay := NewPaymentRepository(db)
	fmt.Println(pay)
	track := pb.CreatePaymentRequest{
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Amount:        345.6,
		PaymentMethod: "salom",
		PaymentStatus: "jksdlkja;fsjk",
	}
	response, err := pay.CreatePayment(&track)
	if err != nil {
		fmt.Println("_________")
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.ReservationResponse{})

	}

}
func TestPaymentRepository_DeletePayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	pay := NewPaymentRepository(db)
	fmt.Println(pay)
	payDelete := pb.IdRequest{
		Id: "cfd06c0a-aa56-4755-9e1e-3779647f916e",
	}
	response, err := pay.DeletedPayment(&payDelete)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.PaymentResponse{})

	}

}
func TestPaymentRepository_UpdatePayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	pay := NewPaymentRepository(db)
	fmt.Println(pay)
	paymentUpdate := pb.UpdatePaymentRequest{
		Id:            "cfd06c0a-aa56-4755-9e1e-3779647f916e",
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Amount:        753.5,
		PaymentMethod: "salom",
		PaymentStatus: "+++++++++++++++++++++++",
	}
	response, err := pay.UpdatePayment(&paymentUpdate)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.PaymentResponse{})

	}
}

func TestPaymentRepository_GetPayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	pay := NewPaymentRepository(db)

	getPayment := pb.IdRequest{
		Id: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
	}

	fmt.Println(pay)
	expected := pb.PaymentResponse{}

	response, err := pay.GetByIdPayment(&getPayment)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.PaymentResponse{})

	}
}

func TestAllPaymentRepository_GetPayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	payment := NewPaymentRepository(db)

	getAllPayments := pb.GetAllPaymentRequest{
		Id:            "cfd06c0a-aa56-4755-9e1e-3779647f916e",
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Amount:        753.5,
		PaymentMethod: "salom",
		PaymentStatus: "+++++++++++++++++++++++",
	}

	fmt.Println(payment)
	expected := pb.PaymentsResponse{}

	response, err := payment.GetAllPayment(&getAllPayments)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.PaymentsResponse{})

	}
}
