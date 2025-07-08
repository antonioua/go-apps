package kata

func kata() {
	// channel a contains 3 messages
	a := make(chan string, 3)
	a <- "foo"
	a <- "bar"
	a <- "baz"
	close(a)

	// channel b contains 2 messages
	b := make(chan string, 2)
	b <- "hello"
	b <- "world"
	close(b)

	// your implementation
	c := Merge(a, b)

	// when messages are consumed from c, it must return all 5 messages from a and b,
	// while the order of the messages is not defined
	<-c // "foo"
	<-c // "bar"
	<-c // "hello"
	<-c // "world"
	<-c // "baz"

	// channel c must be closed at this point
}

func Merge(a chan, b chan ) {

}
