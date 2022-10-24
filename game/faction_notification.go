package game

type notificationStruct struct {
	Header, Text string
}

func (f *faction) GetNotifications() []*notificationStruct {
	return f.notificationsForThisTurn
}

func (f *faction) addNotification(header, text string) {
	f.notificationsForThisTurn = append(f.notificationsForThisTurn, &notificationStruct{Header: header, Text: text})
}

func (f *faction) clearNotifications() {
	f.notificationsForThisTurn = nil
}
