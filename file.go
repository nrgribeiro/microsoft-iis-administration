package iis

type File struct {
	Name            string               `json:"name"`
	PhysicalPath    string               `json:"physical_path"`
	Type			string				 `json:"type"`
	Alias 			string				 `json:"alias"`
	Parent 			[]FileParent		 `json:"parent"`
}

type FileParent struct {
	Id  string `json:"id"`
}