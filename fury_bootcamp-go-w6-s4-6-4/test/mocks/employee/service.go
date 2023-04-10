package employee

import (
	"context"
	"errors"
	"strconv"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockServiceEmployee struct {
	DataMock     []domain.Employee
	DatamockIBO  []domain.ReportInBO
	Err          string
	MethodCalled bool
}

func (ms *MockServiceEmployee) Save(ctx context.Context, cardNumberId string, name string, lastname string, wharehouseId int) (domain.Employee, error) {
	ms.MethodCalled = true
	for _, value := range ms.DataMock {
		if value.CardNumberID == cardNumberId {
			return domain.Employee{}, errors.New("The card_number_id already exists")
		}
	}
	if ms.Err != "" {
		return domain.Employee{}, errors.New(ms.Err)
	}
	var newEmployee domain.Employee
	var id int = 1
	if len(ms.DataMock) > 0 {
		id = ms.DataMock[len(ms.DataMock)-1].ID + 1
	}
	newEmployee.ID = id
	newEmployee.CardNumberID = cardNumberId
	newEmployee.FirstName = name
	newEmployee.LastName = lastname
	newEmployee.WarehouseID = wharehouseId

	ms.DataMock = append(ms.DataMock, newEmployee)

	return newEmployee, nil
}

func (ms *MockServiceEmployee) GetAllEmployees(ctx context.Context) ([]domain.Employee, error) {
	ms.MethodCalled = true
	if ms.Err != "" {
		return []domain.Employee{}, errors.New(ms.Err)
	}
	return ms.DataMock, nil
}

func (ms *MockServiceEmployee) GetEmployeeByID(ctx context.Context, id int) (domain.Employee, error) {
	ms.MethodCalled = true
	if ms.Err != "" {
		return domain.Employee{}, errors.New(ms.Err)
	}
	for _, value := range ms.DataMock {
		if value.ID == id {
			return value, nil
		}
	}
	return domain.Employee{}, errors.New("employee not found")
}

func (ms *MockServiceEmployee) Update(ctx context.Context, id int, name string, lastname string, wharehouseId *int) (domain.Employee, error) {
	ms.MethodCalled = true
	index := 0
	flag := false
	for i := range ms.DataMock {
		if ms.DataMock[i].ID == id {
			index = i
			flag = true
		}
	}
	if !flag {
		return domain.Employee{}, errors.New("employee not found")
	}
	if name != "" {
		ms.DataMock[index].FirstName = name
	}
	if lastname != "" {
		ms.DataMock[index].LastName = lastname
	}
	if wharehouseId != nil {
		ms.DataMock[index].WarehouseID = *wharehouseId
	}
	return ms.DataMock[index], nil
}

func (ms *MockServiceEmployee) Delete(ctx context.Context, id int) error {
	ms.MethodCalled = true
	index := 0
	flag := false
	for i := range ms.DataMock {
		if ms.DataMock[i].ID == id {
			index = i
			flag = true
		}
	}
	if !flag {
		return errors.New("employee not found")
	}
	ms.DataMock = append(ms.DataMock[:index], ms.DataMock[index+1:]...)
	return nil
}

func (ms *MockServiceEmployee) Report_BO(ctx context.Context, id string) ([]domain.ReportInBO, error) {
	ms.MethodCalled = true
	if ms.Err != "" {
		return []domain.ReportInBO{}, errors.New(ms.Err)
	}
	if id != "" {
		id_e, err := strconv.Atoi(id)
		if err != nil {
			return []domain.ReportInBO{}, errors.New("invalid id format")
		}
		for _, r := range ms.DatamockIBO {
			if r.ID == id_e {
				return []domain.ReportInBO{r}, nil
			}
		}
		return []domain.ReportInBO{}, errors.New("employee not found")
	}
	return ms.DatamockIBO, nil
}
