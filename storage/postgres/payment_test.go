package postgres

import (
	"fmt"
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

		panic(err)
	}
	fmt.Println("+++++++++++++")

	pay := NewPaymentRepository(db)
	fmt.Println(pay)
	pays := pb.CreatePaymentRequest{
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		Amount:        30.00,
		PaymentMethod: "credit_card",
		PaymentStatus: "pending",
	}
	response, err := pay.CreatePayment(&pays)
	if err != nil {
		fmt.Println("_________", err)
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
		Id: "b8811cf2-9352-4dc3-9884-c7dabca7ab8b",
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
		Id:            "b271e2a3-f90d-4b51-bfd5-808bd65f1756",
		ReservationId: "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Amount:        753.5,
		PaymentMethod: "salom",
		PaymentStatus: "pending",
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
		Id: "e7c84f17-6a39-4889-92fe-0a9f3a3c062",
	}

	fmt.Println(pay)
	expected := pb.PaymentResponse{
		Id:            "e7c84f17-6a39-4889-92fe-0a9f3a3c0622",
		ReservationId: "a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2",
		Amount:        30.00,
		PaymentMethod: "credit_card",
		PaymentStatus: "pending",
	}

	response, err := pay.GetByIdPayment(&getPayment)
	if err != nil {
		fmt.Println("+++++++", err)
		panic(err)
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.PaymentResponse{})

	}
}

func TestAllPaymentRepository_GetAllPayment(t *testing.T) {
	db, err := Connection()
	if err != nil {

		panic(err)
	}

	payment := NewPaymentRepository(db)

	getAllPayments := pb.GetAllPaymentRequest{
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
