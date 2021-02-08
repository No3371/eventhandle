# eventhandle
go pub/sub implementation based on channels.

The 'subscribers' indicates the channels, the publisher/handle does not have to know anything about who is subscribing, it just create channals and deliver the events, the event type is interface{} so consensus is needed.

## Usage
**Subscribe to a handle**
`eventhandle.Subscribe(bufferSize int)`

**Unsubscribe from a handle**
`eventhandle.Subscribe(subscriber chan interface{})`

**Publish through a handle**
`eventhandle.Publish(e interface{})`
