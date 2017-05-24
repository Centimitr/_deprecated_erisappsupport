package entity

type Reader struct {
	mapping map[string]Book
}

func (r *Reader) Init() {
	r.mapping = make(map[string]Book)
}

func (r *Reader) Open(locator string) error {
	b := Book{}
	err := b.Init(locator)
	if err != nil {
		return err
	}
	r.mapping[locator] = b
	return nil
}

func (r *Reader) Get(locator string) (b Book, err error) {
	if _, ok := r.mapping[locator]; !ok {
		err = r.Open(locator)
	}
	return r.mapping[locator], err
}
