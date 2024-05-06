package middleware

type MW interface {
	Start(args ...string)
	Stop(args ...string)
	Write(args ...string)
}

type MWs map[string]MW

func (mws MWs) Start(args ...string) {
	for _, mw := range mws {
		mw.Start(args...)
	}
}

func (mws MWs) Stop(args ...string) {
	for _, mw := range mws {
		mw.Stop(args...)
	}
}

func (mws MWs) Write(middleware string, args ...string) {
	mws[middleware].Write(args...)
}

func GetMiddleware() MWs {
	Middlewares := make(map[string]MW)
	Middlewares[TimeChecker] = &TimeCheck{}

	return Middlewares
}
