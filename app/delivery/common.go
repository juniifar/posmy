package delivery

// Ping handler
func (impl *Deliveries) Ping() {
	ctx := impl.Ctx
	ctx.WriteString("Pong")
}

// Index handler
func (impl *Deliveries) Index() {
	ctx := impl.Ctx

	ctx.WriteString("OK")
}
