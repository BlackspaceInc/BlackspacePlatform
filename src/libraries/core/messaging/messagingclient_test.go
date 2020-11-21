package messaging

/*
func TestMessageHandlerLoop(t *testing.T) {

	var invocations = 0

	var handlerFunction = func(d amqp.Delivery) {
		logrus.Println("In handlerFunction")
		invocations = invocations + 1
	}

	Convey("Given", t, func() {
		var messageChannel = make(chan amqp.Delivery, 1)
		go consumeLoop(messageChannel, handlerFunction)

		Convey("When", func() {
			d := amqp.Delivery{Body: []byte(""), ConsumerTag: ""}
			messageChannel <- d
			messageChannel <- d
			messageChannel <- d
			Convey("Then", func() {
				So(invocations, ShouldEqual, 3)
			})
		})
	})
}
*/
