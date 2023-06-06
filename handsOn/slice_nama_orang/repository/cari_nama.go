package repository

import (
	"cari_nama/model"
	"strings"
)

func CariMuhammad(name []string) []model.NameMuhammad {
	var muhName []model.NameMuhammad
	
	for _, v := range name {
		if strings.Contains(strings.ToLower(v), "muhammad") {
			newNameMuh := model.NameMuhammad{v}
			
			muhName = append(muhName, newNameMuh)
		}
	}
	
	return muhName
}
