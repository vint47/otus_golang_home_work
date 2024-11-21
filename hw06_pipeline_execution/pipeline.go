package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func clearChanel(out Out) {
	go func() {
		for temp := range out {
			_ = temp
		}
	}()
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	stageCover := func(in In, stage Stage) Out {
		out := make(Bi)

		go func() {
			defer close(out)
			stageOut := stage(in)

			for {
				select {
				case <-done:
					clearChanel(stageOut)
					return
				case val, ok := <-stageOut:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case out <- val:
					}
				}
			}
		}()

		return out
	}

	for _, stage := range stages {
		in = stageCover(in, stage)
	}

	return in
}
