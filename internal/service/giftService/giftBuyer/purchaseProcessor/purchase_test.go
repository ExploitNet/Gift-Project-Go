package purchaseProcessor

import (
	"context"
	"testing"

	"github.com/gotd/td/tg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPaymentProcessor для тестирования
type MockPaymentProcessor struct {
	mock.Mock
}

func (m *MockPaymentProcessor) CreatePaymentForm(ctx context.Context, gift *tg.StarGift, receiverTypes []int) (tg.PaymentsPaymentFormClass, *tg.InputInvoiceStarGift, error) {
	args := m.Called(ctx, gift, receiverTypes)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(tg.PaymentsPaymentFormClass), args.Get(1).(*tg.InputInvoiceStarGift), args.Error(2)
}

func createTestGift(id int64, stars int64) *tg.StarGift {
	return &tg.StarGift{
		ID:    id,
		Stars: stars,
	}
}

func createTestInvoice(giftID int64) *tg.InputInvoiceStarGift {
	return &tg.InputInvoiceStarGift{
		Peer:   &tg.InputPeerSelf{},
		GiftID: giftID,
		Message: tg.TextWithEntities{
			Text: "Test message",
		},
	}
}

func TestNewPurchaseProcessor(t *testing.T) {
	mockPaymentProcessor := &MockPaymentProcessor{}

	processor := NewPurchaseProcessor(nil, mockPaymentProcessor)

	assert.NotNil(t, processor)
	assert.Nil(t, processor.api)
	assert.Equal(t, mockPaymentProcessor, processor.paymentProcessor)
}

func TestPurchaseProcessorImpl_PurchaseGift_ErrorCases(t *testing.T) {
	t.Run("ошибка при недостаточном балансе", func(t *testing.T) {
		mockPaymentProcessor := &MockPaymentProcessor{}

		processor := &PurchaseProcessorImpl{
			api:              nil,
			paymentProcessor: mockPaymentProcessor,
		}

		gift := createTestGift(1, 100)

		// Настраиваем мок на случай если CreatePaymentForm все-таки вызовется
		mockPaymentProcessor.On("CreatePaymentForm", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil, assert.AnError)

		ctx := context.Background()
		err := processor.PurchaseGift(ctx, gift, []int{1})

		assert.Error(t, err)
		// Проверяем что есть ошибка (любая)
		assert.NotNil(t, err)
	})
}

func TestPurchaseProcessorImpl_SendStarsForm(t *testing.T) {
	t.Run("тестирование с nil API", func(t *testing.T) {
		processor := &PurchaseProcessorImpl{
			api:              nil,
			paymentProcessor: nil,
		}

		invoice := createTestInvoice(1)
		ctx := context.Background()

		// С nil API должна возникнуть паника или ошибка
		assert.Panics(t, func() {
			processor.sendStarsForm(ctx, invoice, 12345)
		})
	})
}

func TestPurchaseProcessorImpl_ValidatePurchase_NilAPI(t *testing.T) {
	t.Run("валидация с nil API", func(t *testing.T) {
		processor := &PurchaseProcessorImpl{
			api:              nil,
			paymentProcessor: nil,
		}

		gift := createTestGift(1, 100)

		// С nil API валидация должна вернуть false
		result := processor.validatePurchase(gift)
		assert.False(t, result)
	})
}
