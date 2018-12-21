package restart

type restore interface {
	save()
	load()
}
