package main

import "simple/pkg/database"

func main() {
	if generator, err := database.NewGenerator(&database.GenConfig{
		DSN:             "root:wui11413@tcp(127.0.0.1:3306)/simple?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai",
		OutPath:         "./internal/types/query",
		ModelPkgPath:    "./internal/types/entity",
		TablePrefix:     "sys_",
		WithQueryFilter: true,
	}); err != nil {
		panic(err)
	} else {
		//generator.GenerateAllModel().v
		generator.GenerateModelsWithRelations()
	}
}
