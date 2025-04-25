package consumers



func RunConsumers() {
	NewTempHumidityConsumer().Run()
}
