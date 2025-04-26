package consumers



func RunConsumers() {
	// Add consumers here with `go` keyword
	NewTempHumidityConsumer().Run()
}
