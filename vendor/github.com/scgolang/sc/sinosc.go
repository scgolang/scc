package sc

// SinOsc is a table-lookup sinewave oscillator
type SinOsc struct {
	// Freq is frequency in Hz
	Freq Input
	// Phase is the initial phase offset
	Phase Input
}

func (sin *SinOsc) defaults() {
	if sin.Freq == nil {
		sin.Freq = C(440)
	}
	if sin.Phase == nil {
		sin.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (sin SinOsc) Rate(rate int8) Input {
	CheckRate(rate)
	(&sin).defaults()
	return UgenInput("SinOsc", rate, 0, 1, sin.Freq, sin.Phase)
}

func init() {
	if err := RegisterSynthdef("sine", func(params Params) Ugen {
		var (
			add   = params.Add("add", 0)
			mul   = params.Add("mul", 1)
			out   = params.Add("out", 0)
			freq  = params.Add("freq", 440)
			phase = params.Add("phase", 0)
		)
		return Out{
			Bus: out,
			Channels: SinOsc{
				Freq:  freq,
				Phase: phase,
			}.Rate(AR).MulAdd(mul, add),
		}.Rate(AR)
	}); err != nil {
		panic(err)
	}
	if err := RegisterSynthdef("sine_ctl", func(params Params) Ugen {
		var (
			add   = params.Add("add", 0)
			mul   = params.Add("mul", 1)
			out   = params.Add("out", 0)
			freq  = params.Add("freq", 440)
			phase = params.Add("phase", 0)
		)
		return Out{
			Bus: out,
			Channels: SinOsc{
				Freq:  freq,
				Phase: phase,
			}.Rate(KR).MulAdd(mul, add),
		}.Rate(KR)
	}); err != nil {
		panic(err)
	}
}
