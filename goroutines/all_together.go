package goroutines

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type AOut string
type BOut string
type COut struct {
	data string
}
type CIn struct {
	A AOut
	B BOut
}

type Input struct {
	A string
	B string
	C string
}

func InitGatherAndProcess() {
	ctx := context.Background()
	input := Input{
		A: "a",
		B: "b",
		C: "c",
	}
	cOut, err := gatherAndProcess(ctx, input)
	if err != nil {
		panic(err)
	}
	println(cOut.data)
}
func gatherAndProcess(ctx context.Context, data Input) (COut, error) {

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		outC: make(chan COut, 1),
		inC:  make(chan CIn, 1),
	}

	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)

	if err != nil {
		return COut{}, err
	}

	p.inC <- inputC
	out, err := p.waitForC(ctx)

	return out, err
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.err <- err
			return
		}
		p.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.err <- err
			return
		}
		p.outB <- bOut
	}()

	go func() {
		cOut, err := getResultC(ctx, data.C)
		if err != nil {
			p.err <- err
			return
		}
		p.outC <- cOut
	}()
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0

	for count < 2 {
		select {
		case a := <-p.outA:
			inputC.A = a
			count++
		case b := <-p.outB:
			inputC.B = b
			count++
		case err := <-p.err:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}

	return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.err:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	err  chan error
}

var once sync.Once

func simulateTimeResponse() {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	time.Sleep(time.Duration(rand.Intn(110)) * time.Millisecond)
}

func getResultA(ctx context.Context, in string) (AOut, error) {
	simulateTimeResponse()
	res := AOut(fmt.Sprintf("result is the input: %s", in))
	fmt.Println(res)
	return res, nil
}

func getResultB(ctx context.Context, in string) (BOut, error) {
	simulateTimeResponse()
	res := BOut(fmt.Sprintf("result is the input: %s", in))
	fmt.Println(res)
	return res, nil
}
func getResultC(ctx context.Context, in string) (COut, error) {
	simulateTimeResponse()
	var cOut = COut{fmt.Sprintf("result is the input: %s", in)}

	return cOut, nil
}
