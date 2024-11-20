package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageCover(done In, val interface{}, stage Stage) Out {
	out := make(Bi)
	in := make(Bi)
	in <- val

	go func() {
		defer close(out)
		defer close(in)
		select {
		case <-done:
			return
		case out <- stage(in):
		}
	}()

	return out
}
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(in)
	}
	return in
}

func ExecutePipeline_(in In, done In, stages ...Stage) Out {

	localOut := make(Bi)

	go func() {
		defer close(localOut)

		for value := range in {
			//for {
			//	select {
			//	case <-done:
			//		return
			//	case value, okFromIn := <-in:
			//		if !okFromIn {
			//			return
			//		}
			for _, stage := range stages {
				select {
				case <-done:
					return
				case val, ok := <-stageCover(done, value, stage):
					if !ok {
						return
					}
					value = val
				}
			}
			localOut <- value
			//	}
			//
		}
	}()

	return localOut
	//chanel := in
	//for _, stage := range stages {
	//	select {
	//	case <-done:
	//		return nil
	//	default:
	//		chanel = stage(chanel)
	//	}
	//
	//}
	//
	//return chanel

	//out := make(Bi)
	//wg := sync.WaitGroup{}
	//go func() {
	//	defer close(out)
	//
	//	for val := range in {
	//		wg.Add(1)
	//		go func(val interface{}) {
	//			for _, stage := range stages {
	//
	//			}
	//			out <- strconv.Itoa(val.(int))
	//			wg.Done()
	//		}(val)
	//
	//	}
	//	wg.Wait()
	//}()
	//
	//return out
}

//func sendToStages(val interface{}, done In, stages ...Stage) Out {
//
//	inLocal := make(Bi)
//	outLocal := make(Bi)
//	inLocal <- val
//	for _, stage := range stages {
//
//		out := stage(inLocal)
//		select {
//		case <-done:
//
//		}
//	}
//}
