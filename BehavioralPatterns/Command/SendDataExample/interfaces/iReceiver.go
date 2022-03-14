package interfaces

import "base/BehavioralPatterns/Command/SendDataExample/domain"

type IReceiver interface {
	ReceiveData(data domain.Data)
}
