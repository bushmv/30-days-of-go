package main

func main() {
	r := NewNotificationRepository(NewSliceRepository(), NewConsoleNotifier())

	_, _ = r.Save(&Item{ID: NoID, Title: "a"})
	_, _ = r.Save(&Item{ID: NoID, Title: "b"})
	_, _ = r.Save(&Item{ID: NoID, Title: "c"})
	_, _ = r.GetById(1)
	_, _ = r.GetById(99)
	_, _ = r.GetAll()
}
