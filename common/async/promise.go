package async

// A wrapper around goroutines that mimics the behaviour
// of promises in JavaScript with a golang twist :)
type Promise struct {
	channel chan promiseResult
}

// Takes a callback and wraps it into a promise using the goroutines
// Note: Can be improved to Accept Functions with more parameters using golang "reflect" package
func NewPromise(callback func() (interface{}, error)) Promise {
	promise := Promise{
		channel: make(chan promiseResult),
	}
	go func() {
		defer close(promise.channel)
		res, err := callback()
		if err != nil {
			promise.reject(err)
		}
		promise.resolve(res)
	}()
	return promise
}

// Await for the promise to complete and return the result
//if successful or error in case of failure
func (p *Promise) Await() (interface{}, error) {
	result := <-p.channel
	return result.result, result.err
}

// Awaits for all Promises to be resolved or one get rejected
func WaitAll(promises ...Promise) Promise {
	promise := NewPromise(func() (interface{}, error) { return waitAll(promises...) })
	return promise
}

func waitAll(promises ...Promise) (interface{}, error) {
	all := make([]interface{}, 0, len(promises))
	for _, pr := range promises {
		result, err := pr.Await()
		if err != nil {
			return nil, err
		}
		all = append(all, result)
	}

	return all, nil
}

// Represents the result of the promise in both case reject/resolve
type promiseResult struct {
	result interface{}
	err    error
}

// Promise complete with success
func (p *Promise) resolve(result interface{}) {
	p.channel <- promiseResult{
		result: result,
		err:    nil,
	}
}

// Promise complete with an error
func (p *Promise) reject(err error) {
	p.channel <- promiseResult{
		result: nil,
		err:    err,
	}
}
