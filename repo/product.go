package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {

	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Very Sweet Orange, Sweet like Sadie Sink",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/W4CyOmvwDLOtnG7gB2uDEg7m-dOaXcbbNQefKDn0sxU/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tcHNk/L29yYW5nZS1pc29s/YXRlZC1hbHBoYS1s/YXllcl82MTA1Mzkt/NTA2LmpwZz9nYT1H/QTEuMS4xMDA5Mjg5/MjAyLjE3OTEwOTc5/MTQmc2VtdD1haXNf/aHlicmlkJnc9NzQw/JnE9ODA",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Greem Apple, Greener then Me",
		Price:       40,
		ImgUrl:      "https://imgs.search.brave.com/Uy-TwyDej-DhWzNmsI5dxtkBMtUdm6AxE0AIzfEHP0o/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L2ZyZWUtcHNkL2ds/aXN0ZW5pbmctcmVk/LWFwcGxlLWNvdmVy/ZWQtZnJlc2gtd2F0/ZXItZHJvcGxldHMt/aXNvbGF0ZWQtYmxh/Y2stYmFja2dyb3Vu/ZF84NDQ0My01OTY4/Ni5qcGc_c2VtdD1h/aXNfaHlicmlkJnc9/NzQwJnE9ODA",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Sweet & Best Looking Banana, You can eat it as a preworkout",
		Price:       30,
		ImgUrl:      "https://imgs.search.brave.com/OxsX3sE85TBAL3Uj1K-x2JwMWlgVOQA2y2gBX1zpsJw/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tcGhv/dG8vcGVlbGVkLWJh/bmFuYS1pc29sYXRl/ZC13aGl0ZS1iYWNr/Z3JvdW5kLXdpdGgt/Y2xpcHBpbmctcGF0/aF84ODI4MS05Ny5q/cGc_c2VtdD1haXNf/aHlicmlkJnc9NzQw/JnE9ODA",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Mango",
		Description: "Sweet & Best Looking Mango, Sweet and juicy like dure fishan",
		Price:       300,
		ImgUrl:      "https://imgs.search.brave.com/Qdt4LAKeCEpJDKmADYsJ7APYmx0yskyz8lD1T047StY/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly90My5m/dGNkbi5uZXQvanBn/LzE5LzUyLzc5LzI2/LzM2MF9GXzE5NTI3/OTI2NTBfaGt6THpY/cmwxY2xjQTkzVlVq/ejVIRmxBa1prNnhp/bnkuanBn",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Graps",
		Description: "Sweet Graps, You can eat it or make wine witrh it",
		Price:       400,
		ImgUrl:      "https://imgs.search.brave.com/gfEvUuxF_dFGQBCbemgOdAZzQ9R8ij07rtrCzP-VRok/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5nZXR0eWltYWdl/cy5jb20vaWQvMTA0/ODIyMTczL3Bob3Rv/L2Nsb3NlLXVwLW9m/LWdyYXBlcy5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9bUJu/N19SM3J0SWlCN0VB/dzFVQUgwOElvNldi/NHoxdTVSM0VJQ1Z4/MlhZOD0",
	}
	prd6 := &Product{
		ID:          6,
		Title:       "Papaya",
		Description: "Sweet Papaya, Best Papaya in this world",
		Price:       50,
		ImgUrl:      "https://imgs.search.brave.com/LwbfjH94V-AQgY8tO06FwxkgqIBSNfaeSRxGvy1F4Js/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMjYv/NzUxLzE3My9zbWFs/bC9nZW5lcmF0aXZl/LWFpLW1hY3JvLWZy/ZXNoLWhhbGYtb2Yt/cGFwYXlhLWZydWl0/LWJhY2tncm91bmQt/dHJvcGljYWwtZXhv/dGljLWNsb3NldXAt/d2l0aC1kcm9wcy1w/aG90by5qcGc",
	}

	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd5)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd6)
}
